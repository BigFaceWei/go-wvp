package initialize

import (
	"net/http"

	"wvp-go/server/internal/media/zlm"

	"github.com/gin-gonic/gin"
)

// initZLMHookRoutes registers ZLMediaKit webhook endpoints.
// These match wvp-GB28181-pro: /index/hook/on_publish, /index/hook/on_play, etc.
func initZLMHookRoutes(router *gin.Engine) {
	router.POST("/index/hook/on_publish", zlmHookHandler)
	router.POST("/index/hook/on_play", zlmHookHandler)
	router.POST("/index/hook/on_stream_changed", zlmHookHandler)
	router.POST("/index/hook/on_stream_not_found", zlmHookHandler)
	router.POST("/index/hook/on_stream_none_reader", zlmHookHandler)
	router.POST("/index/hook/on_rtp_server_timeout", zlmHookHandler)
	router.POST("/index/hook/on_send_rtp_stopped", zlmHookHandler)
	router.POST("/index/hook/on_record_mp4", zlmHookHandler)
	router.POST("/index/hook/on_server_keepalive", zlmHookHandler)
	router.POST("/index/hook/on_server_started", zlmHookHandler)
}

// zlmHookHandler is the gin-compatible handler that delegates to the ZLM hook handler.
func zlmHookHandler(c *gin.Context) {
	handler := zlm.GetHookHandler()
	if handler == nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "hook handler not initialized"})
		return
	}
	handler.HandleHook(c.Writer, c.Request)
}
