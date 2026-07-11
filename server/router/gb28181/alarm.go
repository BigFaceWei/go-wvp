package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type AlarmRouter struct{}

var alarmRouter = new(AlarmRouter)

func (a *AlarmRouter) InitAlarmRouter(Router *gin.RouterGroup) {
	alarmRouterGroup := Router.Group("/alarm")
	{
		alarmRouterGroup.GET("/list", gb28181.GetAlarmList)
		alarmRouterGroup.GET("/statistics", gb28181.GetAlarmStatistics)
		alarmRouterGroup.GET("/:id", gb28181.GetAlarm)
		alarmRouterGroup.POST("/:id/handle", gb28181.HandleAlarm)
		alarmRouterGroup.DELETE("/:id", gb28181.DeleteAlarm)
	}
}
