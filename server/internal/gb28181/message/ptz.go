package message

import (
	"encoding/xml"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type PTZHandler struct {
	logger *zap.Logger
}

func NewPTZHandler(logger *zap.Logger) *PTZHandler {
	return &PTZHandler{logger: logger}
}

type PTZCommand struct {
	CmdType   string `xml:"CmdType"`
	SN        string `xml:"SN"`
	DeviceID  string `xml:"DeviceID"`
	PTZCmd    string `xml:"PTZCmd"`
}

type PTZCmdCode struct {
	XMLName  xml.Name `xml:"Control"`
	CmdType  string   `xml:"CmdType"`
	SN       string   `xml:"SN"`
	DeviceID string   `xml:"DeviceID"`
	PTZCmd   string   `xml:"PTZCmd"`
}

type PTZAction struct {
	Direction string
	Speed     int
	Zoom      int
	Focus     int
	Iris      int
}

const (
	PTZStop        = 0x00
	PTZUp          = 0x01
	PTZDown        = 0x02
	PTZLeft        = 0x04
	PTZRight       = 0x08
	PTZUpLeft      = 0x05
	PTZUpRight     = 0x09
	PTZDownLeft    = 0x06
	PTZDownRight   = 0x0A
	PTZZoomIn      = 0x10
	PTZZoomOut     = 0x20
	PTZFocusNear   = 0x40
	PTZFocusFar    = 0x80
	PTZIrisOpen    = 0x100
	PTZIrisClose   = 0x200
	PTZAuto        = 0x400
)

var ptzDirectionMap = map[string]int{
	"up":         PTZUp,
	"down":       PTZDown,
	"left":       PTZLeft,
	"right":      PTZRight,
	"up_left":    PTZUpLeft,
	"up_right":   PTZUpRight,
	"down_left":  PTZDownLeft,
	"down_right": PTZDownRight,
	"stop":       PTZStop,
}

func (h *PTZHandler) BuildPTZCommand(deviceID, direction string, speed int, zoom int) (string, error) {
	cmdCode := 0

	if dirCode, ok := ptzDirectionMap[strings.ToLower(direction)]; ok {
		cmdCode |= dirCode
	} else {
		return "", fmt.Errorf("invalid direction: %s", direction)
	}

	if speed < 0 {
		speed = 0
	}
	if speed > 255 {
		speed = 255
	}

	if zoom > 0 {
		cmdCode |= PTZZoomIn
	} else if zoom < 0 {
		cmdCode |= PTZZoomOut
	}

	ptzCmd := fmt.Sprintf("%02X%02X%02X%02X%02X%02X%02X%02X",
		0xFF,
		(cmdCode>>8)&0xFF,
		cmdCode&0xFF,
		speed,
		speed,
		0x00,
		0x00,
		0x00,
	)

	return ptzCmd, nil
}

// Common DeviceControl XML builder helper
func buildDeviceControlXML(cmdType, deviceID, sn string, extraFields map[string]string) ([]byte, error) {
	var sb strings.Builder
	sb.WriteString(xml.Header)
	sb.WriteString(fmt.Sprintf("<Control>\n"))
	sb.WriteString(fmt.Sprintf("  <CmdType>%s</CmdType>\n", cmdType))
	sb.WriteString(fmt.Sprintf("  <SN>%s</SN>\n", sn))
	sb.WriteString(fmt.Sprintf("  <DeviceID>%s</DeviceID>\n", deviceID))
	for k, v := range extraFields {
		sb.WriteString(fmt.Sprintf("  <%s>%s</%s>\n", k, v, k))
	}
	sb.WriteString("</Control>")
	return []byte(sb.String()), nil
}

func (h *PTZHandler) BuildPTZXML(deviceID, ptzCmd string, sn string) ([]byte, error) {
	cmd := PTZCmdCode{
		CmdType:  "DeviceControl",
		SN:       sn,
		DeviceID: deviceID,
		PTZCmd:   ptzCmd,
	}

	output, err := xml.MarshalIndent(cmd, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal PTZ command failed: %w", err)
	}

	return append([]byte(xml.Header), output...), nil
}

func (h *PTZHandler) BuildPresetXML(deviceID, sn, presetID, cmdType string) ([]byte, error) {
	// cmdType: "set" / "goto" / "remove"
	return buildDeviceControlXML("DeviceControl", deviceID, sn, map[string]string{
		"PresetCmd": cmdType,
		"PresetID":  presetID,
	})
}

func (h *PTZHandler) BuildCruiseXML(deviceID, sn, cruiseID string, cruisePoints []string) ([]byte, error) {
	// cmdType: "start" / "stop" / "add"
	return buildDeviceControlXML("DeviceControl", deviceID, sn, map[string]string{
		"CruiseCmd": cruiseID,
	})
}

func (h *PTZHandler) BuildScanXML(deviceID, sn, scanID, cmdType string) ([]byte, error) {
	return buildDeviceControlXML("DeviceControl", deviceID, sn, map[string]string{
		"ScanCmd": cmdType,
		"ScanID":  scanID,
	})
}

func (h *PTZHandler) ParsePTZCommand(body []byte) (*PTZCommand, error) {
	cmd := &PTZCommand{}
	if err := xml.Unmarshal(body, cmd); err != nil {
		return nil, fmt.Errorf("parse PTZ command failed: %w", err)
	}

	if cmd.CmdType != "DeviceControl" {
		return nil, fmt.Errorf("invalid command type: %s", cmd.CmdType)
	}

	return cmd, nil
}

func (h *PTZHandler) HandlePTZCommand(deviceID, direction string, speed, zoom int) error {
	ptzCmd, err := h.BuildPTZCommand(deviceID, direction, speed, zoom)
	if err != nil {
		return err
	}

	h.logger.Info("PTZ command built",
		zap.String("device_id", deviceID),
		zap.String("direction", direction),
		zap.Int("speed", speed),
		zap.Int("zoom", zoom),
		zap.String("cmd", ptzCmd),
	)

	return nil
}

// HandlePreset is kept for backward compatibility — use BuildPresetXML for building XML.
func (h *PTZHandler) HandlePreset(deviceID, presetID string) error {
	h.logger.Info("Preset command",
		zap.String("device_id", deviceID),
		zap.String("preset_id", presetID),
	)
	return nil
}

// HandleCruise is kept for backward compatibility — use BuildCruiseXML for building XML.
func (h *PTZHandler) HandleCruise(deviceID, cruiseID string, cruisePoints []string) error {
	h.logger.Info("Cruise command",
		zap.String("device_id", deviceID),
		zap.String("cruise_id", cruiseID),
		zap.Int("points", len(cruisePoints)),
	)
	return nil
}
