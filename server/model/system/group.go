package system

import (
	"wvp-go/server/global"
)

// BusinessGroup 业务分组
type BusinessGroup struct {
	global.GVA_MODEL
	Name     string `gorm:"type:varchar(128);not null" json:"name" form:"name"`
	ParentID string `gorm:"type:varchar(64);index;default:''" json:"parent_id" form:"parent_id"`
	DeviceID string `gorm:"type:varchar(64)" json:"device_id" form:"device_id"`
	RegionID string `gorm:"type:varchar(64)" json:"region_id" form:"region_id"`
	Sort     int    `gorm:"type:int;default:0" json:"sort" form:"sort"`
	Remark   string `gorm:"type:varchar(256)" json:"remark" form:"remark"`
}

func (BusinessGroup) TableName() string {
	return "wvp_group"
}
