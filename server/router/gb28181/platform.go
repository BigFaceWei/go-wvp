package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type PlatformRouter struct{}

var platformRouter = new(PlatformRouter)

func (p *PlatformRouter) InitPlatformRouter(Router *gin.RouterGroup) {
	platformRouterGroup := Router.Group("/platform")
	{
		platformRouterGroup.GET("/list", gb28181.GetPlatformList)
		platformRouterGroup.GET("/online", gb28181.GetOnlinePlatforms)
		platformRouterGroup.GET("/:id", gb28181.GetPlatform)
		platformRouterGroup.POST("", gb28181.CreatePlatform)
		platformRouterGroup.PUT("/:id", gb28181.UpdatePlatform)
		platformRouterGroup.DELETE("/:id", gb28181.DeletePlatform)
		platformRouterGroup.GET("/:id/channels", gb28181.GetPlatformChannels)
		platformRouterGroup.POST("/:id/sync", gb28181.SyncPlatformCatalog)
		platformRouterGroup.POST("/:id/start", gb28181.StartPlatformCascade)
		platformRouterGroup.POST("/:id/stop", gb28181.StopPlatformCascade)
	}
}
