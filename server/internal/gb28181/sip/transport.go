package sip

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

type TransportProtocol string

const (
	UDP TransportProtocol = "UDP"
	TCP TransportProtocol = "TCP"
)

type Transport interface {
	Listen(addr string) error
	Send(addr string, data []byte) error
	SetMessageHandler(handler func(addr string, data []byte))
	Close() error
	Protocol() TransportProtocol
}

// tcpConnState holds per-connection state for TCP message framing.
type tcpConnState struct {
	conn net.Conn
	buf  bytes.Buffer // accumulated partial data waiting to form a complete SIP message
}

type UDPTransport struct {
	conn       *net.UDPConn
	handler    func(addr string, data []byte)
	running    bool
	mu         sync.RWMutex
	stopCh     chan struct{}
	logger     *zap.Logger
}

func NewUDPTransport(logger *zap.Logger) *UDPTransport {
	if logger == nil {
		logger = zap.NewNop()
	}
	return &UDPTransport{
		logger: logger,
		stopCh: make(chan struct{}),
	}
}

func (t *UDPTransport) Protocol() TransportProtocol {
	return UDP
}

func (t *UDPTransport) Listen(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return fmt.Errorf("resolve UDP addr failed: %w", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return fmt.Errorf("listen UDP failed: %w", err)
	}

	t.conn = conn
	t.running = true

	t.logger.Info("UDP transport started", zap.String("addr", addr))

	go t.readLoop()

	return nil
}

func (t *UDPTransport) readLoop() {
	buf := make([]byte, 65535)
	for {
		select {
		case <-t.stopCh:
			return
		default:
		}

		if t.conn == nil {
			return
		}
		t.conn.SetReadDeadline(time.Now().Add(time.Second))
		n, remoteAddr, err := t.conn.ReadFromUDP(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			t.logger.Error("UDP read error", zap.Error(err))
			continue
		}

		data := make([]byte, n)
		copy(data, buf[:n])

		if t.handler != nil {
			go t.handler(remoteAddr.String(), data)
		}
	}
}

func (t *UDPTransport) Send(addr string, data []byte) error {
	if t.conn == nil {
		return fmt.Errorf("UDP transport not connected")
	}
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return fmt.Errorf("resolve addr failed: %w", err)
	}

	_, err = t.conn.WriteToUDP(data, udpAddr)
	if err != nil {
		return fmt.Errorf("send UDP failed: %w", err)
	}

	return nil
}

func (t *UDPTransport) SetMessageHandler(handler func(addr string, data []byte)) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.handler = handler
}

func (t *UDPTransport) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.running {
		return nil
	}

	t.running = false
	close(t.stopCh)

	if t.conn != nil {
		return t.conn.Close()
	}
	return nil
}

type TCPTransport struct {
	listener    net.Listener
	handler     func(addr string, data []byte)
	running     bool
	mu          sync.RWMutex
	connections map[string]*tcpConnState
	stopCh      chan struct{}
	logger      *zap.Logger
}

func NewTCPTransport(logger *zap.Logger) *TCPTransport {
	if logger == nil {
		logger = zap.NewNop()
	}
	return &TCPTransport{
		logger:      logger,
		connections: make(map[string]*tcpConnState),
		stopCh:      make(chan struct{}),
	}
}

func (t *TCPTransport) Protocol() TransportProtocol {
	return TCP
}

func (t *TCPTransport) Listen(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen TCP failed: %w", err)
	}

	t.listener = listener
	t.running = true

	t.logger.Info("TCP transport started", zap.String("addr", addr))

	go t.acceptLoop()

	return nil
}

func (t *TCPTransport) acceptLoop() {
	for {
		select {
		case <-t.stopCh:
			return
		default:
		}

		conn, err := t.listener.Accept()
		if err != nil {
			t.logger.Error("TCP accept error", zap.Error(err))
			continue
		}

		addr := conn.RemoteAddr().String()
		state := &tcpConnState{conn: conn}
		t.mu.Lock()
		t.connections[addr] = state
		t.mu.Unlock()

		go t.handleConnection(state, addr)
	}
}

// readSIPMessage tries to extract one complete SIP message from the buffer.
// Returns nil if we don't have enough data yet.
func readSIPMessage(buf *bytes.Buffer) []byte {
	data := buf.Bytes()
	// Find the empty line (\r\n\r\n) that separates headers from body
	headerEnd := bytes.Index(data, []byte("\r\n\r\n"))
	if headerEnd == -1 {
		// Incomplete headers — need more data
		return nil
	}

	// Parse Content-Length from headers
	headersPart := string(data[:headerEnd])
	contentLength := 0
	for _, line := range strings.Split(headersPart, "\r\n") {
		if strings.HasPrefix(strings.ToLower(line), "content-length:") {
			val := strings.TrimSpace(line[len("content-length:"):])
			if n, err := strconv.Atoi(val); err == nil {
				contentLength = n
			}
			break
		}
	}

	// Body starts after \r\n\r\n
	bodyStart := headerEnd + 4
	totalLen := bodyStart + contentLength

	if len(data) < totalLen {
		// Haven't received the full body yet
		return nil
	}

	// Extract the complete message
	msg := make([]byte, totalLen)
	copy(msg, data[:totalLen])

	// Remove the consumed message from the buffer
	remaining := make([]byte, buf.Len()-totalLen)
	copy(remaining, data[totalLen:])
	buf.Reset()
	buf.Write(remaining)

	return msg
}

func (t *TCPTransport) handleConnection(state *tcpConnState, addr string) {
	defer func() {
		state.conn.Close()
		t.mu.Lock()
		delete(t.connections, addr)
		t.mu.Unlock()
	}()

	buf := make([]byte, 65535)
	for {
		select {
		case <-t.stopCh:
			return
		default:
		}

		state.conn.SetReadDeadline(time.Now().Add(time.Second))
		n, err := state.conn.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				// On timeout, try to extract any pending complete messages from buffer
				for {
					msg := readSIPMessage(&state.buf)
					if msg == nil {
						break
					}
					if t.handler != nil {
						go t.handler(addr, msg)
					}
				}
				continue
			}
			// Before closing, flush any remaining complete messages
			for {
				msg := readSIPMessage(&state.buf)
				if msg == nil {
					break
				}
				if t.handler != nil {
					t.handler(addr, msg)
				}
			}
			t.logger.Error("TCP read error", zap.String("addr", addr), zap.Error(err))
			return
		}

		// Append new data to connection buffer
		state.buf.Write(buf[:n])

		// Extract all complete messages from buffer
		for {
			msg := readSIPMessage(&state.buf)
			if msg == nil {
				break
			}
			if t.handler != nil {
				go t.handler(addr, msg)
			}
		}
	}
}

func (t *TCPTransport) Send(addr string, data []byte) error {
	t.mu.RLock()
	state, ok := t.connections[addr]
	t.mu.RUnlock()

	if !ok {
		newConn, err := net.Dial("tcp", addr)
		if err != nil {
			return fmt.Errorf("connect TCP failed: %w", err)
		}
		state = &tcpConnState{conn: newConn}
		t.mu.Lock()
		t.connections[addr] = state
		t.mu.Unlock()
	}

	_, err := state.conn.Write(data)
	if err != nil {
		t.mu.Lock()
		delete(t.connections, addr)
		t.mu.Unlock()
		state.conn.Close()
		return fmt.Errorf("send TCP failed: %w", err)
	}

	return nil
}

func (t *TCPTransport) SetMessageHandler(handler func(addr string, data []byte)) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.handler = handler
}

func (t *TCPTransport) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.running {
		return nil
	}

	t.running = false
	close(t.stopCh)

	for addr, state := range t.connections {
		state.conn.Close()
		delete(t.connections, addr)
	}

	if t.listener != nil {
		return t.listener.Close()
	}
	return nil
}

func NewTransport(protocol TransportProtocol, logger *zap.Logger) Transport {
	switch protocol {
	case TCP:
		return NewTCPTransport(logger)
	default:
		return NewUDPTransport(logger)
	}
}
