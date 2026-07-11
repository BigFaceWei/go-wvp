package sip

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type Client struct {
	config         *ClientConfig
	transport      Transport
	transactionMgr *TransactionManager
	dialogMgr      *DialogManager
	running        bool
	mu             sync.RWMutex
	logger         *zap.Logger
}

type ClientConfig struct {
	LocalIP    string
	LocalPort  int
	Domain     string
	ServerID   string
	RemoteIP   string
	RemotePort int
	Transport  string
}

func NewClient(config *ClientConfig, logger *zap.Logger) *Client {
	protocol := UDP
	if config.Transport == "TCP" {
		protocol = TCP
	}

	return &Client{
		config:         config,
		transport:      NewTransport(protocol, logger),
		transactionMgr: NewTransactionManager(logger),
		dialogMgr:      NewDialogManager(logger),
		logger:         logger,
	}
}

func (c *Client) Start() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	addr := fmt.Sprintf("%s:%d", c.config.LocalIP, c.config.LocalPort)
	if err := c.transport.Listen(addr); err != nil {
		return fmt.Errorf("start SIP client failed: %w", err)
	}

	c.transport.SetMessageHandler(c.handleMessage)
	c.running = true

	c.logger.Info("SIP client started", zap.String("addr", addr))
	return nil
}

func (c *Client) Stop() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.running {
		return nil
	}

	c.running = false
	return c.transport.Close()
}

func (c *Client) handleMessage(addr string, data []byte) {
	msg, err := ParseMessage(data)
	if err != nil {
		c.logger.Error("Parse SIP message failed", zap.Error(err))
		return
	}

	if !msg.IsRequest {
		branch := msg.GetHeader("Via")
		txn := c.transactionMgr.FindByBranch(branch)
		if txn != nil {
			txn.HandleResponse(msg)
		}
	}
}

func (c *Client) SendRequest(method, requestURI string, headers map[string]string, body []byte) (*Transaction, error) {
	remoteAddr := fmt.Sprintf("%s:%d", c.config.RemoteIP, c.config.RemotePort)

	builder := NewBuilder()
	from := fmt.Sprintf("<sip:%s@%s>;tag=%s", c.config.ServerID, c.config.Domain, generateTag())
	to := fmt.Sprintf("<sip:%s@%s>", c.config.ServerID, c.config.Domain)
	callID := generateCallID()
	cseq := fmt.Sprintf("1 %s", method)

	msg := builder.BuildRequest(method, requestURI, from, to, callID, cseq)
	for key, value := range headers {
		msg.SetHeader(key, value)
	}
	if body != nil {
		msg.SetBody(body)
	}

	txn, err := c.transactionMgr.CreateClientTransaction(msg, c.transport, remoteAddr)
	if err != nil {
		return nil, err
	}

	if err := txn.SendRequest(); err != nil {
		return nil, err
	}

	return txn, nil
}

func (c *Client) SendRegister(expires int) (*Transaction, error) {
	requestURI := fmt.Sprintf("sip:%s", c.config.Domain)
	headers := map[string]string{
		"Expires": fmt.Sprintf("%d", expires),
		"Contact": fmt.Sprintf("<sip:%s@%s:%d>", c.config.ServerID, c.config.LocalIP, c.config.LocalPort),
	}

	return c.SendRequest("REGISTER", requestURI, headers, nil)
}

func (c *Client) SendKeepalive(deviceID string) (*Transaction, error) {
	requestURI := fmt.Sprintf("sip:%s@%s", deviceID, c.config.Domain)
	return c.SendRequest("OPTIONS", requestURI, nil, nil)
}