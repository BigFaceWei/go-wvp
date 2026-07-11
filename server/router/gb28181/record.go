package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type RecordRouter struct{}

var recordRouter = new(RecordRouter)

func (r *RecordRouter) InitRecordRouter(Router *gin.RouterGroup) {
	recordRouterGroup := Router.Group("/record")
	{
		recordRouterGroup.GET("/plan/list", gb28181.GetRecordPlanList)
		recordRouterGroup.GET("/plan/:id", gb28181.GetRecordPlan)
		recordRouterGroup.POST("/plan", gb28181.CreateRecordPlan)
		recordRouterGroup.PUT("/plan/:id", gb28181.UpdateRecordPlan)
		recordRouterGroup.DELETE("/plan/:id", gb28181.DeleteRecordPlan)
		recordRouterGroup.POST("/plan/:id/start", gb28181.StartRecordPlan)
		recordRouterGroup.POST("/plan/:id/stop", gb28181.StopRecordPlan)
		recordRouterGroup.GET("/file/list", gb28181.GetRecordFileList)
		recordRouterGroup.GET("/file/download/:id", gb28181.DownloadRecordFile)
		recordRouterGroup.GET("/file/preview/:id", gb28181.PreviewRecordFile)
		recordRouterGroup.POST("/playback", gb28181.PlaybackRecord)
		recordRouterGroup.POST("/playback/stop/:id", gb28181.StopPlayback)
	}
}
