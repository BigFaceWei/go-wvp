package system

import (
	"time"

	"wvp-go/server/global"
)

type DevicePosition struct {
	global.GVA_MODEL
	DeviceID    string    `gorm:"type:varchar(64);index;not null" json:"device_id" form:"device_id"`
	ChannelID   string    `gorm:"type:varchar(64);index;not null" json:"channel_id" form:"channel_id"`
	Longitude   float64   `gorm:"type:decimal(10,7);not null" json:"longitude" form:"longitude"`
	Latitude    float64   `gorm:"type:decimal(10,7);not null" json:"latitude" form:"latitude"`
	Altitude    float64   `gorm:"type:decimal(10,2);default:0" json:"altitude" form:"altitude"`
	Speed       float64   `gorm:"type:decimal(10,2);default:0" json:"speed" form:"speed"`
	Direction   float64   `gorm:"type:decimal(10,2);default:0" json:"direction" form:"direction"`
	GPSStatus   int       `gorm:"type:tinyint;default:0" json:"gps_status" form:"gps_status"` // 0:无效 1:有效
	ReportTime  time.Time `gorm:"index" json:"report_time" form:"report_time"`
}

func (DevicePosition) TableName() string {
	return "wvp_device_position"
}
