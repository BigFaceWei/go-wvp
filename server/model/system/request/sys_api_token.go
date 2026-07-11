package request

import (
	"wvp-go/server/model/common/request"
	"wvp-go/server/model/system"
)

type SysApiTokenSearch struct {
	system.SysApiToken
	request.PageInfo
    Status *bool `json:"status" form:"status"`
}
