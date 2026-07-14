package system

import (
	"wvp-go/server/global"
)

// Region 行政区域
type Region struct {
	global.GVA_MODEL
	Code     string `gorm:"type:varchar(64);primaryKey;not null" json:"code" form:"code"`
	Name     string `gorm:"type:varchar(128);not null" json:"name" form:"name"`
	ParentID string `gorm:"type:varchar(64);index;default:''" json:"parent_id" form:"parent_id"`
	Level    int    `gorm:"type:int;default:0" json:"level" form:"level"` // 0-省,1-市,2-区县,3-街道
	Sort     int    `gorm:"type:int;default:0" json:"sort" form:"sort"`
	Remark   string `gorm:"type:varchar(256)" json:"remark" form:"remark"`
}

func (Region) TableName() string {
	return "wvp_region"
}
