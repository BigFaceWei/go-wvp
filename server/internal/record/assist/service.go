package record

import (
	"time"

	"wvp-go/server/global"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
)

type AssistService struct {
	logger *zap.Logger
}

func NewAssistService(logger *zap.Logger) *AssistService {
	return &AssistService{logger: logger}
}

func (s *AssistService) IndexRecordFile(file *system.RecordFile) error {
	global.GVA_LOG.Info("Indexing record file",
		zap.String("file_name", file.FileName),
		zap.String("device_id", file.DeviceID),
	)
	return nil
}

func (s *AssistService) CleanupExpiredFiles(retentionDays int) (int, error) {
	cutoff := time.Now().AddDate(0, 0, -retentionDays)

	var files []system.RecordFile
	if err := global.GVA_DB.Where("end_time < ?", cutoff).Find(&files).Error; err != nil {
		return 0, err
	}

	for _, file := range files {
		global.GVA_LOG.Info("Cleaning up expired record file",
			zap.String("file_name", file.FileName),
			zap.Time("end_time", file.EndTime),
		)
	}

	if err := global.GVA_DB.Where("end_time < ?", cutoff).Delete(&system.RecordFile{}).Error; err != nil {
		return 0, err
	}

	return len(files), nil
}

func (s *AssistService) InitDefaultPlans() error {
	var count int64
	global.GVA_DB.Model(&system.RecordPlan{}).Count(&count)
	if count > 0 {
		return nil
	}

	defaultPlan := system.RecordPlan{
		Name:       "默认录像计划",
		RecordType: 1,
		Enable:     true,
		Status:     0,
		StartTime:  "00:00:00",
		EndTime:    "23:59:59",
		FileDays:   30,
	}

	if err := global.GVA_DB.Create(&defaultPlan).Error; err != nil {
		return err
	}

	global.GVA_LOG.Info("Default record plan created")
	return nil
}
