package gb28181

import (
	"fmt"
	"path/filepath"
	"time"

	"wvp-go/server/global"
	"wvp-go/server/model/system"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetRecordPlanList 获取录像计划列表
// @Summary 获取录像计划列表
// @Tags 录像管理
// @Param page query int true "页码"
// @Param page_size query int true "每页条数"
// @Success 200 {object} response.Response{data=gin.H{list=[]system.RecordPlan,total=int64}}
// @Router /api/record/plan/list [get]
func GetRecordPlanList(c *gin.Context) {
	var pageInfo struct {
		Page     int    `form:"page" binding:"required,min=1"`
		PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
		DeviceID string `form:"device_id"`
		Name     string `form:"name"`
	}

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	db := global.GVA_DB.Model(&system.RecordPlan{})

	if pageInfo.DeviceID != "" {
		db = db.Where("device_id = ?", pageInfo.DeviceID)
	}
	if pageInfo.Name != "" {
		db = db.Where("name LIKE ?", "%"+pageInfo.Name+"%")
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("count record plans failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	var plans []system.RecordPlan
	if err := db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).
		Limit(pageInfo.PageSize).
		Order("created_at DESC").
		Find(&plans).Error; err != nil {
		global.GVA_LOG.Error("query record plans failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, gin.H{
		"list":      plans,
		"total":     total,
		"page":      pageInfo.Page,
		"page_size": pageInfo.PageSize,
	})
}

func GetRecordPlan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var plan system.RecordPlan
	if err := global.GVA_DB.Where("id = ?", id).First(&plan).Error; err != nil {
		response.Fail(c, response.DB_RECORD_NOT_FOUND, nil)
		return
	}

	response.Success(c, plan)
}

func CreateRecordPlan(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		DeviceID    string `json:"device_id" binding:"required"`
		ChannelID   string `json:"channel_id" binding:"required"`
		StreamType  int    `json:"stream_type"`
		RecordType  int    `json:"record_type"`
		StartTime   string `json:"start_time"`
		EndTime     string `json:"end_time"`
		FileDays    int    `json:"file_days"`
		StoragePath string `json:"storage_path"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	plan := system.RecordPlan{
		Name:        req.Name,
		DeviceID:    req.DeviceID,
		ChannelID:   req.ChannelID,
		StreamType:  req.StreamType,
		RecordType:  req.RecordType,
		Enable:      true,
		Status:      0,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		FileDays:    req.FileDays,
		StoragePath: req.StoragePath,
	}

	if plan.StreamType == 0 {
		plan.StreamType = 1
	}
	if plan.RecordType == 0 {
		plan.RecordType = 1
	}
	if plan.FileDays <= 0 {
		plan.FileDays = 30
	}

	if err := global.GVA_DB.Create(&plan).Error; err != nil {
		global.GVA_LOG.Error("create record plan failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, plan)
}

func UpdateRecordPlan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var plan system.RecordPlan
	if err := global.GVA_DB.Where("id = ?", id).First(&plan).Error; err != nil {
		response.Fail(c, response.DB_RECORD_NOT_FOUND, nil)
		return
	}

	var req struct {
		Name        string `json:"name"`
		DeviceID    string `json:"device_id"`
		ChannelID   string `json:"channel_id"`
		StreamType  *int   `json:"stream_type"`
		RecordType  *int   `json:"record_type"`
		Enable      *bool  `json:"enable"`
		StartTime   string `json:"start_time"`
		EndTime     string `json:"end_time"`
		FileDays    *int   `json:"file_days"`
		StoragePath string `json:"storage_path"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	if req.Name != "" {
		plan.Name = req.Name
	}
	if req.DeviceID != "" {
		plan.DeviceID = req.DeviceID
	}
	if req.ChannelID != "" {
		plan.ChannelID = req.ChannelID
	}
	if req.StreamType != nil {
		plan.StreamType = *req.StreamType
	}
	if req.RecordType != nil {
		plan.RecordType = *req.RecordType
	}
	if req.Enable != nil {
		plan.Enable = *req.Enable
	}
	if req.StartTime != "" {
		plan.StartTime = req.StartTime
	}
	if req.EndTime != "" {
		plan.EndTime = req.EndTime
	}
	if req.FileDays != nil {
		plan.FileDays = *req.FileDays
	}
	if req.StoragePath != "" {
		plan.StoragePath = req.StoragePath
	}

	if err := global.GVA_DB.Save(&plan).Error; err != nil {
		global.GVA_LOG.Error("update record plan failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, plan)
}

func DeleteRecordPlan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	if err := global.GVA_DB.Where("id = ?", id).Delete(&system.RecordPlan{}).Error; err != nil {
		global.GVA_LOG.Error("delete record plan failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, nil)
}

func StartRecordPlan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var plan system.RecordPlan
	if err := global.GVA_DB.Where("id = ?", id).First(&plan).Error; err != nil {
		response.Fail(c, response.DB_RECORD_NOT_FOUND, nil)
		return
	}

	plan.Status = 1
	if err := global.GVA_DB.Save(&plan).Error; err != nil {
		global.GVA_LOG.Error("start record plan failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	global.GVA_LOG.Info("Record plan started",
		zap.Uint("plan_id", plan.ID),
		zap.String("device_id", plan.DeviceID),
	)

	response.SuccessWithMessage(c, "录像计划已启动", plan)
}

func StopRecordPlan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var plan system.RecordPlan
	if err := global.GVA_DB.Where("id = ?", id).First(&plan).Error; err != nil {
		response.Fail(c, response.DB_RECORD_NOT_FOUND, nil)
		return
	}

	plan.Status = 0
	if err := global.GVA_DB.Save(&plan).Error; err != nil {
		global.GVA_LOG.Error("stop record plan failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	global.GVA_LOG.Info("Record plan stopped",
		zap.Uint("plan_id", plan.ID),
		zap.String("device_id", plan.DeviceID),
	)

	response.SuccessWithMessage(c, "录像计划已停止", plan)
}

func GetRecordFileList(c *gin.Context) {
	var pageInfo struct {
		Page     int    `form:"page" binding:"required,min=1"`
		PageSize int    `form:"page_size" binding:"required,min=1,max=100"`
		DeviceID string `form:"device_id"`
		ChannelID string `form:"channel_id"`
	}

	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	db := global.GVA_DB.Model(&system.RecordFile{})

	if pageInfo.DeviceID != "" {
		db = db.Where("device_id = ?", pageInfo.DeviceID)
	}
	if pageInfo.ChannelID != "" {
		db = db.Where("channel_id = ?", pageInfo.ChannelID)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("count record files failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	var files []system.RecordFile
	if err := db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).
		Limit(pageInfo.PageSize).
		Order("start_time DESC").
		Find(&files).Error; err != nil {
		global.GVA_LOG.Error("query record files failed", zap.Error(err))
		response.Fail(c, response.DB_ERROR, nil)
		return
	}

	response.Success(c, gin.H{
		"list":      files,
		"total":     total,
		"page":      pageInfo.Page,
		"page_size": pageInfo.PageSize,
	})
}

func DownloadRecordFile(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var file system.RecordFile
	if err := global.GVA_DB.Where("id = ?", id).First(&file).Error; err != nil {
		response.Fail(c, response.RECORD_NOT_FOUND, nil)
		return
	}

	fileName := file.FileName
	if fileName == "" {
		fileName = fmt.Sprintf("%s_%s.ts", file.DeviceID, file.StartTime.Format("20060102150405"))
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	c.Header("Content-Type", "application/octet-stream")
	c.File(filepath.Join(file.FilePath, fileName))
}

func PreviewRecordFile(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	var file system.RecordFile
	if err := global.GVA_DB.Where("id = ?", id).First(&file).Error; err != nil {
		response.Fail(c, response.RECORD_NOT_FOUND, nil)
		return
	}

	previewURL := fmt.Sprintf("/api/record/file/preview/%s/%s/%d", file.DeviceID, file.ChannelID, file.StartTime.Unix())

	response.Success(c, gin.H{
		"file_id":     file.ID,
		"device_id":   file.DeviceID,
		"channel_id":  file.ChannelID,
		"file_name":   file.FileName,
		"file_size":   file.FileSize,
		"start_time":  file.StartTime,
		"end_time":    file.EndTime,
		"duration":    file.Duration,
		"preview_url": previewURL,
		"play_url":    fmt.Sprintf("rtsp://localhost/record/%s/%s", file.DeviceID, file.ChannelID),
	})
}

func PlaybackRecord(c *gin.Context) {
	var req struct {
		DeviceID  string `json:"device_id" binding:"required"`
		ChannelID string `json:"channel_id" binding:"required"`
		StartTime string `json:"start_time" binding:"required"`
		EndTime   string `json:"end_time" binding:"required"`
		Speed     float64 `json:"speed"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	startTime, err := time.Parse("2006-01-02T15:04:05", req.StartTime)
	if err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	endTime, err := time.Parse("2006-01-02T15:04:05", req.EndTime)
	if err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}

	speed := req.Speed
	if speed <= 0 {
		speed = 1.0
	}

	playbackURL := fmt.Sprintf("rtsp://localhost/record/%s/%s?start=%d&end=%d&speed=%.1f",
		req.DeviceID, req.ChannelID, startTime.Unix(), endTime.Unix(), speed)

	global.GVA_LOG.Info("Record playback request",
		zap.String("device_id", req.DeviceID),
		zap.Time("start", startTime),
		zap.Time("end", endTime),
		zap.Float64("speed", speed),
	)

	response.Success(c, gin.H{
		"device_id":    req.DeviceID,
		"channel_id":   req.ChannelID,
		"start_time":   startTime,
		"end_time":     endTime,
		"speed":        speed,
		"playback_url": playbackURL,
	})
}

func StopPlayback(c *gin.Context) {
	deviceID := c.Param("id")
	if deviceID == "" {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	global.GVA_LOG.Info("Stop playback request", zap.String("device_id", deviceID))
	response.Success(c, nil)
}
