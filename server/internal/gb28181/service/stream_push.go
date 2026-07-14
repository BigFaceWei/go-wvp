package service

import (
	"fmt"

	"wvp-go/server/global"
	"wvp-go/server/internal/media/zlm"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
)

type StreamPushService struct {
	logger *zap.Logger
}

func NewStreamPushService(logger *zap.Logger) *StreamPushService {
	if logger == nil {
		logger = global.GVA_LOG
	}
	return &StreamPushService{logger: logger}
}

func (s *StreamPushService) getZLMClient() *zlm.Client {
	if len(global.GVA_CONFIG.WVP.ZLMediaKit.Servers) == 0 {
		return nil
	}
	for _, cfg := range global.GVA_CONFIG.WVP.ZLMediaKit.Servers {
		if cfg.Default {
			return zlm.NewClient(&zlm.ZLMConfig{
				ID:       cfg.ID,
				IP:       cfg.IP,
				HTTPPort: cfg.HTTPPort,
				Secret:   cfg.Secret,
				Default:  cfg.Default,
			}, s.logger)
		}
	}
	cfg := global.GVA_CONFIG.WVP.ZLMediaKit.Servers[0]
	return zlm.NewClient(&zlm.ZLMConfig{
		ID:       cfg.ID,
		IP:       cfg.IP,
		HTTPPort: cfg.HTTPPort,
		Secret:   cfg.Secret,
		Default:  cfg.Default,
	}, s.logger)
}

func (s *StreamPushService) GetList(name, stream string, page, pageSize int) ([]system.StreamPush, int64, error) {
	var pushes []system.StreamPush
	db := global.GVA_DB.Model(&system.StreamPush{})
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if stream != "" {
		db = db.Where("stream LIKE ?", "%"+stream+"%")
	}
	var total int64
	db.Count(&total)
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&pushes).Error; err != nil {
		return nil, 0, err
	}
	return pushes, total, nil
}

func (s *StreamPushService) Get(id string) (*system.StreamPush, error) {
	var push system.StreamPush
	if err := global.GVA_DB.Where("id = ?", id).First(&push).Error; err != nil {
		return nil, err
	}
	return &push, nil
}

func (s *StreamPushService) Create(push *system.StreamPush) error {
	if push.App == "" {
		push.App = "push"
	}
	if push.Stream == "" {
		push.Stream = fmt.Sprintf("push_%s", push.Name)
	}
	push.Status = false
	push.Online = false
	return global.GVA_DB.Create(push).Error
}

func (s *StreamPushService) Update(push *system.StreamPush) error {
	return global.GVA_DB.Save(push).Error
}

func (s *StreamPushService) Delete(id string) error {
	_ = s.Stop(id)
	return global.GVA_DB.Delete(&system.StreamPush{}, "id = ?", id).Error
}

func (s *StreamPushService) Start(id string) error {
	push, err := s.Get(id)
	if err != nil {
		return fmt.Errorf("push not found: %w", err)
	}
	if push.Status {
		return fmt.Errorf("push already running")
	}

	client := s.getZLMClient()
	if client == nil {
		return fmt.Errorf("no ZLMediaKit server configured")
	}

	streamCfg := global.GVA_CONFIG.WVP.ZLMediaKit.Stream
	vhost := streamCfg.Vhost
	if vhost == "" {
		vhost = "__defaultVhost__"
	}

	key, err := client.AddStreamPusherProxy(vhost, push.App, push.Stream, push.PushURL)
	if err != nil {
		return fmt.Errorf("start stream pusher failed: %w", err)
	}

	push.Status = true
	push.StreamKey = key
	push.MediaServerID = client.GetID()
	return global.GVA_DB.Save(push).Error
}

func (s *StreamPushService) Stop(id string) error {
	push, err := s.Get(id)
	if err != nil {
		return fmt.Errorf("push not found: %w", err)
	}
	if !push.Status {
		return nil
	}

	client := s.getZLMClient()
	if client != nil {
		streamCfg := global.GVA_CONFIG.WVP.ZLMediaKit.Stream
		vhost := streamCfg.Vhost
		if vhost == "" {
			vhost = "__defaultVhost__"
		}
		_ = client.DelStreamPusherProxy(vhost, push.App, push.Stream)
	}

	push.Status = false
	push.Online = false
	push.StreamKey = ""
	return global.GVA_DB.Save(push).Error
}
