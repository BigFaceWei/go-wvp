package sip

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

type TransactionState int

const (
	TransactionTrying TransactionState = iota
	TransactionProceeding
	TransactionCompleted
	TransactionConfirmed
	TransactionTerminated
)

type Transaction struct {
	ID        string
	Request   *SIPMessage
	BranchID  string
	Method    string
	State     TransactionState
	Transport Transport
	RemoteAddr string
	Responses []*SIPMessage
	TimerA    *time.Timer
	TimerB    *time.Timer
	TimerD    *time.Timer
	CreatedAt time.Time
	mu        sync.RWMutex
	logger    *zap.Logger
}

type TransactionManager struct {
	transactions map[string]*Transaction
	mu           sync.RWMutex
	logger       *zap.Logger
}

func NewTransactionManager(logger *zap.Logger) *TransactionManager {
	return &TransactionManager{
		transactions: make(map[string]*Transaction),
		logger:       logger,
	}
}

func (tm *TransactionManager) CreateClientTransaction(request *SIPMessage, transport Transport, remoteAddr string) (*Transaction, error) {
	branchID := request.GetHeader("Via")
	if branchID == "" {
		return nil, fmt.Errorf("missing Via header")
	}

	txnID := generateTransactionID(branchID, request.RequestLine.Method)

	txn := &Transaction{
		ID:         txnID,
		Request:    request,
		BranchID:   branchID,
		Method:     request.RequestLine.Method,
		State:      TransactionTrying,
		Transport:  transport,
		RemoteAddr: remoteAddr,
		CreatedAt:  time.Now(),
		logger:     tm.logger,
	}

	tm.mu.Lock()
	tm.transactions[txnID] = txn
	tm.mu.Unlock()

	tm.logger.Debug("Transaction created",
		zap.String("id", txnID),
		zap.String("method", request.RequestLine.Method),
		zap.String("remote", remoteAddr),
	)

	return txn, nil
}

func (tm *TransactionManager) CreateServerTransaction(request *SIPMessage, transport Transport, remoteAddr string) (*Transaction, error) {
	branchID := request.GetHeader("Via")
	if branchID == "" {
		return nil, fmt.Errorf("missing Via header")
	}

	txnID := generateTransactionID(branchID, request.RequestLine.Method)

	txn := &Transaction{
		ID:         txnID,
		Request:    request,
		BranchID:   branchID,
		Method:     request.RequestLine.Method,
		State:      TransactionProceeding,
		Transport:  transport,
		RemoteAddr: remoteAddr,
		CreatedAt:  time.Now(),
		logger:     tm.logger,
	}

	tm.mu.Lock()
	tm.transactions[txnID] = txn
	tm.mu.Unlock()

	tm.logger.Debug("Server transaction created",
		zap.String("id", txnID),
		zap.String("method", request.RequestLine.Method),
		zap.String("remote", remoteAddr),
	)

	return txn, nil
}

func (tm *TransactionManager) GetTransaction(id string) (*Transaction, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	txn, ok := tm.transactions[id]
	return txn, ok
}

func (tm *TransactionManager) RemoveTransaction(id string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	delete(tm.transactions, id)
}

func (tm *TransactionManager) FindByBranch(branchID string) *Transaction {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	for _, txn := range tm.transactions {
		if txn.BranchID == branchID {
			return txn
		}
	}
	return nil
}

func generateTransactionID(branchID, method string) string {
	return fmt.Sprintf("%s-%s", branchID, method)
}

func (t *Transaction) SendRequest() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	data := []byte(t.Request.String())
	err := t.Transport.Send(t.RemoteAddr, data)
	if err != nil {
		return fmt.Errorf("send request failed: %w", err)
	}

	t.State = TransactionTrying
	t.startTimers()

	return nil
}

func (t *Transaction) SendResponse(response *SIPMessage) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	data := []byte(response.String())
	err := t.Transport.Send(t.RemoteAddr, data)
	if err != nil {
		return fmt.Errorf("send response failed: %w", err)
	}

	t.Responses = append(t.Responses, response)

	if response.StatusLine != nil {
		if response.StatusLine.StatusCode >= 200 {
			t.State = TransactionCompleted
		} else {
			t.State = TransactionProceeding
		}
	}

	return nil
}

func (t *Transaction) HandleResponse(response *SIPMessage) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Responses = append(t.Responses, response)

	if response.StatusLine != nil {
		if response.StatusLine.StatusCode >= 200 {
			t.State = TransactionCompleted
		} else if response.StatusLine.StatusCode >= 100 {
			t.State = TransactionProceeding
		}
	}

	t.logger.Debug("Response received",
		zap.String("id", t.ID),
		zap.Int("status", response.StatusLine.StatusCode),
	)
}

func (t *Transaction) startTimers() {
	t.TimerA = time.AfterFunc(time.Second, func() {
		t.mu.Lock()
		defer t.mu.Unlock()
		if t.State == TransactionTrying {
			t.Transport.Send(t.RemoteAddr, []byte(t.Request.String()))
			t.TimerA.Reset(2 * time.Second)
		}
	})

	t.TimerB = time.AfterFunc(64*time.Second, func() {
		t.mu.Lock()
		defer t.mu.Unlock()
		t.State = TransactionTerminated
		t.logger.Warn("Transaction timeout", zap.String("id", t.ID))
	})

	t.TimerD = time.AfterFunc(32*time.Second, func() {
		t.mu.Lock()
		defer t.mu.Unlock()
		t.State = TransactionTerminated
	})
}

func (t *Transaction) StopTimers() {
	if t.TimerA != nil {
		t.TimerA.Stop()
	}
	if t.TimerB != nil {
		t.TimerB.Stop()
	}
	if t.TimerD != nil {
		t.TimerD.Stop()
	}
}

func (t *Transaction) IsTerminated() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.State == TransactionTerminated || t.State == TransactionCompleted
}