package sip

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type Server struct {
	config          *ServerConfig
	transport       Transport
	transactionMgr  *TransactionManager
	dialogMgr       *DialogManager
	messageHandlers map[string]func(*SIPMessage, string)
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
		messageHandlers: make(map[string]func(*SIPMessage, string)),
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

func (s *Server) RegisterHandler(method string, handler func(*SIPMessage, string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messageHandlers[method] = handler
}

func (s *Server) handleMessage(addr string, data []byte) {
	msg, err := ParseMessage(data)
	if err != nil {
		s.logger.Error("Parse SIP message failed",
			zap.String("addr", addr),
			zap.Error(err),
		)
		return
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
		handler(msg, addr)
	} else {
		s.logger.Warn("No handler for SIP method",
			zap.String("method", method),
		)
		s.sendResponse(txn, 405, "Method Not Allowed")
	}
}

func (s *Server) handleResponse(msg *SIPMessage, addr string) {
	branch := msg.GetHeader("Via")
	txn := s.transactionMgr.FindByBranch(branch)
	if txn != nil {
		txn.HandleResponse(msg)
	}
}

func (s *Server) sendResponse(txn *Transaction, statusCode int, reason string) {
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
		response.SetHeader("Via", txn.Request.GetHeader("Via"))
		response.SetHeader("From", txn.Request.GetHeader("From"))
		response.SetHeader("To", txn.Request.GetHeader("To"))
		response.SetHeader("Call-ID", txn.Request.GetHeader("Call-ID"))
		response.SetHeader("CSeq", txn.Request.GetHeader("CSeq"))
	}

	txn.SendResponse(response)
}

func (s *Server) SendRequest(method, requestURI string, headers map[string]string, body []byte) (*Transaction, error) {
	builder := NewBuilder()
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

	txn, err := s.transactionMgr.CreateClientTransaction(msg, s.transport, requestURI)
	if err != nil {
		return nil, err
	}

	if err := txn.SendRequest(); err != nil {
		return nil, err
	}

	return txn, nil
}