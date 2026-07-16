package sip

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"go.uber.org/zap"
)

// MessageHandler processes a SIP request and sends an appropriate response
// via the provided Transaction. The handler MUST call txn.SendResponse() for
// every request it handles (including non-2xx responses).
type MessageHandler func(msg *SIPMessage, addr string, txn *Transaction)

type Server struct {
	config          *ServerConfig
	transport       Transport
	transactionMgr  *TransactionManager
	dialogMgr       *DialogManager
	messageHandlers map[string]MessageHandler
	running         bool
	mu              sync.RWMutex
	logger          *zap.Logger
}

type ServerConfig struct {
	ListenIP    string
	ListenPort  int
	Domain      string
	ServerID    string
	Transport   string
	SIPLog      bool
}

func NewServer(config *ServerConfig, logger *zap.Logger) *Server {
	protocol := UDP
	if config.Transport == "TCP" {
		protocol = TCP
	}

	return &Server{
		config:          config,
		transport:       NewTransport(protocol, logger),
		transactionMgr:  NewTransactionManager(logger),
		dialogMgr:       NewDialogManager(logger),
		messageHandlers: make(map[string]MessageHandler),
		logger:          logger,
	}
}

func (s *Server) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	addr := fmt.Sprintf("%s:%d", s.config.ListenIP, s.config.ListenPort)
	if err := s.transport.Listen(addr); err != nil {
		return fmt.Errorf("start SIP server failed: %w", err)
	}

	s.transport.SetMessageHandler(s.handleMessage)
	s.running = true

	s.logger.Info("SIP server started",
		zap.String("addr", addr),
		zap.String("domain", s.config.Domain),
	)

	return nil
}

func (s *Server) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return nil
	}

	s.running = false
	return s.transport.Close()
}

func (s *Server) RegisterHandler(method string, handler MessageHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messageHandlers[method] = handler
}

func (s *Server) handleMessage(addr string, data []byte) {
	msg, err := ParseMessage(data)
	if err != nil {
		if s.config.SIPLog {
			s.logger.Warn("SIP Parse failed",
				zap.String("addr", addr),
				zap.String("raw", string(data)),
				zap.Error(err),
			)
		} else {
			s.logger.Error("Parse SIP message failed",
				zap.String("addr", addr),
				zap.Error(err),
			)
		}
		return
	}

	if s.config.SIPLog {
		s.logger.Info("\n" + FormatIncoming(msg, addr))
	}

	if msg.IsRequest {
		s.handleRequest(msg, addr)
	} else {
		s.handleResponse(msg, addr)
	}
}

func (s *Server) handleRequest(msg *SIPMessage, addr string) {
	method := msg.RequestLine.Method

	txn, err := s.transactionMgr.CreateServerTransaction(msg, s.transport, addr)
	if err != nil {
		s.logger.Error("Create server transaction failed",
			zap.String("method", method),
			zap.Error(err),
		)
		return
	}

	s.mu.RLock()
	handler, exists := s.messageHandlers[method]
	s.mu.RUnlock()

	if exists {
		// Handler is responsible for sending its own response via txn.SendResponse()
		handler(msg, addr, txn)
	} else {
		s.logger.Warn("No handler for SIP method",
			zap.String("method", method),
		)
		// Default: send 405 Method Not Allowed
		s.sendResponse(txn, 405, "Method Not Allowed", nil)
	}
}

func (s *Server) handleResponse(msg *SIPMessage, addr string) {
	branch := msg.GetHeader("Via")
	txn := s.transactionMgr.FindByBranch(branch)
	if txn != nil {
		txn.HandleResponse(msg)
	}
}

// sendResponse constructs a SIP response from the transaction request context
// and sends it. The optional extraHeaders map allows adding or overriding headers
// (e.g. WWW-Authenticate for 401 challenges).
func (s *Server) sendResponse(txn *Transaction, statusCode int, reason string, extraHeaders map[string]string) {
	response := &SIPMessage{
		IsRequest: false,
		StatusLine: &StatusLine{
			SIPVersion:   "SIP/2.0",
			StatusCode:   statusCode,
			ReasonPhrase: reason,
		},
		Headers: make(map[string][]string),
	}

	if txn.Request != nil {
		// Copy ALL Via headers in order (RFC 3261 §8.2.6)
		if vias, ok := txn.Request.Headers["Via"]; ok {
			for _, v := range vias {
				response.AddHeader("Via", v)
			}
		}
		response.SetHeader("From", txn.Request.GetHeader("From"))
		response.SetHeader("To", txn.Request.GetHeader("To"))
		response.SetHeader("Call-ID", txn.Request.GetHeader("Call-ID"))
		response.SetHeader("CSeq", txn.Request.GetHeader("CSeq"))
	}

	// Apply any extra headers (e.g. WWW-Authenticate for 401)
	for key, value := range extraHeaders {
		response.SetHeader(key, value)
	}

	if s.config.SIPLog {
		s.logger.Info("\n" + FormatOutgoing(response, txn.RemoteAddr))
	}

	txn.SendResponse(response)
}

// SendUnauthorized sends a 401 response with WWW-Authenticate Digest challenge.
func (s *Server) SendUnauthorized(txn *Transaction, realm, nonce string) {
	extraHeaders := map[string]string{
		"WWW-Authenticate": fmt.Sprintf(`Digest realm="%s", nonce="%s", algorithm=MD5, qop="auth"`, realm, nonce),
	}
	s.sendResponse(txn, 401, "Unauthorized", extraHeaders)
}

// SendOK sends a 200 OK response for the given transaction.
func (s *Server) SendOK(txn *Transaction) {
	s.sendResponse(txn, 200, "OK", nil)
}

func (s *Server) SendRequest(method, requestURI string, headers map[string]string, body []byte) (*Transaction, error) {
	builder := NewBuilder()
	builder.SetListenAddr(fmt.Sprintf("%s:%d", s.config.ListenIP, s.config.ListenPort))
	from := fmt.Sprintf("<sip:%s@%s>;tag=%s", s.config.ServerID, s.config.Domain, generateTag())
	to := fmt.Sprintf("<sip:%s@%s>", s.config.ServerID, s.config.Domain)
	callID := generateCallID()
	cseq := fmt.Sprintf("1 %s", method)

	msg := builder.BuildRequest(method, requestURI, from, to, callID, cseq)
	for key, value := range headers {
		msg.SetHeader(key, value)
	}
	if body != nil {
		msg.SetBody(body)
	}

	if s.config.SIPLog {
		s.logger.Info("\n" + FormatOutgoing(msg, requestURI))
	}

	txn, err := s.transactionMgr.CreateClientTransaction(msg, s.transport, requestURI)
	if err != nil {
		return nil, err
	}

	if err := txn.SendRequest(); err != nil {
		return nil, err
	}

	return txn, nil
}

func (s *Server) SendRequestTo(method, requestURI, targetAddr string, headers map[string]string, body []byte) (*Transaction, error) {
	host, port := parseAddr(targetAddr)
	if host == s.config.ListenIP || host == "127.0.0.1" || host == "0.0.0.0" {
		if port == s.config.ListenPort || port == 0 {
			s.logger.Warn("Skip sending to self",
				zap.String("target", targetAddr),
			)
			return nil, fmt.Errorf("cannot send to self: %s", targetAddr)
		}
	}

	builder := NewBuilder()
	builder.SetListenAddr(fmt.Sprintf("%s:%d", s.config.ListenIP, s.config.ListenPort))
	from := fmt.Sprintf("<sip:%s@%s>;tag=%s", s.config.ServerID, s.config.Domain, generateTag())
	to := headers["To"]
	if to == "" {
		to = fmt.Sprintf("<sip:%s@%s>", s.config.ServerID, s.config.Domain)
	}
	callID := generateCallID()
	cseq := fmt.Sprintf("1 %s", method)

	msg := builder.BuildRequest(method, requestURI, from, to, callID, cseq)
	for key, value := range headers {
		msg.SetHeader(key, value)
	}
	if body != nil {
		msg.SetBody(body)
	}

	if s.config.SIPLog {
		s.logger.Info("\n" + FormatOutgoing(msg, targetAddr))
	}

	txn, err := s.transactionMgr.CreateClientTransaction(msg, s.transport, targetAddr)
	if err != nil {
		return nil, err
	}

	if err := txn.SendRequest(); err != nil {
		return nil, err
	}

	return txn, nil
}

func parseAddr(addr string) (string, int) {
	if idx := strings.LastIndex(addr, ":"); idx != -1 {
		host := addr[:idx]
		port, _ := strconv.Atoi(addr[idx+1:])
		return host, port
	}
	return addr, 0
}

func (s *Server) SendInvite(deviceID, targetAddr, sdp, subject string) (*Transaction, error) {
	requestURI := fmt.Sprintf("sip:%s@%s", deviceID, s.config.Domain)
	contact := fmt.Sprintf("<sip:%s@%s:%d>", s.config.ServerID, s.config.ListenIP, s.config.ListenPort)
	headers := map[string]string{
		"To":      fmt.Sprintf("<sip:%s@%s>", deviceID, s.config.Domain),
		"Contact": contact,
	}
	if subject != "" {
		headers["Subject"] = subject
	}
	return s.SendRequestTo("INVITE", requestURI, targetAddr, headers, []byte(sdp))
}

// SendACK sends an ACK request for an established INVITE dialog.
func (s *Server) SendACK(targetAddr, callID, fromTag, toTag, cseq string) error {
	ackMsg := fmt.Sprintf("ACK sip:%s SIP/2.0\r\n"+
		"Via: SIP/2.0/%s %s:%d;branch=%s\r\n"+
		"From: <sip:%s@%s>;tag=%s\r\n"+
		"To: <sip:%s@%s>;tag=%s\r\n"+
		"Call-ID: %s\r\n"+
		"CSeq: %s ACK\r\n"+
		"Max-Forwards: 70\r\n"+
		"Content-Length: 0\r\n"+
		"\r\n",
		targetAddr,
		s.config.Transport, s.config.ListenIP, s.config.ListenPort, generateBranch(),
		s.config.ServerID, s.config.Domain, fromTag,
		s.config.ServerID, s.config.Domain, toTag,
		callID, cseq)

	data := []byte(ackMsg)
	return s.transport.Send(targetAddr, data)
}
