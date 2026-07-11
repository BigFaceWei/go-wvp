package request

import (
	"wvp-go/server/model/common/request"
	"wvp-go/server/model/system"
)

type SysLoginLogSearch struct {
	system.SysLoginLog
	request.PageInfo
}
