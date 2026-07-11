package message

import (
	"encoding/xml"
	"fmt"

	"wvp-go/server/global"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
)

type CatalogHandler struct {
	logger *zap.Logger
}

func NewCatalogHandler(logger *zap.Logger) *CatalogHandler {
	return &CatalogHandler{logger: logger}
}

type CatalogRequest struct {
	XMLName   xml.Name `xml:"Query"`
	CmdType   string   `xml:"CmdType"`
	SN        string   `xml:"SN"`
	DeviceID  string   `xml:"DeviceID"`
	StartPos  int      `xml:"StartNum"`
	MaxCount  int      `xml:"MaxNum"`
}

type CatalogResponse struct {
	XMLName    xml.Name        `xml:"Response"`
	CmdType    string          `xml:"CmdType"`
	SN         string          `xml:"SN"`
	DeviceID   string          `xml:"DeviceID"`
	SumNum     int             `xml:"SumNum"`
	DeviceList CatalogDeviceList `xml:"DeviceList"`
}

type CatalogDeviceList struct {
	Num   int              `xml:"Num,attr"`
	Items []CatalogItem    `xml:"Item"`
}

type CatalogItem struct {
	XMLName        xml.Name `xml:"Item"`
	DeviceID       string   `xml:"DeviceID"`
	Name           string   `xml:"Name"`
	Manufacturer   string   `xml:"DeviceManufacturer"`
	Model          string   `xml:"DeviceModel"`
	Owner          string   `xml:"Owner"`
	CivilCode      string   `xml:"CivilCode"`
	Address        string   `xml:"Address"`
	Parental       int      `xml:"Parental"`
	ParentID       string   `xml:"ParentID"`
	SafetyWay      int      `xml:"SafetyWay"`
	RegisterWay    int      `xml:"RegisterWay"`
	Secrecy        int      `xml:"Secrecy"`
	Status         string   `xml:"Status"`
}

func (h *CatalogHandler) ParseCatalogRequest(body []byte) (*CatalogRequest, error) {
	req := &CatalogRequest{}
	if err := xml.Unmarshal(body, req); err != nil {
		return nil, fmt.Errorf("parse catalog request failed: %w", err)
	}

	if req.CmdType != "Catalog" {
		return nil, fmt.Errorf("invalid cmd type: %s", req.CmdType)
	}

	return req, nil
}

func (h *CatalogHandler) HandleCatalogRequest(req *CatalogRequest) (*CatalogResponse, error) {
	var channels []system.DeviceChannel
	result := global.GVA_DB.Where("device_id = ?", req.DeviceID).
		Offset(req.StartPos).
		Limit(req.MaxCount).
		Find(&channels)

	if result.Error != nil {
		return nil, fmt.Errorf("query channels failed: %w", result.Error)
	}

	var total int64
	global.GVA_DB.Model(&system.DeviceChannel{}).Where("device_id = ?", req.DeviceID).Count(&total)

	items := make([]CatalogItem, len(channels))
	for i, ch := range channels {
		status := "OFF"
		if ch.Status == "ON" {
			status = "ON"
		}

		items[i] = CatalogItem{
			DeviceID:     ch.ChannelID,
			Name:         ch.Name,
			Manufacturer: ch.Manufacturer,
			Model:        ch.Model,
			Owner:        ch.Owner,
			CivilCode:    ch.CivilCode,
			Address:      ch.Address,
			Parental:     ch.Parental,
			ParentID:     ch.ParentID,
			SafetyWay:    ch.SafetyWay,
			RegisterWay:  ch.RegisterWay,
			Secrecy:      ch.Secrecy,
			Status:       status,
		}
	}

	response := &CatalogResponse{
		CmdType:  "Catalog",
		SN:       req.SN,
		DeviceID: req.DeviceID,
		SumNum:   int(total),
		DeviceList: CatalogDeviceList{
			Num:   len(items),
			Items: items,
		},
	}

	h.logger.Info("Catalog query handled",
		zap.String("device_id", req.DeviceID),
		zap.Int("channels", len(items)),
		zap.Int("total", int(total)),
	)

	return response, nil
}

func (h *CatalogHandler) BuildCatalogXML(resp *CatalogResponse) ([]byte, error) {
	output, err := xml.MarshalIndent(resp, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal catalog response failed: %w", err)
	}

	return append([]byte(xml.Header), output...), nil
}

func (h *CatalogHandler) SendCatalogQuery(deviceID string) error {
	device := &system.Device{}
	result := global.GVA_DB.Where("device_id = ?", deviceID).First(device)
	if result.Error != nil {
		return fmt.Errorf("device not found: %s", deviceID)
	}

	if !device.Online {
		return fmt.Errorf("device offline: %s", deviceID)
	}

	h.logger.Info("Send catalog query",
		zap.String("device_id", deviceID),
	)

	return nil
}