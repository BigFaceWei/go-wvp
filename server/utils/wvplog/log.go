package wvplog

import (
	"wvp-go/server/global"

	"go.uber.org/zap"
)

func Info(msg string, fields ...zap.Field) {
	global.GVA_LOG.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	global.GVA_LOG.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	global.GVA_LOG.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	global.GVA_LOG.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	global.GVA_LOG.Fatal(msg, fields...)
}

func With(fields ...zap.Field) *zap.Logger {
	return global.GVA_LOG.With(fields...)
}

func DeviceLog(deviceID string) *zap.Logger {
	return global.GVA_LOG.With(zap.String("device_id", deviceID))
}

func ChannelLog(deviceID, channelID string) *zap.Logger {
	return global.GVA_LOG.With(
		zap.String("device_id", deviceID),
		zap.String("channel_id", channelID),
	)
}

func SIPLog(direction, method string) *zap.Logger {
	return global.GVA_LOG.With(
		zap.String("sip_direction", direction),
		zap.String("sip_method", method),
	)
}

func ZLMLog(serverID string) *zap.Logger {
	return global.GVA_LOG.With(zap.String("zlm_server", serverID))
}