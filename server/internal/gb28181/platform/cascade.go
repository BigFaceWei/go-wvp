package platform

import (
	"fmt"
	"sync"
	"time"

	"wvp-go/server/global"
	gbsip "wvp-go/server/internal/gb28181/sip"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
)

type CascadeManager struct {
	sipServer      *gbsip.Server
	sipClients     map[string]*gbsip.Client
	deviceChannels map[string][]system.DeviceChannel
	mu             sync.RWMutex
	logger         *zap.Logger
	running        bool
}

func NewCascadeManager(sipServer *gbsip.Server, logger *zap.Logger) *CascadeManager {
	return &CascadeManager{
		sipServer:      sipServer,
		sipClients:     make(map[string]*gbsip.Client),
		deviceChannels: make(map[string][]system.DeviceChannel),
		logger:         logger,
	}
}

func (cm *CascadeManager) Start() error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.running = true
	cm.logger.Info("Platform cascade manager started")
	return nil
}

func (cm *CascadeManager) Stop() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for platformID, client := range cm.sipClients {
		if err := client.Stop(); err != nil {
			cm.logger.Error("stop SIP client failed",
				zap.String("platform_id", platformID),
				zap.Error(err),
			)
		}
	}

	cm.running = false
	cm.logger.Info("Platform cascade manager stopped")
}

func (cm *CascadeManager) RegisterToUpperPlatform(platform system.Platform) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	clientConfig := &gbsip.ClientConfig{
		LocalIP:    global.GVA_CONFIG.WVP.SIP.ListenIP,
		LocalPort:  global.GVA_CONFIG.WVP.SIP.ListenPort,
		Domain:     platform.Domain,
		ServerID:   platform.PlatformID,
		RemoteIP:   platform.IP,
		RemotePort: platform.Port,
		Transport:  platform.Transport,
	}

	client := gbsip.NewClient(clientConfig, cm.logger)
	if err := client.Start(); err != nil {
		return fmt.Errorf("start SIP client for platform %s failed: %w", platform.PlatformID, err)
	}

	expires := platform.Expires
	if expires <= 0 {
		expires = 3600
	}

	txn, err := client.SendRegister(expires)
	if err != nil {
		return fmt.Errorf("register to platform %s failed: %w", platform.PlatformID, err)
	}

	cm.sipClients[platform.PlatformID] = client

	global.GVA_DB.Model(&system.Platform{}).Where("platform_id = ?", platform.PlatformID).Updates(map[string]interface{}{
		"status":        1,
		"register_time": time.Now(),
	})

	cm.logger.Info("Registered to upper platform",
		zap.String("platform_id", platform.PlatformID),
		zap.String("platform_ip", platform.IP),
		zap.Int("platform_port", platform.Port),
	)

	_ = txn
	return nil
}

func (cm *CascadeManager) UnregisterFromPlatform(platformID string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	client, ok := cm.sipClients[platformID]
	if !ok {
		return fmt.Errorf("platform %s not connected", platformID)
	}

	txn, err := client.SendRegister(0)
	if err != nil {
		return fmt.Errorf("unregister from platform %s failed: %w", platformID, err)
	}

	if err := client.Stop(); err != nil {
		cm.logger.Error("stop SIP client failed",
			zap.String("platform_id", platformID),
			zap.Error(err),
		)
	}

	delete(cm.sipClients, platformID)

	global.GVA_DB.Model(&system.Platform{}).Where("platform_id = ?", platformID).Update("status", 0)

	cm.logger.Info("Unregistered from upper platform", zap.String("platform_id", platformID))

	_ = txn
	return nil
}

func (cm *CascadeManager) SyncCatalog(platformID string) error {
	cm.mu.RLock()
	client, ok := cm.sipClients[platformID]
	cm.mu.RUnlock()

	if !ok {
		return fmt.Errorf("platform %s not connected", platformID)
	}

	var platform system.Platform
	if err := global.GVA_DB.Where("platform_id = ?", platformID).First(&platform).Error; err != nil {
		return fmt.Errorf("platform %s not found: %w", platformID, err)
	}

	var channels []system.DeviceChannel
	if err := global.GVA_DB.Find(&channels).Error; err != nil {
		return fmt.Errorf("query channels failed: %w", err)
	}

	cm.logger.Info("Syncing catalog to platform",
		zap.String("platform_id", platformID),
		zap.Int("channel_count", len(channels)),
	)

	for _, channel := range channels {
		cm.mu.Lock()
		cm.deviceChannels[channel.DeviceID] = append(cm.deviceChannels[channel.DeviceID], channel)
		cm.mu.Unlock()

		_ = client
	}

	global.GVA_DB.Model(&system.Platform{}).Where("platform_id = ?", platformID).Update("channel_count", len(channels))

	cm.logger.Info("Catalog synced to platform",
		zap.String("platform_id", platformID),
		zap.Int("channel_count", len(channels)),
	)

	return nil
}

func (cm *CascadeManager) HandleDeviceCatalog(deviceID string, channels []system.DeviceChannel) {
	cm.mu.Lock()
	cm.deviceChannels[deviceID] = channels
	cm.mu.Unlock()

	cm.logger.Info("Device catalog updated",
		zap.String("device_id", deviceID),
		zap.Int("channel_count", len(channels)),
	)
}

func (cm *CascadeManager) GetPlatformByDeviceID(deviceID string) *system.Platform {
	var channel system.DeviceChannel
	if err := global.GVA_DB.Where("channel_id = ?", deviceID).First(&channel).Error; err != nil {
		return nil
	}

	var platformChannel system.PlatformChannel
	if err := global.GVA_DB.Where("channel_id = ?", channel.ChannelID).First(&platformChannel).Error; err != nil {
		return nil
	}

	var platform system.Platform
	if err := global.GVA_DB.Where("platform_id = ?", platformChannel.PlatformID).First(&platform).Error; err != nil {
		return nil
	}

	return &platform
}

func (cm *CascadeManager) GetConnectedPlatforms() []system.Platform {
	var platforms []system.Platform
	global.GVA_DB.Where("status = ?", 1).Find(&platforms)
	return platforms
}

func (cm *CascadeManager) KeepaliveToUpperPlatforms() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Copy the client map under read lock so we don't hold lock during network I/O
		cm.mu.RLock()
		clients := make(map[string]*gbsip.Client, len(cm.sipClients))
		for id, client := range cm.sipClients {
			clients[id] = client
		}
		cm.mu.RUnlock()

		for platformID, client := range clients {
			_, err := client.SendKeepalive(platformID)
			if err != nil {
				cm.logger.Error("send keepalive to platform failed",
					zap.String("platform_id", platformID),
					zap.Error(err),
				)
			}
		}
	}
}
