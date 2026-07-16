package zlm

import (
	"wvp-go/server/global"

	"go.uber.org/zap"
)

// InitZLMHooks initializes the ZLMediaKit integration:
// 1. Creates a ZLM client for the default media server
// 2. Registers hook routes on the gin engine
// 3. Syncs ZLM configuration (setServerConfig with hook URLs)
// This must be called AFTER the HTTP server has started.
func InitZLMHooks() {
	logger := zap.L().Named("zlm-init")

	servers := global.GVA_CONFIG.WVP.ZLMediaKit.Servers
	if len(servers) == 0 {
		logger.Warn("No ZLMediaKit servers configured, skipping ZLM hook init")
		return
	}

	// Find default server
	var defaultServer *ZLMConfig
	for i, s := range servers {
		if s.Default {
			defaultServer = &ZLMConfig{
				ID:            s.ID,
				IP:            s.IP,
				HTTPPort:      s.HTTPPort,
				APIPort:       s.HTTPPort, // ZLM serves API on HTTP port
				Secret:        s.Secret,
				Default:       s.Default,
				HookIP:        global.GVA_CONFIG.WVP.SIP.ListenIP,
				HookPort:      global.GVA_CONFIG.System.Addr,
				MediaHTTPPort: s.HTTPPort,
			}
			break
		}
		_ = i
	}
	if defaultServer == nil {
		s := servers[0]
		defaultServer = &ZLMConfig{
			ID:            s.ID,
			IP:            s.IP,
			HTTPPort:      s.HTTPPort,
			APIPort:       s.HTTPPort,
			Secret:        s.Secret,
			Default:       s.Default,
			HookIP:        global.GVA_CONFIG.WVP.SIP.ListenIP,
			HookPort:      global.GVA_CONFIG.System.Addr,
			MediaHTTPPort: s.HTTPPort,
		}
	}

	logger.Info("Initializing ZLM integration",
		zap.String("zlm_ip", defaultServer.IP),
		zap.Int("zlm_http_port", defaultServer.HTTPPort),
		zap.String("hook_ip", defaultServer.HookIP),
		zap.Int("hook_port", defaultServer.HookPort),
	)

	client := NewClient(defaultServer, logger)
	handler := NewHookHandler(client, logger)
	handler.RegisterDefaultHandlers()

	// Store for use by other components
	global.GVA_ZLM_HANDLER = handler
	global.GVA_ZLM_CLIENT = client

	// Sync ZLM configuration (getServerConfig + setServerConfig)
	if err := handler.SyncZLMConfig(); err != nil {
		logger.Error("Failed to sync ZLM configuration", zap.Error(err))
	} else {
		logger.Info("ZLMediaKit configuration synced successfully")
	}
}

// GetHookHandler returns the global ZLM hook handler.
func GetHookHandler() *HookHandler {
	if global.GVA_ZLM_HANDLER != nil {
		return global.GVA_ZLM_HANDLER.(*HookHandler)
	}
	return nil
}

// GetZLMClient returns the global ZLM client.
func GetZLMClient() *Client {
	if global.GVA_ZLM_CLIENT != nil {
		return global.GVA_ZLM_CLIENT.(*Client)
	}
	return nil
}
