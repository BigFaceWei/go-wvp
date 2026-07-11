package gb28181

import (
	"wvp-go/server/global"
	"wvp-go/server/model/system"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetPositionList 获取位置列表
// @Summary 获取位置列表
// @Tags 移动位置
// @Param page query int true "页码"
// @Param page_size query int true "每页条数"
// @Success 200 {object} response.Response{data=gin.H{list=[]system.DevicePosition,total=int64}}
// @Router /api/position/list [get]
func GetPositionList(c *gin.Context) {
	var pageInfo struct {
		Page     int    `form:"page" binding:"required,min=1"`
		PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
		DeviceID string `form:"device_id"`
	}

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	db := global.GVA_DB.Model(&system.DevicePosition{})

	if pageInfo.DeviceID != "" {
		db = db.Where("device_id = ?", pageInfo.DeviceID)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("count positions failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	var positions []system.DevicePosition
	if err := db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).
		Limit(pageInfo.PageSize).
		Order("report_time DESC").
		Find(&positions).Error; err != nil {
		global.GVA_LOG.Error("query positions failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, gin.H{
		"list":      positions,
		"total":     total,
		"page":      pageInfo.Page,
		"page_size": pageInfo.PageSize,
	})
}

func GetDevicePosition(c *gin.Context) {
	deviceID := c.Param("id")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var position system.DevicePosition
	if err := global.GVA_DB.Where("device_id = ?", deviceID).Order("report_time DESC").First(&position).Error; err != nil {
		response.Fail(c, response.DB_RECORD_NOT_FOUND, nil)
		return
	}

	response.Success(c, position)
}

func GetPositionTrack(c *gin.Context) {
	deviceID := c.Param("id")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var positions []system.DevicePosition
	if err := global.GVA_DB.Where("device_id = ?", deviceID).
		Order("report_time ASC").
		Find(&positions).Error; err != nil {
		global.GVA_LOG.Error("query position track failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, positions)
}

func DeletePosition(c *gin.Context) {
	deviceID := c.Param("id")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	if err := global.GVA_DB.Where("device_id = ?", deviceID).Delete(&system.DevicePosition{}).Error; err != nil {
		global.GVA_LOG.Error("delete position failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, nil)
}
