package system

import (
	"time"

	"wvp-go/server/global"
)

type Alarm struct {
	global.GVA_MODEL
	DeviceID     string    `gorm:"type:varchar(64);index;not null" json:"device_id" form:"device_id"`
	ChannelID    string    `gorm:"type:varchar(64);index;not null" json:"channel_id" form:"channel_id"`
	AlarmType    string    `gorm:"type:varchar(32);index;not null" json:"alarm_type" form:"alarm_type"`
	AlarmTypeVal int       `gorm:"type:int" json:"alarm_type_val" form:"alarm_type_val"`
	AlarmPriority int      `gorm:"type:tinyint;default:0" json:"alarm_priority" form:"alarm_priority"`
	AlarmTime    time.Time `gorm:"index" json:"alarm_time" form:"alarm_time"`
	Longitude    float64   `gorm:"type:decimal(10,7)" json:"longitude" form:"longitude"`
	Latitude     float64   `gorm:"type:decimal(10,7)" json:"latitude" form:"latitude"`
	Description  string    `gorm:"type:varchar(512)" json:"description" form:"description"`
	HandleStatus int       `gorm:"type:tinyint;default:0" json:"handle_status" form:"handle_status"` // 0:未处理 1:已处理 2:已忽略
	HandleUser   string    `gorm:"type:varchar(64)" json:"handle_user" form:"handle_user"`
	HandleTime   time.Time `json:"handle_time" form:"handle_time"`
	HandleResult string    `gorm:"type:varchar(256)" json:"handle_result" form:"handle_result"`
}

func (Alarm) TableName() string {
	return "wvp_alarm"
}
