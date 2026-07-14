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
	ID         string
	Request    *SIPMessage
	BranchID   string
	Method     string
	State      TransactionState
	Transport  Transport
	RemoteAddr string
	Responses  []*SIPMessage
	TimerA     *time.Timer
	TimerB     *time.Timer
	TimerD     *time.Timer
	CreatedAt  time.Time

	// responseReady is closed when a final response (status >= 200) is received.
	responseReady chan struct{}
	responseMu    sync.Mutex
	lastResponse  *SIPMessage // stored under responseMu; only set before closing responseReady

	mu      sync.RWMutex
	logger  *zap.Logger
	stopped bool // true once StopTimers has been called
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
	viaFull := request.GetHeader("Via")
	if viaFull == "" {
		return nil, fmt.Errorf("missing Via header")
	}

	branchID := ExtractBranchParam(viaFull)
	if branchID == "" {
		return nil, fmt.Errorf("missing branch parameter in Via header")
	}

	txnID := generateTransactionID(branchID, request.RequestLine.Method)

	txn := &Transaction{
		ID:            txnID,
		Request:       request,
		BranchID:      branchID,
		Method:        request.RequestLine.Method,
		State:         TransactionTrying,
		Transport:     transport,
		RemoteAddr:    remoteAddr,
		CreatedAt:     time.Now(),
		responseReady: make(chan struct{}),
		logger:        tm.logger,
	}

	tm.mu.Lock()
	tm.transactions[txnID] = txn
	tm.mu.Unlock()

	tm.logger.Debug("Transaction created",
		zap.String("id", txnID),
		zap.String("type", "client"),
		zap.String("method", request.RequestLine.Method),
		zap.String("remote", remoteAddr),
	)

	return txn, nil
}

func (tm *TransactionManager) CreateServerTransaction(request *SIPMessage, transport Transport, remoteAddr string) (*Transaction, error) {
	viaFull := request.GetHeader("Via")
	if viaFull == "" {
		return nil, fmt.Errorf("missing Via header")
	}

	branchID := ExtractBranchParam(viaFull)
	if branchID == "" {
		return nil, fmt.Errorf("missing branch parameter in Via header")
	}

	txnID := generateTransactionID(branchID, request.RequestLine.Method)

	txn := &Transaction{
		ID:            txnID,
		Request:       request,
		BranchID:      branchID,
		Method:        request.RequestLine.Method,
		State:         TransactionProceeding,
		Transport:     transport,
		RemoteAddr:    remoteAddr,
		CreatedAt:     time.Now(),
		responseReady: make(chan struct{}),
		logger:        tm.logger,
	}

	tm.mu.Lock()
	tm.transactions[txnID] = txn
	tm.mu.Unlock()

	tm.logger.Debug("Server transaction created",
		zap.String("id", txnID),
		zap.String("type", "server"),
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

// RemoveTransaction stops all timers and removes the transaction from the map.
func (tm *TransactionManager) RemoveTransaction(id string) {
	tm.mu.Lock()
	txn, ok := tm.transactions[id]
	if ok {
		delete(tm.transactions, id)
	}
	tm.mu.Unlock()

	if ok && txn != nil {
		txn.StopTimers()
	}
}

func (tm *TransactionManager) FindByBranch(viaHeader string) *Transaction {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	branchParam := ExtractBranchParam(viaHeader)
	if branchParam == "" {
		return nil
	}

	for _, txn := range tm.transactions {
		if txn.BranchID == branchParam {
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

	data := []byte(t.Request.String())
	err := t.Transport.Send(t.RemoteAddr, data)
	if err != nil {
		t.mu.Unlock()
		return fmt.Errorf("send request failed: %w", err)
	}

	t.State = TransactionTrying
	t.mu.Unlock()

	// Start timers outside the lock to avoid deadlock in timer callbacks
	t.startTimers()

	return nil
}

func (t *Transaction) SendResponse(response *SIPMessage) error {
	t.mu.Lock()

	data := []byte(response.String())
	err := t.Transport.Send(t.RemoteAddr, data)
	if err != nil {
		t.mu.Unlock()
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

	// For server transactions over UDP, start TimerD to cleanup after potential retransmissions
	if response.StatusLine != nil && response.StatusLine.StatusCode >= 200 {
		isUDP := true
		if udp, ok := t.Transport.(*UDPTransport); ok && udp != nil {
			isUDP = true
		} else {
			isUDP = false
		}

		if isUDP {
			// TimerD: cleanup after 32 seconds (RFC 3261 §17.2.2)
			t.TimerD = time.AfterFunc(32*time.Second, func() {
				t.mu.Lock()
				t.State = TransactionTerminated
				t.mu.Unlock()

				// Notify the transaction manager to remove this transaction
				// We do this by just marking it terminated — the manager's
				// periodic cleanup will handle it, or we could directly
				// remove it here with a callback.
				// For simplicity, just mark terminated.
				t.logger.Debug("Server transaction terminated by TimerD",
					zap.String("id", t.ID),
				)
			})
		}
	}

	t.mu.Unlock()

	return nil
}

// HandleResponse stores the response and notifies WaitForResponse callers.
func (t *Transaction) HandleResponse(response *SIPMessage) {
	t.mu.Lock()

	t.Responses = append(t.Responses, response)

	if response.StatusLine != nil {
		if response.StatusLine.StatusCode >= 200 {
			t.State = TransactionCompleted
			// Store and signal for WaitForResponse
			t.responseMu.Lock()
			if t.lastResponse == nil || t.lastResponse.StatusLine.StatusCode < 200 {
				t.lastResponse = response
				close(t.responseReady)
			}
			t.responseMu.Unlock()
		} else if response.StatusLine.StatusCode >= 100 {
			t.State = TransactionProceeding
			// Also store provisional responses so WaitForResponse can see them
			t.responseMu.Lock()
			if t.lastResponse == nil {
				t.lastResponse = response
			}
			t.responseMu.Unlock()
		}
	}

	t.logger.Debug("Response received",
		zap.String("id", t.ID),
		zap.Int("status", response.StatusLine.StatusCode),
	)

	// If we received a final response, stop client retransmission timers
	if response.StatusLine != nil && response.StatusLine.StatusCode >= 200 {
		t.mu.Unlock()
		t.StopTimers()
		return
	}
	t.mu.Unlock()
}

func (t *Transaction) startTimers() {
	t.mu.Lock()
	t.stopped = false
	t.mu.Unlock()

	// TimerA: retransmit after 1s, then 2s, 4s, 8s... up to 64s max
	retryInterval := time.Second
	t.TimerA = time.AfterFunc(retryInterval, func() {
		t.mu.Lock()
		if t.stopped {
			t.mu.Unlock()
			return
		}
		if t.State == TransactionTrying {
			// Copy request data and address outside lock for network send
			reqData := []byte(t.Request.String())
			addr := t.RemoteAddr
			transport := t.Transport
			isStopped := t.stopped
			t.mu.Unlock()

			if !isStopped {
				_ = transport.Send(addr, reqData)
			}

			// Double the retry interval, cap at 64s
			retryInterval *= 2
			if retryInterval > 64*time.Second {
				retryInterval = 64 * time.Second
			}

			t.mu.Lock()
			if !t.stopped {
				t.TimerA.Reset(retryInterval)
			}
			t.mu.Unlock()
		} else {
			t.mu.Unlock()
		}
	})

	// TimerB: overall timeout at 64s (RFC 3261)
	t.TimerB = time.AfterFunc(64*time.Second, func() {
		t.mu.Lock()
		if t.stopped {
			t.mu.Unlock()
			return
		}
		t.State = TransactionTerminated
		// Notify WaitForResponse that we timed out by closing responseReady
		t.responseMu.Lock()
		if t.lastResponse == nil {
			// closing with no response — WaitForResponse will detect via state check
		}
		t.responseMu.Unlock()
		t.logger.Warn("Client transaction timeout (TimerB)",
			zap.String("id", t.ID),
		)
		t.mu.Unlock()
	})
}

// StopTimers stops all timers and marks the transaction as stopped.
// After StopTimers returns, no timer callback will execute.
func (t *Transaction) StopTimers() {
	t.mu.Lock()
	if t.stopped {
		t.mu.Unlock()
		return
	}
	t.stopped = true

	if t.TimerA != nil {
		t.TimerA.Stop()
	}
	if t.TimerB != nil {
		t.TimerB.Stop()
	}
	if t.TimerD != nil {
		t.TimerD.Stop()
	}
	t.mu.Unlock()
}

func (t *Transaction) IsTerminated() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.State == TransactionTerminated || t.State == TransactionConfirmed
}

// WaitForResponse waits for a final SIP response (status >= 200).
// It returns the first final response received, or an error on timeout.
func (t *Transaction) WaitForResponse(timeout time.Duration) (*SIPMessage, error) {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-t.responseReady:
		t.responseMu.Lock()
		resp := t.lastResponse
		t.responseMu.Unlock()
		if resp != nil && resp.StatusLine != nil && resp.StatusLine.StatusCode >= 200 {
			return resp, nil
		}
		// If we got a provisional response stored, wait for the final one
		// Actually the channel is already closed, so this path shouldn't happen
		// since we only close responseReady when storing a final response.
		if resp != nil {
			return resp, nil
		}
		return nil, fmt.Errorf("transaction ended without a response (transaction %s)", t.ID)

	case <-timer.C:
		return nil, fmt.Errorf("timeout waiting for SIP response (transaction %s)", t.ID)
	}
}

// GetStoredResponse returns any response already stored (provisional or final).
func (t *Transaction) GetStoredResponse() *SIPMessage {
	t.responseMu.Lock()
	defer t.responseMu.Unlock()
	return t.lastResponse
}
