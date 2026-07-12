package sip

import (
	"fmt"
	"net"
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
	listener   net.Listener
	handler    func(addr string, data []byte)
	running    bool
	mu         sync.RWMutex
	connections map[string]net.Conn
	stopCh     chan struct{}
	logger     *zap.Logger
}

func NewTCPTransport(logger *zap.Logger) *TCPTransport {
	if logger == nil {
		logger = zap.NewNop()
	}
	return &TCPTransport{
		logger:      logger,
		connections: make(map[string]net.Conn),
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
		t.mu.Lock()
		t.connections[addr] = conn
		t.mu.Unlock()

		go t.handleConnection(conn, addr)
	}
}

func (t *TCPTransport) handleConnection(conn net.Conn, addr string) {
	defer func() {
		conn.Close()
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

		conn.SetReadDeadline(time.Now().Add(time.Second))
		n, err := conn.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			t.logger.Error("TCP read error", zap.String("addr", addr), zap.Error(err))
			return
		}

		data := make([]byte, n)
		copy(data, buf[:n])

		if t.handler != nil {
			go t.handler(addr, data)
		}
	}
}

func (t *TCPTransport) Send(addr string, data []byte) error {
	t.mu.RLock()
	conn, ok := t.connections[addr]
	t.mu.RUnlock()

	if !ok {
		newConn, err := net.Dial("tcp", addr)
		if err != nil {
			return fmt.Errorf("connect TCP failed: %w", err)
		}
		conn = newConn
		t.mu.Lock()
		t.connections[addr] = conn
		t.mu.Unlock()
	}

	_, err := conn.Write(data)
	if err != nil {
		t.mu.Lock()
		delete(t.connections, addr)
		t.mu.Unlock()
		conn.Close()
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

	for addr, conn := range t.connections {
		conn.Close()
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