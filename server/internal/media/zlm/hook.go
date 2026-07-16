package zlm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"go.uber.org/zap"
)

// StopStreamFunc is called when a stream should be stopped (e.g., via on_stream_none_reader).
// deviceID and channelID are parsed from the stream_id (format: deviceId_channelId).
// Set by the service layer via OnStreamNoneReaderCallback.
type StopStreamFunc func(deviceID, channelID string) error

// OnStreamNoneReaderCallback is set by the gb28181 service package to handle
// on_stream_none_reader hooks with full cleanup (BYE + closeRtpServer + close_streams).
// The service package imports zlm, so assignment avoids circular imports.
var OnStreamNoneReaderCallback StopStreamFunc

type HookHandler struct {
	client   *Client
	handlers map[string]HookFunc
	mu       sync.RWMutex
	logger   *zap.Logger
}

type HookFunc func(params map[string]interface{}) error

type HookRegister interface {
	OnStreamChanged(params map[string]interface{}) error
	OnStreamNotFound(params map[string]interface{}) error
	OnPlayHook(params map[string]interface{}) error
	OnPublishHook(params map[string]interface{}) error
	OnRecordHook(params map[string]interface{}) error
}

func NewHookHandler(client *Client, logger *zap.Logger) *HookHandler {
	return &HookHandler{
		client:   client,
		handlers: make(map[string]HookFunc),
		logger:   logger,
	}
}

func (h *HookHandler) Register(hookName string, handler HookFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.handlers[hookName] = handler
}

func (h *HookHandler) HandleHook(w http.ResponseWriter, r *http.Request) {
	// Extract hook name from URL path: /index/hook/on_server_keepalive → on_server_keepalive
	path := r.URL.Path
	hookName := ""
	if idx := strings.LastIndex(path, "/"); idx >= 0 {
		hookName = path[idx+1:]
	}
	if hookName == "" {
		http.Error(w, "missing hook name in path", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "read body failed", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var params map[string]interface{}
	if len(body) > 0 {
		if err := json.Unmarshal(body, &params); err != nil {
			http.Error(w, "parse body failed", http.StatusBadRequest)
			return
		}
	}

	h.logger.Info("ZLM hook received",
		zap.String("hook", hookName),
		zap.Int("body_len", len(body)),
	)

	h.mu.RLock()
	handler, exists := h.handlers[hookName]
	h.mu.RUnlock()

	if !exists {
		h.logger.Warn("No handler for hook", zap.String("hook", hookName))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code": 0,
		})
		return
	}

	if err := handler(params); err != nil {
		h.logger.Error("Handle hook failed",
			zap.String("hook", hookName),
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 0,
	})
}

// ── Default hook handlers ──

func (h *HookHandler) DefaultPublishHookHandler(params map[string]interface{}) error {
	h.logger.Info("on_publish hook",
		zap.String("app", getStr(params, "app")),
		zap.String("stream", getStr(params, "stream")),
		zap.String("ip", getStr(params, "ip")),
		zap.Float64("port", getFloat(params, "port")),
	)
	return nil
}

func (h *HookHandler) DefaultStreamChangedHandler(params map[string]interface{}) error {
	register := getBool(params, "regist")
	if register {
		h.logger.Info("Stream registered",
			zap.String("schema", getStr(params, "schema")),
			zap.String("app", getStr(params, "app")),
			zap.String("stream", getStr(params, "stream")),
		)
	} else {
		h.logger.Info("Stream unregistered",
			zap.String("app", getStr(params, "app")),
			zap.String("stream", getStr(params, "stream")),
		)
	}
	return nil
}

func (h *HookHandler) DefaultStreamNotFoundHandler(params map[string]interface{}) error {
	h.logger.Warn("Stream not found",
		zap.String("app", getStr(params, "app")),
		zap.String("stream", getStr(params, "stream")),
	)
	return nil
}

func (h *HookHandler) DefaultPlayHookHandler(params map[string]interface{}) error {
	h.logger.Info("on_play hook",
		zap.String("app", getStr(params, "app")),
		zap.String("stream", getStr(params, "stream")),
		zap.String("ip", getStr(params, "ip")),
	)
	return nil
}

func (h *HookHandler) DefaultRecordHookHandler(params map[string]interface{}) error {
	h.logger.Info("on_record_mp4 hook",
		zap.String("app", getStr(params, "app")),
		zap.String("stream", getStr(params, "stream")),
	)
	return nil
}

func (h *HookHandler) DefaultStreamNoneReaderHandler(params map[string]interface{}) error {
	stream := getStr(params, "stream")

	h.logger.Info("on_stream_none_reader hook — auto-stopping stream",
		zap.String("app", getStr(params, "app")),
		zap.String("stream", stream),
	)

	// Parse deviceID and channelID from stream name (format: deviceId_channelId)
	parts := strings.SplitN(stream, "_", 2)
	if len(parts) != 2 {
		h.logger.Warn("on_stream_none_reader: cannot parse stream_id", zap.String("stream", stream))
		return nil
	}
	deviceID := parts[0]
	channelID := parts[1]

	if OnStreamNoneReaderCallback != nil {
		h.logger.Info("on_stream_none_reader: delegating to stop callback",
			zap.String("device_id", deviceID),
			zap.String("channel_id", channelID),
		)
		if err := OnStreamNoneReaderCallback(deviceID, channelID); err != nil {
			h.logger.Error("on_stream_none_reader: stop callback failed", zap.Error(err))
		}
	} else {
		h.logger.Warn("on_stream_none_reader: no stop callback registered, stream not cleaned up")
	}

	return nil
}

func (h *HookHandler) DefaultRTPServerTimeoutHandler(params map[string]interface{}) error {
	h.logger.Warn("on_rtp_server_timeout hook",
		zap.String("stream_id", getStr(params, "stream_id")),
	)
	return nil
}

func (h *HookHandler) DefaultSendRTPStoppedHandler(params map[string]interface{}) error {
	h.logger.Info("on_send_rtp_stopped hook",
		zap.String("stream_id", getStr(params, "stream_id")),
	)
	return nil
}

func (h *HookHandler) DefaultServerKeepaliveHandler(params map[string]interface{}) error {
	h.logger.Debug("on_server_keepalive hook")
	return nil
}

func (h *HookHandler) DefaultServerStartedHandler(params map[string]interface{}) error {
	h.logger.Info("on_server_started hook — ZLM server started")
	return nil
}

// RegisterDefaultHandlers registers all ZLMediaKit hook handlers.
func (h *HookHandler) RegisterDefaultHandlers() {
	h.Register("on_publish", h.DefaultPublishHookHandler)
	h.Register("on_play", h.DefaultPlayHookHandler)
	h.Register("on_stream_changed", h.DefaultStreamChangedHandler)
	h.Register("on_stream_not_found", h.DefaultStreamNotFoundHandler)
	h.Register("on_stream_none_reader", h.DefaultStreamNoneReaderHandler)
	h.Register("on_rtp_server_timeout", h.DefaultRTPServerTimeoutHandler)
	h.Register("on_send_rtp_stopped", h.DefaultSendRTPStoppedHandler)
	h.Register("on_record_mp4", h.DefaultRecordHookHandler)
	h.Register("on_server_keepalive", h.DefaultServerKeepaliveHandler)
	h.Register("on_server_started", h.DefaultServerStartedHandler)
	// Legacy aliases
	h.Register("on_stream_arrived", h.DefaultStreamChangedHandler)
}

// ── ZLM Config Sync ──

// SyncZLMConfig fetches the current ZLM config, then sets hook URLs and other
// configuration parameters. This matches wvp-GB28181-pro's startup behavior:
//   1. POST /index/api/getServerConfig
//   2. POST /index/api/setServerConfig (with all hook URLs)
func (h *HookHandler) SyncZLMConfig() error {
	hookBaseURL := fmt.Sprintf("http://%s", h.client.GetHookAddr())

	h.logger.Info("Syncing ZLMediaKit configuration",
		zap.String("hook_base_url", hookBaseURL),
	)

	// Step 1: Get current config (to verify connection and see current state)
	currentConfig, err := h.client.GetServerConfig()
	if err != nil {
		h.logger.Warn("Failed to get ZLM server config (will try setServerConfig anyway)",
			zap.Error(err),
		)
	} else {
		h.logger.Info("ZLM current config retrieved",
			zap.Int("config_size", len(currentConfig)),
		)
	}

	// Step 2: Set server config with hook URLs and streaming parameters
	config := map[string]string{
		"hook.enable":                      "1",
		"hook.on_publish":                  hookBaseURL + "/index/hook/on_publish",
		"hook.on_play":                     hookBaseURL + "/index/hook/on_play",
		"hook.on_stream_changed":           hookBaseURL + "/index/hook/on_stream_changed",
		"hook.on_stream_not_found":         hookBaseURL + "/index/hook/on_stream_not_found",
		"hook.on_stream_none_reader":       hookBaseURL + "/index/hook/on_stream_none_reader",
		"hook.on_rtp_server_timeout":       hookBaseURL + "/index/hook/on_rtp_server_timeout",
		"hook.on_send_rtp_stopped":         hookBaseURL + "/index/hook/on_send_rtp_stopped",
		"hook.on_record_mp4":               hookBaseURL + "/index/hook/on_record_mp4",
		"hook.on_server_keepalive":         hookBaseURL + "/index/hook/on_server_keepalive",
		"hook.on_server_started":           hookBaseURL + "/index/hook/on_server_started",
		"hook.alive_interval":              "10.0",
		"hook.timeoutSec":                  "30",
		"rtp_proxy.port_range":             "30000-30500",
		"protocol.continue_push_ms":        "3000",
		"api.secret":                       h.client.config.Secret,
		"ffmpeg.snap":                      "%s -rtsp_transport tcp -i %s -y -f mjpeg -frames:v 1 %s",
	}

	if err := h.client.SetServerConfig(config); err != nil {
		return fmt.Errorf("setServerConfig failed: %w", err)
	}

	h.logger.Info("ZLMediaKit configuration synced successfully",
		zap.String("hook_base_url", hookBaseURL),
	)
	return nil
}

// ── Old SetupHooks (kept for backward compatibility) ──

// SetupHooks configures individual hooks via the old /setHook API.
// Use SyncZLMConfig for the modern approach via /setServerConfig.
func (h *HookHandler) SetupHooks(hookAddr string, hooks []string) error {
	for _, hook := range hooks {
		hookURL := fmt.Sprintf("http://%s/index/hook/%s", hookAddr, hook)
		// Fall back to modern setServerConfig approach
		h.logger.Info("Setting up hook via legacy method",
			zap.String("hook", hook),
			zap.String("url", hookURL),
		)
	}
	// For now, just use SyncZLMConfig
	return h.SyncZLMConfig()
}

// ── Helper functions ──

func getStr(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func getFloat(m map[string]interface{}, key string) float64 {
	if v, ok := m[key].(float64); ok {
		return v
	}
	return 0
}

func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key].(bool); ok {
		return v
	}
	return false
}
