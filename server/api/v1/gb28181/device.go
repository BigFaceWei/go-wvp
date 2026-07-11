package gb28181

import (
	"wvp-go/server/global"
	"wvp-go/server/model/system"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetDeviceList 获取设备列表
// @Summary 获取设备列表
// @Description 分页查询设备列表，支持按设备ID、名称、在线状态筛选
// @Tags 设备管理
// @Param page query int true "页码"
// @Param page_size query int true "每页条数"
// @Param device_id query string false "设备ID"
// @Param name query string false "设备名称"
// @Param online query bool false "在线状态"
// @Success 200 {object} response.Response{data=gin.H{list=[]system.Device,total=int64}}
// @Router /api/device/list [get]
func GetDeviceList(c *gin.Context) {
	var pageInfo struct {
		Page     int    `form:"page" binding:"required,min=1"`
		PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
		DeviceID string `form:"device_id"`
		Name     string `form:"name"`
		Online   *bool  `form:"online"`
	}

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	db := global.GVA_DB.Model(&system.Device{})

	if pageInfo.DeviceID != "" {
		db = db.Where("device_id LIKE ?", "%"+pageInfo.DeviceID+"%")
	}
	if pageInfo.Name != "" {
		db = db.Where("name LIKE ?", "%"+pageInfo.Name+"%")
	}
	if pageInfo.Online != nil {
		db = db.Where("online = ?", *pageInfo.Online)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("count devices failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	var devices []system.Device
	if err := db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).
		Limit(pageInfo.PageSize).
		Order("created_at DESC").
		Find(&devices).Error; err != nil {
		global.GVA_LOG.Error("query devices failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, gin.H{
		"list":      devices,
		"total":     total,
		"page":      pageInfo.Page,
		"page_size": pageInfo.PageSize,
	})
}

// GetDevice 获取设备详情
// @Summary 获取设备详情
// @Description 根据设备ID获取设备详细信息
// @Tags 设备管理
// @Param id path string true "设备ID"
// @Success 200 {object} response.Response{data=system.Device}
// @Router /api/device/{id} [get]
func GetDevice(c *gin.Context) {
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

	response.Success(c, device)
}

func CreateDevice(c *gin.Context) {
	var req struct {
		DeviceID  string `json:"device_id" binding:"required"`
		Name      string `json:"name"`
		Password  string `json:"password"`
		Transport string `json:"transport"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var count int64
	global.GVA_DB.Model(&system.Device{}).Where("device_id = ?", req.DeviceID).Count(&count)
	if count > 0 {
		response.Fail(c, response.DEVICE_ALREADY_EXISTS, nil)
		return
	}

	device := system.Device{
		DeviceID:  req.DeviceID,
		Name:      req.Name,
		Password:  req.Password,
		Transport: req.Transport,
		Online:    false,
	}

	if err := global.GVA_DB.Create(&device).Error; err != nil {
		global.GVA_LOG.Error("create device failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, device)
}

func UpdateDevice(c *gin.Context) {
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

	var req struct {
		Name      string `json:"name"`
		Password  string `json:"password"`
		Transport string `json:"transport"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	if req.Name != "" {
		device.Name = req.Name
	}
	if req.Password != "" {
		device.Password = req.Password
	}
	if req.Transport != "" {
		device.Transport = req.Transport
	}

	if err := global.GVA_DB.Save(&device).Error; err != nil {
		global.GVA_LOG.Error("update device failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, device)
}

func DeleteDevice(c *gin.Context) {
	deviceID := c.Param("id")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	if err := global.GVA_DB.Where("device_id = ?", deviceID).Delete(&system.Device{}).Error; err != nil {
		global.GVA_LOG.Error("delete device failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	global.GVA_DB.Where("device_id = ?", deviceID).Delete(&system.DeviceChannel{})

	response.Success(c, nil)
}

func GetDeviceChannels(c *gin.Context) {
	deviceID := c.Param("id")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var channels []system.DeviceChannel
	if err := global.GVA_DB.Where("device_id = ?", deviceID).Find(&channels).Error; err != nil {
		global.GVA_LOG.Error("query channels failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, channels)
}

func QueryDeviceCatalog(c *gin.Context) {
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

	if !device.Online {
		response.Fail(c, response.DEVICE_OFFLINE, nil)
		return
	}

	response.SuccessWithMessage(c, "目录查询请求已发送", nil)
}

func GetDeviceStatus(c *gin.Context) {
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
		"device_id":      device.DeviceID,
		"online":         device.Online,
		"ip":             device.IP,
		"port":           device.Port,
		"register_time":  device.RegisterTime,
		"keepalive_time": device.KeepaliveTime,
	})
}