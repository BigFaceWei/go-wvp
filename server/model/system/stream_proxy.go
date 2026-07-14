package system

import (
	"wvp-go/server/global"
)

// StreamProxy 拉流代理配置
type StreamProxy struct {
	global.GVA_MODEL
	Name        string `gorm:"type:varchar(128);not null" json:"name" form:"name"`
	StreamType  string `gorm:"type:varchar(32);default:'rtsp'" json:"stream_type" form:"stream_type"` // rtsp/rtmp/http-flv
	URL         string `gorm:"type:varchar(512);not null" json:"url" form:"url"`
	SrcURL      string `gorm:"type:varchar(512)" json:"src_url" form:"src_url"`
	App         string `gorm:"type:varchar(64);default:'proxy'" json:"app" form:"app"`
	Stream      string `gorm:"type:varchar(128);uniqueIndex;not null" json:"stream" form:"stream"`
	MediaServerID string `gorm:"type:varchar(64)" json:"media_server_id" form:"media_server_id"`
	Status      bool   `gorm:"type:tinyint(1);default:0" json:"status" form:"status"`
	Enable      bool   `gorm:"type:tinyint(1);default:1" json:"enable" form:"enable"`
	EnableRTSP  bool   `gorm:"type:tinyint(1);default:1" json:"enable_rtsp" form:"enable_rtsp"`
	EnableRTMP  bool   `gorm:"type:tinyint(1);default:1" json:"enable_rtmp" form:"enable_rtmp"`
	EnableHLS   bool   `gorm:"type:tinyint(1);default:1" json:"enable_hls" form:"enable_hls"`
	EnableFLV   bool   `gorm:"type:tinyint(1);default:1" json:"enable_flv" form:"enable_flv"`
	EnableAudio bool   `gorm:"type:tinyint(1);default:1" json:"enable_audio" form:"enable_audio"`
	FFmppegCmd  string `gorm:"type:varchar(512)" json:"ffmpeg_cmd" form:"ffmpeg_cmd"`
	TimeoutSec  int    `gorm:"type:int;default:30" json:"timeout_sec" form:"timeout_sec"`
	StatusCode  string `gorm:"type:varchar(32)" json:"status_code" form:"status_code"`
	StreamKey   string `gorm:"type:varchar(128)" json:"stream_key" form:"stream_key"`
	Online      bool   `gorm:"type:tinyint(1);default:0" json:"online" form:"online"`
	DeviceID    string `gorm:"type:varchar(64)" json:"device_id" form:"device_id"`
	ChannelID   string `gorm:"type:varchar(64)" json:"channel_id" form:"channel_id"`
}

func (StreamProxy) TableName() string {
	return "wvp_stream_proxy"
}
