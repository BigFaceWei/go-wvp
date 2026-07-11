package gb28181

import (
	"wvp-go/server/global"
	"wvp-go/server/model/system"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetAlarmList 获取报警列表
// @Summary 获取报警列表
// @Tags 报警管理
// @Param page query int true "页码"
// @Param page_size query int true "每页条数"
// @Success 200 {object} response.Response{data=gin.H{list=[]system.Alarm,total=int64}}
// @Router /api/alarm/list [get]
func GetAlarmList(c *gin.Context) {
	var pageInfo struct {
		Page      int    `form:"page" binding:"required,min=1"`
		PageSize  int    `form:"page_size" binding:"required,min=1,max=100"`
		DeviceID  string `form:"device_id"`
		AlarmType string `form:"alarm_type"`
		Status    *int   `form:"handle_status"`
	}

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	db := global.GVA_DB.Model(&system.Alarm{})

	if pageInfo.DeviceID != "" {
		db = db.Where("device_id = ?", pageInfo.DeviceID)
	}
	if pageInfo.AlarmType != "" {
		db = db.Where("alarm_type = ?", pageInfo.AlarmType)
	}
	if pageInfo.Status != nil {
		db = db.Where("handle_status = ?", *pageInfo.Status)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("count alarms failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	var alarms []system.Alarm
	if err := db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).
		Limit(pageInfo.PageSize).
		Order("alarm_time DESC").
		Find(&alarms).Error; err != nil {
		global.GVA_LOG.Error("query alarms failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, gin.H{
		"list":      alarms,
		"total":     total,
		"page":      pageInfo.Page,
		"page_size": pageInfo.PageSize,
	})
}

func GetAlarm(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var alarm system.Alarm
	if err := global.GVA_DB.Where("id = ?", id).First(&alarm).Error; err != nil {
		response.Fail(c, response.DB_RECORD_NOT_FOUND, nil)
		return
	}

	response.Success(c, alarm)
}

func HandleAlarm(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var req struct {
		HandleStatus int    `json:"handle_status" binding:"required"`
		HandleResult string `json:"handle_result"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var alarm system.Alarm
	if err := global.GVA_DB.Where("id = ?", id).First(&alarm).Error; err != nil {
		response.Fail(c, response.DB_RECORD_NOT_FOUND, nil)
		return
	}

	alarm.HandleStatus = req.HandleStatus
	alarm.HandleResult = req.HandleResult
	alarm.HandleTime = global.GVA_DB.NowFunc()

	if err := global.GVA_DB.Save(&alarm).Error; err != nil {
		global.GVA_LOG.Error("update alarm failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, alarm)
}

func DeleteAlarm(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	if err := global.GVA_DB.Where("id = ?", id).Delete(&system.Alarm{}).Error; err != nil {
		global.GVA_LOG.Error("delete alarm failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, nil)
}

func GetAlarmStatistics(c *gin.Context) {
	deviceID := c.Query("device_id")

	db := global.GVA_DB.Model(&system.Alarm{})
	if deviceID != "" {
		db = db.Where("device_id = ?", deviceID)
	}

	var total int64
	db.Count(&total)

	var unhandled int64
	db.Where("handle_status = ?", 0).Count(&unhandled)

	var handled int64
	db.Where("handle_status = ?", 1).Count(&handled)

	var ignored int64
	db.Where("handle_status = ?", 2).Count(&ignored)

	response.Success(c, gin.H{
		"total":     total,
		"unhandled": unhandled,
		"handled":   handled,
		"ignored":   ignored,
	})
}
