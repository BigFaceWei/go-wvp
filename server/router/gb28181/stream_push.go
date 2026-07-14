package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type StreamPushRouter struct{}

var streamPushRouter = new(StreamPushRouter)

func (s *StreamPushRouter) InitStreamPushRouter(Router *gin.RouterGroup) {
	r := Router.Group("/stream/push")
	{
		r.GET("/list", gb28181.GetStreamPushList)
		r.GET("/:id", gb28181.GetStreamPush)
		r.POST("", gb28181.CreateStreamPush)
		r.PUT("/:id", gb28181.UpdateStreamPush)
		r.DELETE("/:id", gb28181.DeleteStreamPush)
		r.POST("/:id/start", gb28181.StartStreamPush)
		r.POST("/:id/stop", gb28181.StopStreamPush)
	}
}
