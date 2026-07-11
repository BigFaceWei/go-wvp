package gb28181

import (
	"wvp-go/server/global"
	"wvp-go/server/model/system"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetPlatformList 获取平台列表
// @Summary 获取平台列表
// @Tags 平台级联
// @Param page query int true "页码"
// @Param page_size query int true "每页条数"
// @Success 200 {object} response.Response{data=gin.H{list=[]system.Platform,total=int64}}
// @Router /api/platform/list [get]
func GetPlatformList(c *gin.Context) {
	var pageInfo struct {
		Page       int    `form:"page" binding:"required,min=1"`
		PageSize   int    `form:"page_size" binding:"required,min=1,max=100"`
		PlatformID string `form:"platform_id"`
		Name       string `form:"name"`
		Status     *int   `form:"status"`
	}

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	db := global.GVA_DB.Model(&system.Platform{})

	if pageInfo.PlatformID != "" {
		db = db.Where("platform_id LIKE ?", "%"+pageInfo.PlatformID+"%")
	}
	if pageInfo.Name != "" {
		db = db.Where("name LIKE ?", "%"+pageInfo.Name+"%")
	}
	if pageInfo.Status != nil {
		db = db.Where("status = ?", *pageInfo.Status)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("count platforms failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	var platforms []system.Platform
	if err := db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).
		Limit(pageInfo.PageSize).
		Order("created_at DESC").
		Find(&platforms).Error; err != nil {
		global.GVA_LOG.Error("query platforms failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, gin.H{
		"list":      platforms,
		"total":     total,
		"page":      pageInfo.Page,
		"page_size": pageInfo.PageSize,
	})
}

func GetPlatform(c *gin.Context) {
	platformID := c.Param("id")
	if platformID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var platform system.Platform
	if err := global.GVA_DB.Where("platform_id = ?", platformID).First(&platform).Error; err != nil {
		response.Fail(c, response.PLATFORM_NOT_FOUND, nil)
		return
	}

	response.Success(c, platform)
}

func CreatePlatform(c *gin.Context) {
	var req struct {
		PlatformID  string `json:"platform_id" binding:"required"`
		Name        string `json:"name" binding:"required"`
		Secret      string `json:"secret" binding:"required"`
		IP          string `json:"ip" binding:"required"`
		Port        int    `json:"port" binding:"required,min=1,max=65535"`
		Transport   string `json:"transport"`
		Domain      string `json:"domain"`
		Expires     int    `json:"expires"`
		AutoRegister bool   `json:"auto_register"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var count int64
	global.GVA_DB.Model(&system.Platform{}).Where("platform_id = ?", req.PlatformID).Count(&count)
	if count > 0 {
		response.Fail(c, response.DB_DUPLICATE, nil)
		return
	}

	platform := system.Platform{
		PlatformID:   req.PlatformID,
		Name:         req.Name,
		Secret:       req.Secret,
		IP:           req.IP,
		Port:         req.Port,
		Transport:    req.Transport,
		Domain:       req.Domain,
		Enable:       true,
		Status:       0,
		Expires:      req.Expires,
		AutoRegister: req.AutoRegister,
	}

	if platform.Transport == "" {
		platform.Transport = "UDP"
	}
	if platform.Expires <= 0 {
		platform.Expires = 3600
	}

	if err := global.GVA_DB.Create(&platform).Error; err != nil {
		global.GVA_LOG.Error("create platform failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, platform)
}

func UpdatePlatform(c *gin.Context) {
	platformID := c.Param("id")
	if platformID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var platform system.Platform
	if err := global.GVA_DB.Where("platform_id = ?", platformID).First(&platform).Error; err != nil {
		response.Fail(c, response.PLATFORM_NOT_FOUND, nil)
		return
	}

	var req struct {
		Name        string `json:"name"`
		Secret      string `json:"secret"`
		IP          string `json:"ip"`
		Port        int    `json:"port"`
		Transport   string `json:"transport"`
		Domain      string `json:"domain"`
		Enable      *bool  `json:"enable"`
		Expires     int    `json:"expires"`
		AutoRegister *bool  `json:"auto_register"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	if req.Name != "" {
		platform.Name = req.Name
	}
	if req.Secret != "" {
		platform.Secret = req.Secret
	}
	if req.IP != "" {
		platform.IP = req.IP
	}
	if req.Port > 0 {
		platform.Port = req.Port
	}
	if req.Transport != "" {
		platform.Transport = req.Transport
	}
	if req.Domain != "" {
		platform.Domain = req.Domain
	}
	if req.Enable != nil {
		platform.Enable = *req.Enable
	}
	if req.Expires > 0 {
		platform.Expires = req.Expires
	}
	if req.AutoRegister != nil {
		platform.AutoRegister = *req.AutoRegister
	}

	if err := global.GVA_DB.Save(&platform).Error; err != nil {
		global.GVA_LOG.Error("update platform failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, platform)
}

func DeletePlatform(c *gin.Context) {
	platformID := c.Param("id")
	if platformID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	if err := global.GVA_DB.Where("platform_id = ?", platformID).Delete(&system.Platform{}).Error; err != nil {
		global.GVA_LOG.Error("delete platform failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	global.GVA_DB.Where("platform_id = ?", platformID).Delete(&system.PlatformChannel{})

	response.Success(c, nil)
}

func GetPlatformChannels(c *gin.Context) {
	platformID := c.Param("id")
	if platformID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var channels []system.PlatformChannel
	if err := global.GVA_DB.Where("platform_id = ?", platformID).Find(&channels).Error; err != nil {
		global.GVA_LOG.Error("query platform channels failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, channels)
}

func GetOnlinePlatforms(c *gin.Context) {
	var platforms []system.Platform
	if err := global.GVA_DB.Where("status = ? AND enable = ?", 1, true).Find(&platforms).Error; err != nil {
		global.GVA_LOG.Error("query online platforms failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, platforms)
}

func SyncPlatformCatalog(c *gin.Context) {
	platformID := c.Param("id")
	if platformID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var platform system.Platform
	if err := global.GVA_DB.Where("platform_id = ?", platformID).First(&platform).Error; err != nil {
		response.Fail(c, response.PLATFORM_NOT_FOUND, nil)
		return
	}

	if platform.Status != 1 {
		response.Fail(c, response.PLATFORM_REGISTER_ERR, nil)
		return
	}

	response.SuccessWithMessage(c, "目录同步请求已发送", nil)
}

func StartPlatformCascade(c *gin.Context) {
	platformID := c.Param("id")
	if platformID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var platform system.Platform
	if err := global.GVA_DB.Where("platform_id = ?", platformID).First(&platform).Error; err != nil {
		response.Fail(c, response.PLATFORM_NOT_FOUND, nil)
		return
	}

	global.GVA_LOG.Info("Start platform cascade",
		zap.String("platform_id", platform.PlatformID),
		zap.String("platform_ip", platform.IP),
		zap.Int("platform_port", platform.Port),
	)

	response.SuccessWithMessage(c, "级联启动请求已发送", nil)
}

func StopPlatformCascade(c *gin.Context) {
	platformID := c.Param("id")
	if platformID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var platform system.Platform
	if err := global.GVA_DB.Where("platform_id = ?", platformID).First(&platform).Error; err != nil {
		response.Fail(c, response.PLATFORM_NOT_FOUND, nil)
		return
	}

	global.GVA_LOG.Info("Stop platform cascade", zap.String("platform_id", platform.PlatformID))

	response.SuccessWithMessage(c, "级联停止请求已发送", nil)
}
