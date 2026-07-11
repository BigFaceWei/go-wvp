package zlm

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

type StreamManager struct {
	client  *Client
	streams map[string]*StreamSession
	mu      sync.RWMutex
	logger  *zap.Logger
}

type StreamSession struct {
	Key        string
	Schema     string
	VHost      string
	App        string
	StreamID   string
	DeviceID   string
	ChannelID  string
	StartTime  time.Time
	MediaInfo  *StreamInfo
}

func NewStreamManager(client *Client, logger *zap.Logger) *StreamManager {
	return &StreamManager{
		client:  client,
		streams: make(map[string]*StreamSession),
		logger:  logger,
	}
}

func (m *StreamManager) buildKey(schema, vhost, app, stream string) string {
	return fmt.Sprintf("%s/%s/%s/%s", schema, vhost, app, stream)
}

func (m *StreamManager) AddStream(schema, vhost, app, stream, deviceID, channelID string) *StreamSession {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := m.buildKey(schema, vhost, app, stream)
	session := &StreamSession{
		Key:       key,
		Schema:    schema,
		VHost:     vhost,
		App:       app,
		StreamID:  stream,
		DeviceID:  deviceID,
		ChannelID: channelID,
		StartTime: time.Now(),
	}
	m.streams[key] = session

	m.logger.Info("Stream added",
		zap.String("key", key),
		zap.String("device_id", deviceID),
		zap.String("channel_id", channelID),
	)

	return session
}

func (m *StreamManager) RemoveStream(schema, vhost, app, stream string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := m.buildKey(schema, vhost, app, stream)
	delete(m.streams, key)

	m.logger.Info("Stream removed", zap.String("key", key))
}

func (m *StreamManager) GetStream(schema, vhost, app, stream string) (*StreamSession, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	key := m.buildKey(schema, vhost, app, stream)
	session, ok := m.streams[key]
	return session, ok
}

func (m *StreamManager) GetStreamByDeviceID(deviceID string) []*StreamSession {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var sessions []*StreamSession
	for _, session := range m.streams {
		if session.DeviceID == deviceID {
			sessions = append(sessions, session)
		}
	}
	return sessions
}

func (m *StreamManager) GetAllStreams() []*StreamSession {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var sessions []*StreamSession
	for _, session := range m.streams {
		sessions = append(sessions, session)
	}
	return sessions
}

func (m *StreamManager) BuildPlayURL(schema, app, stream string) string {
	return fmt.Sprintf("%s://%s/%s/%s", schema, m.client.GetAPI(), app, stream)
}

func (m *StreamManager) BuildWebRTCURL(app, stream string) string {
	return fmt.Sprintf("webrtc://%s/%s/%s", m.client.GetAPI(), app, stream)
}

func (m *StreamManager) GetStreamCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.streams)
}

func (m *StreamManager) StopStream(schema, vhost, app, stream string) error {
	key := m.buildKey(schema, vhost, app, stream)

	if err := m.client.CloseStream(schema, vhost, app, stream); err != nil {
		m.logger.Error("Close stream on ZLM failed",
			zap.String("key", key),
			zap.Error(err),
		)
		return err
	}

	m.RemoveStream(schema, vhost, app, stream)
	return nil
}
