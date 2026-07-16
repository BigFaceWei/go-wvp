package system

import (
	"time"

	"wvp-go/server/global"
)

type Device struct {
	global.GVA_MODEL
	DeviceID      string    `gorm:"type:varchar(64);uniqueIndex;not null" json:"device_id" form:"device_id"`
	Name          string    `gorm:"type:varchar(128)" json:"name" form:"name"`
	Manufacturer  string    `gorm:"type:varchar(128)" json:"manufacturer" form:"manufacturer"`
	Model         string    `gorm:"type:varchar(128)" json:"model" form:"model"`
	Owner         string    `gorm:"type:varchar(128)" json:"owner" form:"owner"`
	CivilCode     string    `gorm:"type:varchar(32)" json:"civil_code" form:"civil_code"`
	Address       string    `gorm:"type:varchar(256)" json:"address" form:"address"`
	Parental      int       `gorm:"type:tinyint;default:0" json:"parental" form:"parental"`
	ParentID      string    `gorm:"type:varchar(64)" json:"parent_id" form:"parent_id"`
	SafetyWay     int       `gorm:"type:tinyint;default:0" json:"safety_way" form:"safety_way"`
	RegisterWay   int       `gorm:"type:tinyint;default:1" json:"register_way" form:"register_way"`
	Secrecy       int       `gorm:"type:tinyint;default:0" json:"secrecy" form:"secrecy"`
	Password      string    `gorm:"type:varchar(128)" json:"password" form:"password"`
	Transport     string    `gorm:"type:varchar(8);default:'UDP'" json:"transport" form:"transport"`
	IP            string    `gorm:"type:varchar(64)" json:"ip" form:"ip"`
	Port          int       `gorm:"type:int;default:5060" json:"port" form:"port"`
	Online        bool      `gorm:"type:tinyint(1);default:0" json:"on_line" form:"on_line"`
	RegisterTime  time.Time `json:"register_time" form:"register_time"`
	KeepaliveTime time.Time `json:"keepalive_time" form:"keepalive_time"`
	ChannelCount  int       `gorm:"type:int;default:0" json:"channel_count" form:"channel_count"`
}

func (Device) TableName() string {
	return "wvp_device"
}

type DeviceChannel struct {
	global.GVA_MODEL
	DeviceID       string `gorm:"type:varchar(64);index;not null" json:"device_id" form:"device_id"`
	ChannelID      string `gorm:"type:varchar(64);uniqueIndex;not null" json:"channel_id" form:"channel_id"`
	ParentID       string `gorm:"type:varchar(64);index" json:"parent_id" form:"parent_id"`
	Name           string `gorm:"type:varchar(128)" json:"name" form:"name"`
	Manufacturer   string `gorm:"type:varchar(128)" json:"manufacturer" form:"manufacturer"`
	Model          string `gorm:"type:varchar(128)" json:"model" form:"model"`
	Owner          string `gorm:"type:varchar(128)" json:"owner" form:"owner"`
	CivilCode      string `gorm:"type:varchar(32)" json:"civil_code" form:"civil_code"`
	Address        string `gorm:"type:varchar(256)" json:"address" form:"address"`
	Parental       int    `gorm:"type:tinyint;default:0" json:"parental" form:"parental"`
	SafetyWay      int    `gorm:"type:tinyint;default:0" json:"safety_way" form:"safety_way"`
	RegisterWay    int    `gorm:"type:tinyint;default:1" json:"register_way" form:"register_way"`
	Secrecy        int    `gorm:"type:tinyint;default:0" json:"secrecy" form:"secrecy"`
	Status         string `gorm:"type:varchar(16);default:'OFF'" json:"status" form:"status"`
	HasAudio       bool   `gorm:"type:tinyint(1);default:0" json:"has_audio" form:"has_audio"`
	Longitude      float64 `gorm:"type:decimal(10,7)" json:"longitude" form:"longitude"`
	Latitude       float64 `gorm:"type:decimal(10,7)" json:"latitude" form:"latitude"`
}

func (DeviceChannel) TableName() string {
	return "wvp_device_channel"
}