package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type PositionRouter struct{}

var positionRouter = new(PositionRouter)

func (p *PositionRouter) InitPositionRouter(Router *gin.RouterGroup) {
	positionRouterGroup := Router.Group("/position")
	{
		positionRouterGroup.GET("/list", gb28181.GetPositionList)
		positionRouterGroup.GET("/:id", gb28181.GetDevicePosition)
		positionRouterGroup.GET("/:id/track", gb28181.GetPositionTrack)
		positionRouterGroup.DELETE("/:id", gb28181.DeletePosition)
	}
}
