package service

import (
	"fmt"

	"wvp-go/server/global"
	"wvp-go/server/internal/media/zlm"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
)

type StreamProxyService struct {
	logger *zap.Logger
}

func NewStreamProxyService(logger *zap.Logger) *StreamProxyService {
	if logger == nil {
		logger = global.GVA_LOG
	}
	return &StreamProxyService{logger: logger}
}

// getZLMClient returns the first default ZLM client.
func (s *StreamProxyService) getZLMClient() *zlm.Client {
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
	// return first server if no default
	cfg := global.GVA_CONFIG.WVP.ZLMediaKit.Servers[0]
	return zlm.NewClient(&zlm.ZLMConfig{
		ID:       cfg.ID,
		IP:       cfg.IP,
		HTTPPort: cfg.HTTPPort,
		Secret:   cfg.Secret,
		Default:  cfg.Default,
	}, s.logger)
}

func (s *StreamProxyService) GetList(name, stream string, page, pageSize int) ([]system.StreamProxy, int64, error) {
	var proxies []system.StreamProxy
	db := global.GVA_DB.Model(&system.StreamProxy{})
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if stream != "" {
		db = db.Where("stream LIKE ?", "%"+stream+"%")
	}
	var total int64
	db.Count(&total)
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&proxies).Error; err != nil {
		return nil, 0, err
	}
	return proxies, total, nil
}

func (s *StreamProxyService) Get(id string) (*system.StreamProxy, error) {
	var proxy system.StreamProxy
	if err := global.GVA_DB.Where("id = ?", id).First(&proxy).Error; err != nil {
		return nil, err
	}
	return &proxy, nil
}

func (s *StreamProxyService) Create(proxy *system.StreamProxy) error {
	if proxy.App == "" {
		proxy.App = "proxy"
	}
	if proxy.Stream == "" {
		proxy.Stream = fmt.Sprintf("proxy_%s", proxy.Name)
	}
	proxy.Status = false
	proxy.Online = false
	return global.GVA_DB.Create(proxy).Error
}

func (s *StreamProxyService) Update(proxy *system.StreamProxy) error {
	return global.GVA_DB.Save(proxy).Error
}

func (s *StreamProxyService) Delete(id string) error {
	// Stop the proxy if it's running before deleting
	_ = s.Stop(id)
	return global.GVA_DB.Delete(&system.StreamProxy{}, "id = ?", id).Error
}

func (s *StreamProxyService) Start(id string) error {
	proxy, err := s.Get(id)
	if err != nil {
		return fmt.Errorf("proxy not found: %w", err)
	}
	if proxy.Status {
		return fmt.Errorf("proxy already running")
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
	app := proxy.App
	if app == "" {
		app = "proxy"
	}
	timeoutSec := "30"
	if proxy.TimeoutSec > 0 {
		timeoutSec = fmt.Sprintf("%d", proxy.TimeoutSec)
	}

	key, err := client.AddStreamProxy(vhost, app, proxy.Stream, proxy.URL, proxy.SrcURL, timeoutSec)
	if err != nil {
		return fmt.Errorf("start stream proxy failed: %w", err)
	}

	proxy.Status = true
	proxy.StreamKey = key
	proxy.MediaServerID = client.GetID()
	return global.GVA_DB.Save(proxy).Error
}

func (s *StreamProxyService) Stop(id string) error {
	proxy, err := s.Get(id)
	if err != nil {
		return fmt.Errorf("proxy not found: %w", err)
	}
	if !proxy.Status {
		return nil // already stopped
	}

	client := s.getZLMClient()
	if client != nil {
		streamCfg := global.GVA_CONFIG.WVP.ZLMediaKit.Stream
		vhost := streamCfg.Vhost
		if vhost == "" {
			vhost = "__defaultVhost__"
		}

		if proxy.StreamKey != "" {
			_ = client.DelFFmpegSource(proxy.StreamKey)
		}
		_ = client.DelStreamProxy(vhost, proxy.App, proxy.Stream)
	}

	proxy.Status = false
	proxy.Online = false
	proxy.StreamKey = ""
	return global.GVA_DB.Save(proxy).Error
}
