package response

type ErrorCode int

const (
	SUCCESS ErrorCode = 0

	INVALID_PARAMS ErrorCode = 400
	UNAUTHORIZED   ErrorCode = 401
	FORBIDDEN      ErrorCode = 403
	NOT_FOUND      ErrorCode = 404
	SERVER_ERROR   ErrorCode = 500

	DEVICE_NOT_FOUND       ErrorCode = 1001
	DEVICE_OFFLINE         ErrorCode = 1002
	DEVICE_ALREADY_EXISTS  ErrorCode = 1003
	DEVICE_REGISTER_FAILED ErrorCode = 1004

	CHANNEL_NOT_FOUND ErrorCode = 1101
	CHANNEL_BUSY      ErrorCode = 1102

	SIP_INIT_FAILED      ErrorCode = 2001
	SIP_SEND_FAILED      ErrorCode = 2002
	SIP_TIMEOUT          ErrorCode = 2003
	SIP_AUTH_FAILED      ErrorCode = 2004

	ZLM_NOT_CONNECTED     ErrorCode = 3001
	ZLM_COMMAND_FAILED    ErrorCode = 3002
	ZLM_STREAM_NOT_FOUND  ErrorCode = 3003
	ZLM_HOOK_FAILED       ErrorCode = 3004

	RECORD_START_FAILED  ErrorCode = 4001
	RECORD_STOP_FAILED   ErrorCode = 4002
	RECORD_NOT_FOUND     ErrorCode = 4003
	RECORD_STORAGE_ERROR ErrorCode = 4004

	PLATFORM_NOT_FOUND    ErrorCode = 5001
	PLATFORM_REGISTER_ERR ErrorCode = 5002
	PLATFORM_CATALOG_ERR  ErrorCode = 5003

	DB_ERROR          ErrorCode = 6001
	DB_RECORD_NOT_FOUND ErrorCode = 6002
	DB_DUPLICATE      ErrorCode = 6003
)

var errorCodeMessages = map[ErrorCode]string{
	SUCCESS: "成功",
	INVALID_PARAMS:   "请求参数无效",
	UNAUTHORIZED:     "未授权",
	FORBIDDEN:        "禁止访问",
	NOT_FOUND:        "资源不存在",
	SERVER_ERROR:     "服务器内部错误",
	DEVICE_NOT_FOUND:       "设备不存在",
	DEVICE_OFFLINE:         "设备离线",
	DEVICE_ALREADY_EXISTS:  "设备已存在",
	DEVICE_REGISTER_FAILED: "设备注册失败",
	CHANNEL_NOT_FOUND: "通道不存在",
	CHANNEL_BUSY:      "通道繁忙",
	SIP_INIT_FAILED:      "SIP服务初始化失败",
	SIP_SEND_FAILED:      "SIP消息发送失败",
	SIP_TIMEOUT:          "SIP请求超时",
	SIP_AUTH_FAILED:      "SIP认证失败",
	ZLM_NOT_CONNECTED:     "ZLM服务未连接",
	ZLM_COMMAND_FAILED:    "ZLM命令执行失败",
	ZLM_STREAM_NOT_FOUND:  "流媒体流不存在",
	ZLM_HOOK_FAILED:       "ZLM回调处理失败",
	RECORD_START_FAILED:  "录制启动失败",
	RECORD_STOP_FAILED:   "录制停止失败",
	RECORD_NOT_FOUND:     "录制文件不存在",
	RECORD_STORAGE_ERROR: "录制存储错误",
	PLATFORM_NOT_FOUND:    "平台不存在",
	PLATFORM_REGISTER_ERR: "平台注册失败",
	PLATFORM_CATALOG_ERR:  "平台目录同步失败",
	DB_ERROR:          "数据库错误",
	DB_RECORD_NOT_FOUND: "数据库记录不存在",
	DB_DUPLICATE:      "数据重复",
}

func (e ErrorCode) Message() string {
	if msg, ok := errorCodeMessages[e]; ok {
		return msg
	}
	return "未知错误"
}

func (e ErrorCode) HTTPStatus() int {
	switch {
	case e == SUCCESS:
		return 200
	case e >= 1000 && e < 2000:
		return 400
	case e >= 2000 && e < 3000:
		return 400
	case e >= 3000 && e < 4000:
		return 502
	case e >= 4000 && e < 5000:
		return 500
	case e >= 5000 && e < 6000:
		return 502
	case e >= 6000:
		return 500
	default:
		return 200
	}
}