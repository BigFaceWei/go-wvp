package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type VideoRouter struct{}

var videoRouter = new(VideoRouter)

func (v *VideoRouter) InitVideoRouter(Router *gin.RouterGroup) {
	videoRouterGroup := Router.Group("/video")
	{
		videoRouterGroup.POST("/play", gb28181.PlayVideo)
		videoRouterGroup.POST("/stop/:id", gb28181.StopVideo)
		videoRouterGroup.POST("/ptz/:id", gb28181.PTZControl)
		videoRouterGroup.GET("/status/:id", gb28181.GetVideoStatus)
	}
}
