package system

import (
	"time"

	"wvp-go/server/global"
)

type RecordPlan struct {
	global.GVA_MODEL
	Name          string `gorm:"type:varchar(128);not null" json:"name" form:"name"`
	DeviceID      string `gorm:"type:varchar(64);index;not null" json:"device_id" form:"device_id"`
	ChannelID     string `gorm:"type:varchar(64);index;not null" json:"channel_id" form:"channel_id"`
	StreamType    int    `gorm:"type:tinyint;default:1" json:"stream_type" form:"stream_type"` // 1:主码流 2:子码流
	RecordType    int    `gorm:"type:tinyint;default:1" json:"record_type" form:"record_type"` // 1:连续录像 2:定时录像 3:告警录像
	Enable        bool   `gorm:"type:tinyint(1);default:1" json:"enable" form:"enable"`
	Status        int    `gorm:"type:tinyint;default:0" json:"status" form:"status"` // 0:停止 1:录像中
	StartTime     string `gorm:"type:varchar(16)" json:"start_time" form:"start_time"`
	EndTime       string `gorm:"type(varchar(16))" json:"end_time" form:"end_time"`
	FileDays      int    `gorm:"type:int;default:30" json:"file_days" form:"file_days"`
	StoragePath   string `gorm:"type:varchar(256)" json:"storage_path" form:"storage_path"`
	CreateTime    time.Time `json:"create_time" form:"create_time"`
	UpdateTime    time.Time `json:"update_time" form:"update_time"`
}

func (RecordPlan) TableName() string {
	return "wvp_record_plan"
}

type RecordFile struct {
	global.GVA_MODEL
	DeviceID     string    `gorm:"type:varchar(64);index;not null" json:"device_id" form:"device_id"`
	ChannelID    string    `gorm:"type:varchar(64);index;not null" json:"channel_id" form:"channel_id"`
	PlanID       uint      `gorm:"index" json:"plan_id" form:"plan_id"`
	FileName     string    `gorm:"type:varchar(256);not null" json:"file_name" form:"file_name"`
	FilePath     string    `gorm:"type:varchar(512);not null" json:"file_path" form:"file_path"`
	FileSize     int64     `gorm:"type:bigint;default:0" json:"file_size" form:"file_size"`
	StartTime    time.Time `gorm:"index" json:"start_time" form:"start_time"`
	EndTime      time.Time `json:"end_time" form:"end_time"`
	Duration     int       `gorm:"type:int;default:0" json:"duration" form:"duration"`
	StorageType  int       `gorm:"type:tinyint;default:1" json:"storage_type" form:"storage_type"` // 1:本地 2:MinIO/S3
	Indexes      string    `gorm:"type:text" json:"indexes" form:"indexes"`
}

func (RecordFile) TableName() string {
	return "wvp_record_file"
}
