package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type DeviceControlRouter struct{}

var deviceControlRouter = new(DeviceControlRouter)

func (d *DeviceControlRouter) InitDeviceControlRouter(Router *gin.RouterGroup) {
	// 设备远程控制
	r := Router.Group("/device/control")
	{
		r.POST("/:deviceId/teleboot", gb28181.RemoteBoot)
		r.POST("/:deviceId/record", gb28181.RecordControl)
		r.POST("/:deviceId/guard", gb28181.GuardControl)
		r.POST("/:deviceId/home", gb28181.HomePosition)
		r.POST("/:deviceId/iframe", gb28181.IFrame)
		r.POST("/:deviceId/dragzoom/in", gb28181.DragZoomIn)
		r.POST("/:deviceId/dragzoom/out", gb28181.DragZoomOut)
	}

	// 预置位
	preset := Router.Group("/device/preset")
	{
		preset.POST("/:deviceId/set", gb28181.PresetSet)
		preset.POST("/:deviceId/goto", gb28181.PresetGoto)
		preset.POST("/:deviceId/remove", gb28181.PresetRemove)
	}

	// 巡航
	cruise := Router.Group("/device/cruise")
	{
		cruise.POST("/:deviceId/start", gb28181.CruiseStart)
		cruise.POST("/:deviceId/stop", gb28181.CruiseStop)
	}

	// 扫描
	scan := Router.Group("/device/scan")
	{
		scan.POST("/:deviceId/start", gb28181.ScanStart)
		scan.POST("/:deviceId/stop", gb28181.ScanStop)
	}

	// 语音广播与对讲
	broadcast := Router.Group("/device/broadcast")
	{
		broadcast.POST("/:deviceId", gb28181.Broadcast)
		broadcast.POST("/:deviceId/audio", gb28181.AudioBroadcast)
	}
}
