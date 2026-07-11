package message

import (
	"encoding/xml"
	"fmt"
	"time"

	"wvp-go/server/global"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
)

type MobilePositionMessage struct {
	XMLName     xml.Name `xml:"Notify"`
	CmdType     string   `xml:"CmdType"`
	SN          string   `xml:"SN"`
	DeviceID    string   `xml:"DeviceID"`
	Longitude   float64  `xml:"Longitude"`
	Latitude    float64  `xml:"Latitude"`
	Altitude    float64  `xml:"Altitude"`
	Speed       float64  `xml:"Speed"`
	Direction   float64  `xml:"Direction"`
	GPSStatus   int      `xml:"GPSStatus"`
	ReportTime  string   `xml:"ReportTime"`
}

type PositionHandler struct {
	logger *zap.Logger
}

func NewPositionHandler(logger *zap.Logger) *PositionHandler {
	return &PositionHandler{
		logger: logger,
	}
}

func (h *PositionHandler) HandleMobilePosition(deviceID string, data []byte) error {
	var posMsg MobilePositionMessage
	if err := xml.Unmarshal(data, &posMsg); err != nil {
		return fmt.Errorf("parse mobile position message failed: %w", err)
	}

	reportTime, _ := time.Parse("2006-01-02T15:04:05", posMsg.ReportTime)
	if reportTime.IsZero() {
		reportTime = time.Now()
	}

	position := system.DevicePosition{
		DeviceID:   deviceID,
		ChannelID:  posMsg.DeviceID,
		Longitude:  posMsg.Longitude,
		Latitude:   posMsg.Latitude,
		Altitude:   posMsg.Altitude,
		Speed:      posMsg.Speed,
		Direction:  posMsg.Direction,
		GPSStatus:  posMsg.GPSStatus,
		ReportTime: reportTime,
	}

	var existing system.DevicePosition
	result := global.GVA_DB.Where("device_id = ? AND channel_id = ?", deviceID, posMsg.DeviceID).First(&existing)
	if result.Error == nil {
		existing.Longitude = posMsg.Longitude
		existing.Latitude = posMsg.Latitude
		existing.Altitude = posMsg.Altitude
		existing.Speed = posMsg.Speed
		existing.Direction = posMsg.Direction
		existing.GPSStatus = posMsg.GPSStatus
		existing.ReportTime = reportTime
		if err := global.GVA_DB.Save(&existing).Error; err != nil {
			h.logger.Error("update device position failed", zap.Error(err))
			return err
		}
	} else {
		if err := global.GVA_DB.Create(&position).Error; err != nil {
			h.logger.Error("save device position failed", zap.Error(err))
			return err
		}
	}

	h.logger.Debug("Mobile position received",
		zap.String("device_id", deviceID),
		zap.Float64("longitude", posMsg.Longitude),
		zap.Float64("latitude", posMsg.Latitude),
	)

	return nil
}
