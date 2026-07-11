package system

import (
	"time"

	"wvp-go/server/global"
)

type Platform struct {
	global.GVA_MODEL
	PlatformID    string    `gorm:"type:varchar(64);uniqueIndex;not null" json:"platform_id" form:"platform_id"`
	Name          string    `gorm:"type:varchar(128)" json:"name" form:"name"`
	Secret        string    `gorm:"type:varchar(128)" json:"secret" form:"secret"`
	IP            string    `gorm:"type:varchar(64)" json:"ip" form:"ip"`
	Port          int       `gorm:"type:int;default:5060" json:"port" form:"port"`
	Transport     string    `gorm:"type:varchar(8);default:'UDP'" json:"transport" form:"transport"`
	Domain        string    `gorm:"type:varchar(128)" json:"domain" form:"domain"`
	Enable        bool      `gorm:"type:tinyint(1);default:1" json:"enable" form:"enable"`
	Status        int       `gorm:"type:tinyint;default:0" json:"status" form:"status"` // 0:离线 1:在线
	ChannelCount  int       `gorm:"type:int;default:0" json:"channel_count" form:"channel_count"`
	Expires       int       `gorm:"type:int;default:3600" json:"expires" form:"expires"`
	AutoRegister  bool      `gorm:"type:tinyint(1);default:1" json:"auto_register" form:"auto_register"`
	CreateTime    time.Time `json:"create_time" form:"create_time"`
	UpdateTime    time.Time `json:"update_time" form:"update_time"`
	RegisterTime  time.Time `json:"register_time" form:"register_time"`
	KeepaliveTime time.Time `json:"keepalive_time" form:"keepalive_time"`
}

func (Platform) TableName() string {
	return "wvp_platform"
}

type PlatformChannel struct {
	global.GVA_MODEL
	PlatformID     string `gorm:"type:varchar(64);index;not null" json:"platform_id" form:"platform_id"`
	ChannelID      string `gorm:"type:varchar(64);index;not null" json:"channel_id" form:"channel_id"`
	DeviceID       string `gorm:"type:varchar(64);index" json:"device_id" form:"device_id"`
	Name           string `gorm:"type:varchar(128)" json:"name" form:"name"`
	ChannelType    int    `gorm:"type:tinyint;default:1" json:"channel_type" form:"channel_type"` // 1:设备通道 2:虚拟通道
	Status         string `gorm:"type:varchar(16);default:'OFF'" json:"status" form:"status"`
	Longitude      float64 `gorm:"type:decimal(10,7)" json:"longitude" form:"longitude"`
	Latitude       float64 `gorm:"type:decimal(10,7)" json:"latitude" form:"latitude"`
}

func (PlatformChannel) TableName() string {
	return "wvp_platform_channel"
}
