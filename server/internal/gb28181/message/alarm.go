package message

import (
	"encoding/xml"
	"fmt"
	"time"

	"wvp-go/server/global"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
)

type AlarmMessage struct {
	XMLName      xml.Name `xml:"Notify"`
	CmdType      string   `xml:"CmdType"`
	SN           string   `xml:"SN"`
	DeviceID     string   `xml:"DeviceID"`
	AlarmType    string   `xml:"AlarmType"`
	AlarmTypeVal int      `xml:"AlarmTypeVal"`
	AlarmPriority int     `xml:"AlarmPriority"`
	AlarmTime    string   `xml:"AlarmTime"`
	Longitude    float64  `xml:"Longitude"`
	Latitude     float64  `xml:"Latitude"`
	Description  string   `xml:"Description"`
}

type AlarmHandler struct {
	logger *zap.Logger
}

func NewAlarmHandler(logger *zap.Logger) *AlarmHandler {
	return &AlarmHandler{
		logger: logger,
	}
}

func (h *AlarmHandler) HandleAlarm(deviceID string, data []byte) error {
	var alarmMsg AlarmMessage
	if err := xml.Unmarshal(data, &alarmMsg); err != nil {
		return fmt.Errorf("parse alarm message failed: %w", err)
	}

	alarmTime, _ := time.Parse("2006-01-02T15:04:05", alarmMsg.AlarmTime)
	if alarmTime.IsZero() {
		alarmTime = time.Now()
	}

	alarm := system.Alarm{
		DeviceID:      deviceID,
		ChannelID:     alarmMsg.DeviceID,
		AlarmType:     alarmMsg.AlarmType,
		AlarmTypeVal:  alarmMsg.AlarmTypeVal,
		AlarmPriority: alarmMsg.AlarmPriority,
		AlarmTime:     alarmTime,
		Longitude:     alarmMsg.Longitude,
		Latitude:      alarmMsg.Latitude,
		Description:   alarmMsg.Description,
		HandleStatus:  0,
	}

	if err := global.GVA_DB.Create(&alarm).Error; err != nil {
		h.logger.Error("save alarm failed", zap.Error(err))
		return err
	}

	h.logger.Info("Alarm received",
		zap.String("device_id", deviceID),
		zap.String("alarm_type", alarmMsg.AlarmType),
		zap.Int("alarm_priority", alarmMsg.AlarmPriority),
		zap.Time("alarm_time", alarmTime),
	)

	return nil
}

func (h *AlarmHandler) GetAlarmTypeString(alarmTypeVal int) string {
	switch alarmTypeVal {
	case 1:
		return "视频丢失"
	case 2:
		return "设备离线"
	case 3:
		return "硬盘故障"
	case 4:
		return "视频遮挡"
	case 5:
		return "移动侦测"
	case 6:
		return "视频帧丢失"
	case 7:
		return "IO报警"
	default:
		return fmt.Sprintf("未知(%d)", alarmTypeVal)
	}
}
