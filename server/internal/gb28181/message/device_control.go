package message

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

// DeviceControlHandler handles building XML for GB28181 DeviceControl commands.
type DeviceControlHandler struct {
	logger *zap.Logger
}

func NewDeviceControlHandler(logger *zap.Logger) *DeviceControlHandler {
	return &DeviceControlHandler{logger: logger}
}

// BuildDeviceControlXML builds a generic DeviceControl command XML.
// cmdName: the sub-command name (e.g. "TelebootCmd", "RecordCmd", "GuardCmd")
// cmdValue: the sub-command value (e.g. "1" for on, "0" for off)
func (h *DeviceControlHandler) BuildDeviceControlXML(deviceID, sn, cmdName, cmdValue string) ([]byte, error) {
	xml := fmt.Sprintf(`<?xml version="1.0" encoding="GB2312"?>
<Control>
  <CmdType>DeviceControl</CmdType>
  <SN>%s</SN>
  <DeviceID>%s</DeviceID>
  <%s>%s</%s>
</Control>`, sn, deviceID, cmdName, cmdValue, cmdName)

	return []byte(xml), nil
}

// BuildHomePositionXML builds a HomePosition command (reset to home).
func (h *DeviceControlHandler) BuildHomePositionXML(deviceID, sn string) ([]byte, error) {
	return h.BuildDeviceControlXML(deviceID, sn, "HomePositionCmd", "1")
}

// BuildGuardXML builds a Guard command for arming/disarming.
// enabled: "1" for arming (布防), "0" for disarming (撤防)
func (h *DeviceControlHandler) BuildGuardXML(deviceID, sn, enabled string) ([]byte, error) {
	return h.BuildDeviceControlXML(deviceID, sn, "GuardCmd", enabled)
}

// BuildRecordXML builds a Record command for remote recording control.
// recordType: "ManualRecord" / "TimingRecord" / "ManualStopRecord"
func (h *DeviceControlHandler) BuildRecordXML(deviceID, sn, recordType string) ([]byte, error) {
	return h.BuildDeviceControlXML(deviceID, sn, "RecordCmd", recordType)
}

// BuildIFrameXML builds an I-Frame command (key frame request).
func (h *DeviceControlHandler) BuildIFrameXML(deviceID, sn string) ([]byte, error) {
	return h.BuildDeviceControlXML(deviceID, sn, "IFrameCmd", "1")
}

// BuildDragZoomXML builds a drag zoom command.
// Parameters are per GB28181 DragZoomIn/DragZoomOut.
func (h *DeviceControlHandler) BuildDragZoomXML(deviceID, sn string, length, width, midpointX, midpointY, lengthX, lengthY string) ([]byte, error) {
	xml := fmt.Sprintf(`<?xml version="1.0" encoding="GB2312"?>
<Control>
  <CmdType>DeviceControl</CmdType>
  <SN>%s</SN>
  <DeviceID>%s</DeviceID>
  <DragZoomIn>
    <Length>%s</Length>
    <Width>%s</Width>
    <MidPointX>%s</MidPointX>
    <MidPointY>%s</MidPointY>
    <LengthX>%s</LengthX>
    <LengthY>%s</LengthY>
  </DragZoomIn>
</Control>`, sn, deviceID, length, width, midpointX, midpointY, lengthX, lengthY)

	return []byte(xml), nil
}

// BuildDragZoomOutXML builds a drag zoom out command.
func (h *DeviceControlHandler) BuildDragZoomOutXML(deviceID, sn string, length, width, midpointX, midpointY, lengthX, lengthY string) ([]byte, error) {
	xml := fmt.Sprintf(`<?xml version="1.0" encoding="GB2312"?>
<Control>
  <CmdType>DeviceControl</CmdType>
  <SN>%s</SN>
  <DeviceID>%s</DeviceID>
  <DragZoomOut>
    <Length>%s</Length>
    <Width>%s</Width>
    <MidPointX>%s</MidPointX>
    <MidPointY>%s</MidPointY>
    <LengthX>%s</LengthX>
    <LengthY>%s</LengthY>
  </DragZoomOut>
</Control>`, sn, deviceID, length, width, midpointX, midpointY, lengthX, lengthY)

	return []byte(xml), nil
}

// BuildBroadcastXML builds a Broadcast command (voice broadcast/alarm triggered broadcast).
func (h *DeviceControlHandler) BuildBroadcastXML(deviceID, sn string) ([]byte, error) {
	return h.BuildDeviceControlXML(deviceID, sn, "BroadcastCmd", "1")
}

// BuildAudioBroadcastXML builds an AudioBroadcast command (two-way talk/voice intercom).
func (h *DeviceControlHandler) BuildAudioBroadcastXML(deviceID, sn string) ([]byte, error) {
	return h.BuildDeviceControlXML(deviceID, sn, "AudioBroadcastCmd", "1")
}

// ExtractCmdTypeFromXML extracts the CmdType value from a DeviceControl XML response.
func (h *DeviceControlHandler) ExtractCmdTypeFromXML(body []byte) string {
	content := string(body)
	start := strings.Index(content, "<CmdType>")
	if start == -1 {
		return ""
	}
	start += len("<CmdType>")
	end := strings.Index(content[start:], "</CmdType>")
	if end == -1 {
		return ""
	}
	return content[start : start+end]
}
