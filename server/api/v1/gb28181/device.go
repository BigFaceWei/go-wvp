package gb28181

import (
	"fmt"
	"strings"
	"time"

	"wvp-go/server/global"
	"wvp-go/server/internal/gb28181/message"
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
		Online   string `form:"online"` // 使用 string 类型，空字符串表示未传值
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
	// 只有明确传了 online 参数才筛选
	if pageInfo.Online != "" {
		online := pageInfo.Online == "true" || pageInfo.Online == "1"
		db = db.Where("online = ?", online)
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

// GetChannelList 获取通道列表（分页+筛选）
func GetChannelList(c *gin.Context) {
	var pageInfo struct {
		Page      int    `form:"page" binding:"required,min=1"`
		PageSize  int    `form:"page_size" binding:"required,min=1,max=100"`
		DeviceID  string `form:"device_id"`
		ChannelID string `form:"channel_id"`
		Name      string `form:"name"`
		OnLine    string `form:"on_line"`
	}

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	db := global.GVA_DB.Model(&system.DeviceChannel{})

	if pageInfo.DeviceID != "" {
		db = db.Where("device_id = ?", pageInfo.DeviceID)
	}
	if pageInfo.ChannelID != "" {
		db = db.Where("channel_id LIKE ?", "%"+pageInfo.ChannelID+"%")
	}
	if pageInfo.Name != "" {
		db = db.Where("name LIKE ?", "%"+pageInfo.Name+"%")
	}
	if pageInfo.OnLine != "" {
		online := pageInfo.OnLine == "true" || pageInfo.OnLine == "1"
		status := "OFF"
		if online {
			status = "ON"
		}
		db = db.Where("status = ?", status)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("count channels failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	var channels []system.DeviceChannel
	if err := db.Offset((pageInfo.Page-1)*pageInfo.PageSize).
		Limit(pageInfo.PageSize).
		Order("created_at DESC").
		Find(&channels).Error; err != nil {
		global.GVA_LOG.Error("query channels failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, gin.H{
		"list":      channels,
		"total":     total,
		"page":      pageInfo.Page,
		"page_size": pageInfo.PageSize,
	})
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

	if global.GVA_SIP_SERVER == nil {
		global.GVA_LOG.Error("SIP server not available")
		response.Fail(c, response.SIP_INIT_FAILED, nil)
		return
	}

	xmlBody := fmt.Sprintf(`<?xml version="1.0" encoding="GB2312"?>
<Query>
  <CmdType>Catalog</CmdType>
  <SN>1</SN>
  <DeviceID>%s</DeviceID>
</Query>`, deviceID)

	requestURI := fmt.Sprintf("sip:%s@%s", deviceID, global.GVA_CONFIG.WVP.SIP.Domain)
	headers := map[string]string{
		"To":           fmt.Sprintf("<sip:%s@%s>", deviceID, global.GVA_CONFIG.WVP.SIP.Domain),
		"Content-Type": "Application/MANSCDP+xml",
	}

	var targetAddr string
	if device.IP == "" {
		global.GVA_LOG.Error("Device IP is empty, cannot send catalog query",
			zap.String("device_id", deviceID),
		)
		response.Fail(c, response.INVALID_PARAMS, fmt.Sprintf("device IP is empty: %s", deviceID))
		return
	}
	if strings.Contains(device.IP, ":") {
		targetAddr = device.IP
	} else {
		targetAddr = fmt.Sprintf("%s:%d", device.IP, device.Port)
	}

	// Register pending query for synchronous wait
	ch := message.GlobalCatalogHandler.RegisterPendingQuery(deviceID, "1")
	defer message.GlobalCatalogHandler.RemovePendingQuery(deviceID)

	_, err := global.GVA_SIP_SERVER.SendRequestTo("MESSAGE", requestURI, targetAddr, headers, []byte(xmlBody))
	if err != nil {
		global.GVA_LOG.Error("Send catalog query failed", zap.Error(err))
		response.Fail(c, response.SIP_SEND_FAILED, nil)
		return
	}

	global.GVA_LOG.Info("Catalog query sent, waiting for response",
		zap.String("device_id", deviceID),
		zap.String("addr", targetAddr),
	)

	// Wait for all catalog responses with 10s timeout
	select {
	case result := <-ch.ResultCh:
		if result.Success && len(result.Items) >= result.SumNum {
			// All items received: delete old channels and batch insert fresh data
			if err := global.GVA_DB.Unscoped().Where("device_id = ?", deviceID).Delete(&system.DeviceChannel{}).Error; err != nil {
				global.GVA_LOG.Error("delete existing channels failed",
					zap.String("device_id", deviceID),
					zap.Error(err),
				)
				response.Fail(c, response.DB_ERROR, nil)
				return
			}

			for _, item := range result.Items {
				channel := system.DeviceChannel{
					DeviceID:     deviceID,
					ChannelID:    item.DeviceID,
					Name:         item.Name,
					Manufacturer: item.Manufacturer,
					Model:        item.Model,
					Owner:        item.Owner,
					CivilCode:    item.CivilCode,
					Address:      item.Address,
					Parental:     item.Parental,
					ParentID:     item.ParentID,
					SafetyWay:    item.SafetyWay,
					RegisterWay:  item.RegisterWay,
					Secrecy:      item.Secrecy,
					Status:       item.Status,
				}
				if err := global.GVA_DB.Create(&channel).Error; err != nil {
					global.GVA_LOG.Error("create channel failed",
						zap.String("device_id", deviceID),
						zap.String("channel_id", item.DeviceID),
						zap.Error(err),
					)
				}
			}

			global.GVA_LOG.Info("Catalog query completed successfully",
				zap.String("device_id", deviceID),
				zap.Int("channels", len(result.Items)),
			)
			// Update channel count on device record
			global.GVA_DB.Table("wvp_device").Where("device_id = ?", deviceID).Update("channel_count", len(result.Items))
			response.Success(c, gin.H{
				"success":  true,
				"expected": result.SumNum,
				"actual":   result.Received,
				"message":  fmt.Sprintf("查询成功，应查询 %d 个，实际查询 %d 个", result.SumNum, result.Received),
			})
		} else {
			// Not all items received — do NOT touch DB, return failure info
			global.GVA_LOG.Warn("Catalog query incomplete",
				zap.String("device_id", deviceID),
				zap.Int("received", result.Received),
				zap.Int("sum_num", result.SumNum),
			)
			response.FailWithDetail(c, response.SIP_TIMEOUT,
				fmt.Sprintf("查询失败，应查询 %d 个，实际查询 %d 个", result.SumNum, result.Received),
				gin.H{
					"success":  false,
					"expected": result.SumNum,
					"actual":   result.Received,
				})
		}

	case <-time.After(10 * time.Second):
		// Timeout — check what we have so far
		if pq, ok := message.GlobalCatalogHandler.GetPendingQuery(deviceID); ok {
			received, sumNum := pq.Progress()

			global.GVA_LOG.Warn("Catalog query timeout",
				zap.String("device_id", deviceID),
				zap.Int("received", received),
				zap.Int("sum_num", sumNum),
			)
			if sumNum > 0 {
				response.FailWithDetail(c, response.SIP_TIMEOUT,
					fmt.Sprintf("查询超时，应查询 %d 个，实际查询 %d 个", sumNum, received),
					gin.H{
						"success":  false,
						"expected": sumNum,
						"actual":   received,
					})
			} else {
				response.FailWithDetail(c, response.SIP_TIMEOUT,
					"查询超时，未收到任何响应",
					gin.H{
						"success":  false,
						"expected": 0,
						"actual":   0,
					})
			}
		} else {
			response.FailWithDetail(c, response.SIP_TIMEOUT,
				"查询超时",
				gin.H{
					"success":  false,
					"expected": 0,
					"actual":   0,
				})
		}
	}
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