package gb28181

import (
	"fmt"

	"wvp-go/server/global"
	"wvp-go/server/internal/gb28181/service"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var deviceControlService = service.NewDeviceControlService(nil)

// RemoteBoot 远程启动设备
func RemoteBoot(c *gin.Context) {
	deviceID := c.Param("deviceId")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.RemoteBoot(deviceID); err != nil {
		global.GVA_LOG.Error("remote boot failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// RecordControl 远程录像控制
func RecordControl(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		RecordType string `json:"record_type" binding:"required"` // ManualRecord/TimingRecord/ManualStopRecord
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.RecordControl(deviceID, req.RecordType); err != nil {
		global.GVA_LOG.Error("record control failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// GuardControl 布防/撤防控制
func GuardControl(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		Enabled bool `json:"enabled"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.GuardControl(deviceID, req.Enabled); err != nil {
		global.GVA_LOG.Error("guard control failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// HomePosition 归位/复位
func HomePosition(c *gin.Context) {
	deviceID := c.Param("deviceId")
	if err := deviceControlService.HomePosition(deviceID); err != nil {
		global.GVA_LOG.Error("home position failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// IFrame 关键帧请求
func IFrame(c *gin.Context) {
	deviceID := c.Param("deviceId")
	if err := deviceControlService.IFrame(deviceID); err != nil {
		global.GVA_LOG.Error("i frame failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// PresetSet 设置预置位
func PresetSet(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		PresetID string `json:"preset_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.PresetSet(deviceID, req.PresetID); err != nil {
		global.GVA_LOG.Error("preset set failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// PresetGoto 调用预置位
func PresetGoto(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		PresetID string `json:"preset_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.PresetGoto(deviceID, req.PresetID); err != nil {
		global.GVA_LOG.Error("preset goto failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// PresetRemove 删除预置位
func PresetRemove(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		PresetID string `json:"preset_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.PresetRemove(deviceID, req.PresetID); err != nil {
		global.GVA_LOG.Error("preset remove failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// CruiseStart 开始巡航
func CruiseStart(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		CruiseID string `json:"cruise_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.CruiseStart(deviceID, req.CruiseID); err != nil {
		global.GVA_LOG.Error("cruise start failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// CruiseStop 停止巡航
func CruiseStop(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		CruiseID string `json:"cruise_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.CruiseStop(deviceID, req.CruiseID); err != nil {
		global.GVA_LOG.Error("cruise stop failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// ScanStart 开始扫描
func ScanStart(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		ScanID string `json:"scan_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.ScanStart(deviceID, req.ScanID); err != nil {
		global.GVA_LOG.Error("scan start failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// ScanStop 停止扫描
func ScanStop(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		ScanID string `json:"scan_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.ScanStop(deviceID, req.ScanID); err != nil {
		global.GVA_LOG.Error("scan stop failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// Broadcast 语音广播（设备报警触发广播）
func Broadcast(c *gin.Context) {
	deviceID := c.Param("deviceId")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.Broadcast(deviceID); err != nil {
		global.GVA_LOG.Error("broadcast failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// AudioBroadcast 语音对讲/喊话
func AudioBroadcast(c *gin.Context) {
	deviceID := c.Param("deviceId")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.AudioBroadcast(deviceID); err != nil {
		global.GVA_LOG.Error("audio broadcast failed", zap.Error(err))
		response.Fail(c, response.ERROR, fmt.Sprintf("audio broadcast failed: %v", err))
		return
	}
	response.Success(c, nil)
}

// DragZoomIn 拖拽放大
func DragZoomIn(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		Length    string `json:"length" binding:"required"`
		Width     string `json:"width" binding:"required"`
		MidpointX string `json:"midpoint_x" binding:"required"`
		MidpointY string `json:"midpoint_y" binding:"required"`
		LengthX   string `json:"length_x" binding:"required"`
		LengthY   string `json:"length_y" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.DragZoomIn(deviceID, req.Length, req.Width, req.MidpointX, req.MidpointY, req.LengthX, req.LengthY); err != nil {
		global.GVA_LOG.Error("drag zoom in failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

// DragZoomOut 拖拽缩小
func DragZoomOut(c *gin.Context) {
	deviceID := c.Param("deviceId")
	var req struct {
		Length    string `json:"length" binding:"required"`
		Width     string `json:"width" binding:"required"`
		MidpointX string `json:"midpoint_x" binding:"required"`
		MidpointY string `json:"midpoint_y" binding:"required"`
		LengthX   string `json:"length_x" binding:"required"`
		LengthY   string `json:"length_y" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := deviceControlService.DragZoomOut(deviceID, req.Length, req.Width, req.MidpointX, req.MidpointY, req.LengthX, req.LengthY); err != nil {
		global.GVA_LOG.Error("drag zoom out failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}
