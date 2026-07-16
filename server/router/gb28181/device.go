package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type DeviceRouter struct{}

var deviceRouter = new(DeviceRouter)

func (d *DeviceRouter) InitDeviceRouter(Router *gin.RouterGroup) {
	deviceRouterGroup := Router.Group("/device")
	{
		deviceRouterGroup.GET("/list", gb28181.GetDeviceList)
		deviceRouterGroup.GET("/:id", gb28181.GetDevice)
		deviceRouterGroup.POST("", gb28181.CreateDevice)
		deviceRouterGroup.PUT("/:id", gb28181.UpdateDevice)
		deviceRouterGroup.DELETE("/:id", gb28181.DeleteDevice)
		deviceRouterGroup.GET("/:id/channels", gb28181.GetDeviceChannels)
		deviceRouterGroup.POST("/:id/catalog", gb28181.QueryDeviceCatalog)
		deviceRouterGroup.GET("/:id/status", gb28181.GetDeviceStatus)
	}

	channelRouterGroup := Router.Group("/channel")
	{
		channelRouterGroup.GET("/list", gb28181.GetChannelList)
	}
}
