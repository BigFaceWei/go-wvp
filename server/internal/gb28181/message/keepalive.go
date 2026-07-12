package message

import (
	"fmt"
	"time"

	"wvp-go/server/global"
	"wvp-go/server/model/system"
	"wvp-go/server/pkg/sip"

	"go.uber.org/zap"
)

type KeepaliveHandler struct {
	logger *zap.Logger
}

func NewKeepaliveHandler(logger *zap.Logger) *KeepaliveHandler {
	return &KeepaliveHandler{logger: logger}
}

type KeepaliveRequest struct {
	DeviceID   string
	Domain     string
	RemoteAddr string
}

func (h *KeepaliveHandler) ParseKeepalive(msg *sip.SIPMessage) (*KeepaliveRequest, error) {
	if !msg.IsRequest || msg.RequestLine.Method != "MESSAGE" {
		return nil, fmt.Errorf("not a MESSAGE message")
	}

	req := &KeepaliveRequest{}

	from := msg.GetHeader("From")
	if from != "" {
		parts := []byte(from)
		start := -1
		for i, b := range parts {
			if b == '<' {
				start = i + 1
			} else if b == '>' && start != -1 {
				uri := string(parts[start:i])
				if len(uri) > 4 && uri[:4] == "sip:" {
					uri = uri[4:]
				}
				if idx := indexOf(uri, '@'); idx != -1 {
					req.DeviceID = uri[:idx]
					req.Domain = uri[idx+1:]
				}
				break
			}
		}
	}

	if contact := msg.GetHeader("Contact"); contact != "" {
		parts := []byte(contact)
		start := -1
		for i, b := range parts {
			if b == '<' {
				start = i + 1
			} else if b == '>' && start != -1 {
				uri := string(parts[start:i])
				if len(uri) > 4 && uri[:4] == "sip:" {
					uri = uri[4:]
				}
				if idx := indexOf(uri, '@'); idx != -1 {
					rest := uri[idx+1:]
					if endIdx := indexOf(rest, '>'); endIdx != -1 {
						rest = rest[:endIdx]
					}
					req.RemoteAddr = rest
				}
				break
			}
		}
	}

	return req, nil
}

func indexOf(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func (h *KeepaliveHandler) HandleKeepalive(req *KeepaliveRequest) error {
	device := &system.Device{}
	result := global.GVA_DB.Where("device_id = ?", req.DeviceID).First(device)
	if result.Error != nil {
		h.logger.Warn("Keepalive from unknown device",
			zap.String("device_id", req.DeviceID),
			zap.String("ip", req.RemoteAddr),
		)
		return fmt.Errorf("device not found: %s", req.DeviceID)
	}

	device.Online = true
	device.KeepaliveTime = time.Now()
	device.IP = req.RemoteAddr

	if err := global.GVA_DB.Save(device).Error; err != nil {
		return fmt.Errorf("update keepalive failed: %w", err)
	}

	h.logger.Debug("Device keepalive",
		zap.String("device_id", req.DeviceID),
		zap.String("ip", req.RemoteAddr),
	)

	return nil
}

type DeviceStatusChecker struct {
	logger   *zap.Logger
	interval time.Duration
	stopCh   chan struct{}
}

func NewDeviceStatusChecker(logger *zap.Logger, interval time.Duration) *DeviceStatusChecker {
	return &DeviceStatusChecker{
		logger:   logger,
		interval: interval,
		stopCh:   make(chan struct{}),
	}
}

func (c *DeviceStatusChecker) Start() {
	go c.checkLoop()
}

func (c *DeviceStatusChecker) Stop() {
	close(c.stopCh)
}

func (c *DeviceStatusChecker) checkLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-c.stopCh:
			return
		case <-ticker.C:
			c.checkOfflineDevices()
		}
	}
}

func (c *DeviceStatusChecker) checkOfflineDevices() {
	timeout := time.Now().Add(-3 * time.Minute)
	result := global.GVA_DB.Model(&system.Device{}).
		Where("online = ? AND keepalive_time < ?", true, timeout).
		Update("online", false)

	if result.Error != nil {
		c.logger.Error("Check offline devices failed", zap.Error(result.Error))
		return
	}

	if result.RowsAffected > 0 {
		c.logger.Info("Devices went offline", zap.Int64("count", result.RowsAffected))
	}
}