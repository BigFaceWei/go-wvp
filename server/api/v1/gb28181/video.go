package gb28181

import (
	"fmt"
	"net/http"

	"wvp-go/server/global"
	"wvp-go/server/internal/gb28181/service"
	"wvp-go/server/model/system"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PlayVideo(c *gin.Context) {
	var req struct {
		DeviceID  string `json:"device_id" binding:"required"`
		ChannelID string `json:"channel_id" binding:"required"`
		SSRC      string `json:"ssrc"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	result, err := service.PlayVideo(req.DeviceID, req.ChannelID, req.SSRC)
	if err != nil {
		response.Fail(c, response.DEVICE_NOT_FOUND, fmt.Sprintf("play failed: %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}

func StopVideo(c *gin.Context) {
	deviceID := c.Param("id")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	if err := service.StopVideo(deviceID); err != nil {
		response.Fail(c, response.DEVICE_NOT_FOUND, fmt.Sprintf("stop failed: %v", err))
		return
	}

	response.Success(c, nil)
}

func PTZControl(c *gin.Context) {
	deviceID := c.Param("id")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var req struct {
		Direction string `json:"direction" binding:"required"`
		Speed     int    `json:"speed"`
		Zoom      int    `json:"zoom"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var device system.Device
	if err := global.GVA_DB.Where("device_id = ?", deviceID).First(&device).Error; err != nil {
		response.Fail(c, response.DEVICE_NOT_FOUND, nil)
		return
	}

	if !device.Online {
		response.Fail(c, response.DEVICE_OFFLINE, nil)
		return
	}

	global.GVA_LOG.Info("PTZ control",
		zap.String("device_id", deviceID),
		zap.String("direction", req.Direction),
		zap.Int("speed", req.Speed),
		zap.Int("zoom", req.Zoom),
	)

	response.Success(c, nil)
}

func GetVideoStatus(c *gin.Context) {
	deviceID := c.Param("id")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var device system.Device
	if err := global.GVA_DB.Where("device_id = ?", deviceID).First(&device).Error; err != nil {
		response.Fail(c, response.DEVICE_NOT_FOUND, nil)
		return
	}

	response.Success(c, gin.H{
		"device_id": device.DeviceID,
		"online":    device.Online,
		"streaming": false,
	})
}
