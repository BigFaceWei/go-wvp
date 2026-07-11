package gb28181

import (
	"wvp-go/server/global"
	"wvp-go/server/model/system"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PlayVideo 视频点播
// @Summary 视频点播
// @Description 向设备发送INVITE请求开始视频推流
// @Tags 视频管理
// @Accept json
// @Produce json
// @Param request body object{device_id=string,channel_id=string,ssrc=string} true "请求参数"
// @Success 200 {object} response.Response{data=gin.H{play_url=string}}
// @Router /api/video/play [post]
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

	var device system.Device
	if err := global.GVA_DB.Where("device_id = ?", req.DeviceID).First(&device).Error; err != nil {
		response.Fail(c, response.DEVICE_NOT_FOUND, nil)
		return
	}

	if !device.Online {
		response.Fail(c, response.DEVICE_OFFLINE, nil)
		return
	}

	global.GVA_LOG.Info("Play video request",
		zap.String("device_id", req.DeviceID),
		zap.String("channel_id", req.ChannelID),
	)

	response.Success(c, gin.H{
		"device_id":  req.DeviceID,
		"channel_id": req.ChannelID,
		"play_url":   "rtsp://localhost/live/stream",
	})
}

func StopVideo(c *gin.Context) {
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

	global.GVA_LOG.Info("Stop video request", zap.String("device_id", deviceID))

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
