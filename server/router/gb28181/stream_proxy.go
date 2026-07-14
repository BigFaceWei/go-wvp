package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type StreamProxyRouter struct{}

var streamProxyRouter = new(StreamProxyRouter)

func (s *StreamProxyRouter) InitStreamProxyRouter(Router *gin.RouterGroup) {
	r := Router.Group("/stream/proxy")
	{
		r.GET("/list", gb28181.GetStreamProxyList)
		r.GET("/:id", gb28181.GetStreamProxy)
		r.POST("", gb28181.CreateStreamProxy)
		r.PUT("/:id", gb28181.UpdateStreamProxy)
		r.DELETE("/:id", gb28181.DeleteStreamProxy)
		r.POST("/:id/start", gb28181.StartStreamProxy)
		r.POST("/:id/stop", gb28181.StopStreamProxy)
	}
}
