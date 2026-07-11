package record

import (
	"fmt"
	"time"

	"wvp-go/server/global"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
)

type PlaybackService struct {
	logger *zap.Logger
}

func NewPlaybackService(logger *zap.Logger) *PlaybackService {
	return &PlaybackService{logger: logger}
}

type PlaybackRequest struct {
	DeviceID  string
	ChannelID string
	StartTime time.Time
	EndTime   time.Time
	Speed     float64
	SSRC      string
}

type PlaybackSession struct {
	SessionID   string
	DeviceID    string
	ChannelID   string
	StartTime   time.Time
	EndTime     time.Time
	Speed       float64
	PlaybackURL string
	Status      string
	CreateTime  time.Time
}

func (s *PlaybackService) StartPlayback(req *PlaybackRequest) (*PlaybackSession, error) {
	if req.StartTime.After(req.EndTime) {
		return nil, fmt.Errorf("start time must be before end time")
	}

	speed := req.Speed
	if speed <= 0 {
		speed = 1.0
	}
	if speed > 16 {
		speed = 16.0
	}

	var files []system.RecordFile
	if err := global.GVA_DB.Where("device_id = ? AND channel_id = ? AND start_time >= ? AND end_time <= ?",
		req.DeviceID, req.ChannelID, req.StartTime, req.EndTime).
		Order("start_time ASC").
		Find(&files).Error; err != nil {
		return nil, fmt.Errorf("query record files failed: %w", err)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no record files found in the specified time range")
	}

	session := &PlaybackSession{
		SessionID:   fmt.Sprintf("playback_%s_%d", req.DeviceID, time.Now().UnixNano()),
		DeviceID:    req.DeviceID,
		ChannelID:   req.ChannelID,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Speed:       speed,
		PlaybackURL: fmt.Sprintf("rtsp://localhost/record/%s/%s", req.DeviceID, req.ChannelID),
		Status:      "playing",
		CreateTime:  time.Now(),
	}

	s.logger.Info("Video playback started",
		zap.String("device_id", req.DeviceID),
		zap.Time("start", req.StartTime),
		zap.Time("end", req.EndTime),
		zap.Float64("speed", speed),
	)

	return session, nil
}

func (s *PlaybackService) StopPlayback(sessionID string) error {
	s.logger.Info("Video playback stopped", zap.String("session_id", sessionID))
	return nil
}

func (s *PlaybackService) PausePlayback(sessionID string) error {
	s.logger.Info("Video playback paused", zap.String("session_id", sessionID))
	return nil
}

func (s *PlaybackService) ResumePlayback(sessionID string) error {
	s.logger.Info("Video playback resumed", zap.String("session_id", sessionID))
	return nil
}

func (s *PlaybackService) SeekPlayback(sessionID string, seekTime time.Time) error {
	s.logger.Info("Video playback seek",
		zap.String("session_id", sessionID),
		zap.Time("seek_time", seekTime),
	)
	return nil
}

func (s *PlaybackService) DownloadRecordings(deviceID, channelID string, startTime, endTime time.Time) (string, error) {
	var files []system.RecordFile
	if err := global.GVA_DB.Where("device_id = ? AND channel_id = ? AND start_time >= ? AND start_time <= ?",
		deviceID, channelID, startTime, endTime).
		Find(&files).Error; err != nil {
		return "", fmt.Errorf("query record files failed: %w", err)
	}

	if len(files) == 0 {
		return "", fmt.Errorf("no record files found in the specified time range")
	}

	downloadURL := fmt.Sprintf("/api/record/file/download/%s/%s/%d", deviceID, channelID, startTime.Unix())
	return downloadURL, nil
}

func (s *PlaybackService) QueryRecordFiles(deviceID, channelID string, startTime, endTime time.Time, page, pageSize int) ([]system.RecordFile, int64, error) {
	db := global.GVA_DB.Model(&system.RecordFile{}).Where("start_time >= ? AND end_time <= ?", startTime, endTime)

	if deviceID != "" {
		db = db.Where("device_id = ?", deviceID)
	}
	if channelID != "" {
		db = db.Where("channel_id = ?", channelID)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var files []system.RecordFile
	if err := db.Order("start_time ASC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&files).Error; err != nil {
		return nil, 0, err
	}

	return files, total, nil
}
