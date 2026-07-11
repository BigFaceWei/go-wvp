package gb28181

import (
	"github.com/gin-gonic/gin"
)

type Router struct{}

var router = new(Router)

func (r *Router) InitGB28181Router(Router *gin.RouterGroup) {
	deviceRouter.InitDeviceRouter(Router)
	videoRouter.InitVideoRouter(Router)
	platformRouter.InitPlatformRouter(Router)
	alarmRouter.InitAlarmRouter(Router)
	positionRouter.InitPositionRouter(Router)
	recordRouter.InitRecordRouter(Router)
}