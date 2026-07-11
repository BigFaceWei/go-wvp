package zlm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

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
	hookName := r.URL.Query().Get("hook")
	if hookName == "" {
		http.Error(w, "missing hook parameter", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "read body failed", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var params map[string]interface{}
	if err := json.Unmarshal(body, &params); err != nil {
		http.Error(w, "parse body failed", http.StatusBadRequest)
		return
	}

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

func (h *HookHandler) DefaultStreamChangedHandler(params map[string]interface{}) error {
	h.logger.Info("Stream changed",
		zap.Any("params", params),
	)
	return nil
}

func (h *HookHandler) DefaultStreamNotFoundHandler(params map[string]interface{}) error {
	h.logger.Warn("Stream not found",
		zap.Any("params", params),
	)
	return nil
}

func (h *HookHandler) DefaultPlayHookHandler(params map[string]interface{}) error {
	h.logger.Info("Play hook",
		zap.Any("params", params),
	)
	return nil
}

func (h *HookHandler) DefaultRecordHookHandler(params map[string]interface{}) error {
	h.logger.Info("Record hook",
		zap.Any("params", params),
	)
	return nil
}

func (h *HookHandler) RegisterDefaultHandlers() {
	h.Register("on_stream_changed", h.DefaultStreamChangedHandler)
	h.Register("on_stream_not_found", h.DefaultStreamNotFoundHandler)
	h.Register("on_play", h.DefaultPlayHookHandler)
	h.Register("on_record", h.DefaultRecordHookHandler)
	h.Register("on_stream_arrived", h.DefaultStreamChangedHandler)
}

type HookServer struct {
	handler *HookHandler
	addr    string
	server  *http.Server
	logger  *zap.Logger
}

func NewHookServer(handler *HookHandler, addr string, logger *zap.Logger) *HookServer {
	return &HookServer{
		handler: handler,
		addr:    addr,
		logger:  logger,
	}
}

func (s *HookServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hook", s.handler.HandleHook)

	s.server = &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	s.logger.Info("Hook server starting", zap.String("addr", s.addr))

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("Hook server error", zap.Error(err))
		}
	}()

	return nil
}

func (s *HookServer) Stop() error {
	if s.server != nil {
		return s.server.Close()
	}
	return nil
}

func (h *HookHandler) SetupHooks(hookAddr string, hooks []string) error {
	for _, hook := range hooks {
		params := map[string]string{
			"hook": fmt.Sprintf("http://%s/hook?hook=%s", hookAddr, hook),
		}
		_, err := h.client.callAPI("/index/api/setHook", params)
		if err != nil {
			h.logger.Error("Setup hook failed",
				zap.String("hook", hook),
				zap.Error(err),
			)
			return err
		}
		h.logger.Info("Hook registered", zap.String("hook", hook))
	}
	return nil
}
