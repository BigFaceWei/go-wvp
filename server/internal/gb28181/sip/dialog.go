package sip

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

type DialogState int

const (
DialogInit DialogState = iota
DialogConfirmed
DialogTerminated
)

type Dialog struct {
ID         string
CallID     string
LocalTag  string
RemoteTag string
LocalURI  string
RemoteURI string
RouteSet  []string
LocalCSeq int
RemoteCSeq int
State     DialogState
Transport Transport
RemoteAddr string
CreatedAt time.Time
ExpiresAt *time.Time
	mu        sync.RWMutex
	logger    *zap.Logger
}

type DialogManager struct {
dialogs map[string]*Dialog
	mu      sync.RWMutex
	logger  *zap.Logger
}

func NewDialogManager(logger *zap.Logger) *DialogManager {
return &DialogManager{
dialogs: make(map[string]*Dialog),
logger: logger,
}
}

func (dm *DialogManager) CreateDialog(callID, localTag, remoteTag, localURI, remoteURI string, transport Transport, remoteAddr string) (*Dialog, error) {
	dialogID := generateDialogID(callID, localTag, remoteTag)

	dm.mu.Lock()
	defer dm.mu.Unlock()

	if _, exists := dm.dialogs[dialogID]; exists {
return nil, fmt.Errorf("dialog already exists: %s", dialogID)
	}

	dialog := &Dialog{
ID:         dialogID,
CallID:     callID,
LocalTag:  localTag,
RemoteTag: remoteTag,
LocalURI:  localURI,
RemoteURI: remoteURI,
State:     DialogInit,
Transport: transport,
RemoteAddr: remoteAddr,
CreatedAt: time.Now(),
logger:    dm.logger,
	}

	dm.dialogs[dialogID] = dialog

	dm.logger.Debug("Dialog created",
zap.String("id", dialogID),
zap.String("call_id", callID),
	)

	return dialog, nil
}

func (dm *DialogManager) GetDialog(id string) (*Dialog, bool) {
dm.mu.RLock()
	defer dm.mu.RUnlock()
	dialog, ok := dm.dialogs[id]
return dialog, ok
}

func (dm *DialogManager) FindByCallID(callID string) *Dialog {
dm.mu.RLock()
	defer dm.mu.RUnlock()

	for _, dialog := range dm.dialogs {
if dialog.CallID == callID {
return dialog
}
	}
return nil
}

func (dm *DialogManager) RemoveDialog(id string) {
dm.mu.Lock()
	defer dm.mu.Unlock()
delete(dm.dialogs, id)
}

func generateDialogID(callID, localTag, remoteTag string) string {
return fmt.Sprintf("%s:%s:%s", callID, localTag, remoteTag)
}

func (d *Dialog) Update(request *SIPMessage) {
d.mu.Lock()
	defer d.mu.Unlock()

if request.IsRequest {
d.RemoteCSeq++
	} else {
d.LocalCSeq++
	}
}

func (d *Dialog) Terminate() {
d.mu.Lock()
	defer d.mu.Unlock()
d.State = DialogTerminated
}

func (d *Dialog) IsTerminated() bool {
d.mu.RLock()
	defer d.mu.RUnlock()
return d.State == DialogTerminated
}