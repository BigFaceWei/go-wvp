package session

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

type InviteSessionState int

const (
InviteSessionInit InviteSessionState = iota
InviteSessionCalling
InviteSessionRinging
InviteSessionAnswered
InviteSessionTerminated
)

type InviteSession struct {
	ID          string
	CallID      string
	FromTag     string
	ToTag       string
	DeviceID    string
	ChannelID   string
	SSRC        string
	MediaIP     string
	MediaPort   int
	State       InviteSessionState
	StartTime   time.Time
	EndTime     *time.Time
	CSeq        int
	mu          sync.RWMutex
	logger      *zap.Logger
}

type SessionManager struct {
	sessions map[string]*InviteSession
	mu       sync.RWMutex
	logger   *zap.Logger
}

func NewSessionManager(logger *zap.Logger) *SessionManager {
return &SessionManager{
sessions: make(map[string]*InviteSession),
logger: logger,
}
}

func (m *SessionManager) CreateSession(deviceID, channelID, ssrc, mediaIP string, mediaPort int) *InviteSession {
m.mu.Lock()
	defer m.mu.Unlock()

sessionID := fmt.Sprintf("%s-%s-%d", deviceID, channelID, time.Now().UnixNano())

session := &InviteSession{
ID:          sessionID,
DeviceID:    deviceID,
ChannelID:   channelID,
SSRC:        ssrc,
MediaIP:     mediaIP,
MediaPort:   mediaPort,
State:       InviteSessionInit,
StartTime:   time.Now(),
CSeq:        1,
logger:      m.logger,
}

m.sessions[sessionID] = session
return session
}

func (m *SessionManager) GetSession(id string) (*InviteSession, bool) {
m.mu.RLock()
	defer m.mu.RUnlock()
	session, ok := m.sessions[id]
return session, ok
}

func (m *SessionManager) GetSessionByDeviceID(deviceID string) *InviteSession {
m.mu.RLock()
	defer m.mu.RUnlock()

for _, session := range m.sessions {
if session.DeviceID == deviceID && session.State != InviteSessionTerminated {
return session
}
	}
return nil
}

func (m *SessionManager) RemoveSession(id string) {
m.mu.Lock()
	defer m.mu.Unlock()
delete(m.sessions, id)
}

func (m *SessionManager) GetSessionCount() int {
m.mu.RLock()
	defer m.mu.RUnlock()
return len(m.sessions)
}

func (s *InviteSession) UpdateState(state InviteSessionState) {
s.mu.Lock()
	defer s.mu.Unlock()
s.State = state

if state == InviteSessionTerminated {
now := time.Now()
s.EndTime = &now
	}
}

func (s *InviteSession) GetState() InviteSessionState {
s.mu.RLock()
	defer s.mu.RUnlock()
return s.State
}

func (s *InviteSession) IsTerminated() bool {
s.mu.RLock()
	defer s.mu.RUnlock()
return s.State == InviteSessionTerminated
}

func (s *InviteSession) GetDuration() time.Duration {
s.mu.RLock()
	defer s.mu.RUnlock()

if s.EndTime != nil {
return s.EndTime.Sub(s.StartTime)
	}
return time.Since(s.StartTime)
}