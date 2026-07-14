package system

import (
	"wvp-go/server/global"
)

// StreamPush 推流管理
type StreamPush struct {
	global.GVA_MODEL
	Name          string `gorm:"type:varchar(128);not null" json:"name" form:"name"`
	App           string `gorm:"type:varchar(64);default:'push'" json:"app" form:"app"`
	Stream        string `gorm:"type:varchar(128);uniqueIndex;not null" json:"stream" form:"stream"`
	MediaServerID string `gorm:"type:varchar(64)" json:"media_server_id" form:"media_server_id"`
	PushType      string `gorm:"type:varchar(32);default:'rtmp'" json:"push_type" form:"push_type"` // rtmp/gb28181
	PushURL       string `gorm:"type:varchar(512)" json:"push_url" form:"push_url"`
	Status        bool   `gorm:"type:tinyint(1);default:0" json:"status" form:"status"`
	Enable        bool   `gorm:"type:tinyint(1);default:1" json:"enable" form:"enable"`
	Online        bool   `gorm:"type:tinyint(1);default:0" json:"online" form:"online"`
	StreamKey     string `gorm:"type:varchar(128)" json:"stream_key" form:"stream_key"`
	DeviceID      string `gorm:"type:varchar(64)" json:"device_id" form:"device_id"`
	ChannelID     string `gorm:"type:varchar(64)" json:"channel_id" form:"channel_id"`
}

func (StreamPush) TableName() string {
	return "wvp_stream_push"
}
