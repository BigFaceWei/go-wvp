package service

import (
	"fmt"

	"wvp-go/server/global"
	"wvp-go/server/internal/gb28181/message"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
)

type DeviceControlService struct {
	logger          *zap.Logger
	controlHandler  *message.DeviceControlHandler
	ptzHandler      *message.PTZHandler
}

func NewDeviceControlService(logger *zap.Logger) *DeviceControlService {
	if logger == nil {
		logger = global.GVA_LOG
	}
	return &DeviceControlService{
		logger:         logger,
		controlHandler: message.NewDeviceControlHandler(logger),
		ptzHandler:     message.NewPTZHandler(logger),
	}
}

// sendControlMessage builds SIP MESSAGE XML and sends it to the device.
func (s *DeviceControlService) sendControlMessage(deviceID string, xmlBody []byte) error {
	srv := global.GVA_SIP_SERVER
	if srv == nil {
		return fmt.Errorf("SIP server not initialized")
	}

	device := &system.Device{}
	if err := global.GVA_DB.Where("device_id = ?", deviceID).First(device).Error; err != nil {
		return fmt.Errorf("device not found: %s", deviceID)
	}
	if !device.Online {
		return fmt.Errorf("device offline: %s", deviceID)
	}

	targetAddr := fmt.Sprintf("%s:%d", device.IP, device.Port)
	requestURI := fmt.Sprintf("sip:%s@%s", deviceID, global.GVA_CONFIG.WVP.SIP.Domain)
	headers := map[string]string{
		"To":           fmt.Sprintf("<sip:%s@%s>", deviceID, global.GVA_CONFIG.WVP.SIP.Domain),
		"Content-Type": "Application/MANSCDP+xml",
	}

	_, err := srv.SendRequestTo("MESSAGE", requestURI, targetAddr, headers, xmlBody)
	return err
}

func (s *DeviceControlService) RemoteBoot(deviceID string) error {
	xml, err := s.controlHandler.BuildDeviceControlXML(deviceID, "1", "TelebootCmd", "1")
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) RecordControl(deviceID, recordType string) error {
	xml, err := s.controlHandler.BuildRecordXML(deviceID, "1", recordType)
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) GuardControl(deviceID string, enabled bool) error {
	val := "0"
	if enabled {
		val = "1"
	}
	xml, err := s.controlHandler.BuildGuardXML(deviceID, "1", val)
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) HomePosition(deviceID string) error {
	xml, err := s.controlHandler.BuildHomePositionXML(deviceID, "1")
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) IFrame(deviceID string) error {
	xml, err := s.controlHandler.BuildIFrameXML(deviceID, "1")
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

// ----- Pre-set / Cruise / Scan -----

func (s *DeviceControlService) PresetSet(deviceID, presetID string) error {
	xml, err := s.ptzHandler.BuildPresetXML(deviceID, "1", presetID, "set")
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) PresetGoto(deviceID, presetID string) error {
	xml, err := s.ptzHandler.BuildPresetXML(deviceID, "1", presetID, "goto")
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) PresetRemove(deviceID, presetID string) error {
	xml, err := s.ptzHandler.BuildPresetXML(deviceID, "1", presetID, "remove")
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) CruiseStart(deviceID, cruiseID string) error {
	xml, err := s.ptzHandler.BuildCruiseXML(deviceID, "1", cruiseID, nil)
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) CruiseStop(deviceID, cruiseID string) error {
	xml, err := s.ptzHandler.BuildCruiseXML(deviceID, "1", cruiseID, nil)
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) ScanStart(deviceID, scanID string) error {
	xml, err := s.ptzHandler.BuildScanXML(deviceID, "1", scanID, "start")
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) ScanStop(deviceID, scanID string) error {
	xml, err := s.ptzHandler.BuildScanXML(deviceID, "1", scanID, "stop")
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

// ----- Voice Broadcast / Audio Broadcast -----

func (s *DeviceControlService) Broadcast(deviceID string) error {
	xml, err := s.controlHandler.BuildBroadcastXML(deviceID, "1")
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) AudioBroadcast(deviceID string) error {
	xml, err := s.controlHandler.BuildAudioBroadcastXML(deviceID, "1")
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

// ----- Drag Zoom -----

func (s *DeviceControlService) DragZoomIn(deviceID string, length, width, midpointX, midpointY, lengthX, lengthY string) error {
	xml, err := s.controlHandler.BuildDragZoomXML(deviceID, "1", length, width, midpointX, midpointY, lengthX, lengthY)
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}

func (s *DeviceControlService) DragZoomOut(deviceID string, length, width, midpointX, midpointY, lengthX, lengthY string) error {
	xml, err := s.controlHandler.BuildDragZoomOutXML(deviceID, "1", length, width, midpointX, midpointY, lengthX, lengthY)
	if err != nil {
		return err
	}
	return s.sendControlMessage(deviceID, xml)
}
