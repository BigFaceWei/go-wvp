package request

import (
	"wvp-go/server/model/common/request"
	"wvp-go/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
