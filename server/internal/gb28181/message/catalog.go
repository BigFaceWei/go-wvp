package message

import (
	"encoding/xml"
	"fmt"
	"sync"
	"time"

	"wvp-go/server/global"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
)

const (
	catalogDedupWindow = 10 * time.Second
	catalogCleanupInterval = 60 * time.Second
)

type catalogDedupEntry struct {
	processedAt time.Time
}

type CatalogHandler struct {
	logger *zap.Logger
	dedup  sync.Map
}

func NewCatalogHandler(logger *zap.Logger) *CatalogHandler {
	h := &CatalogHandler{logger: logger}
	go h.cleanupLoop()
	return h
}

func (h *CatalogHandler) cleanupLoop() {
	ticker := time.NewTicker(catalogCleanupInterval)
	defer ticker.Stop()
	for range ticker.C {
		h.dedup.Range(func(key, value any) bool {
			entry := value.(*catalogDedupEntry)
			if time.Since(entry.processedAt) > catalogCleanupInterval {
				h.dedup.Delete(key)
			}
			return true
		})
	}
}

func (h *CatalogHandler) isDuplicate(deviceID, sn string) bool {
	key := deviceID + ":" + sn
	if v, ok := h.dedup.Load(key); ok {
		entry := v.(*catalogDedupEntry)
		if time.Since(entry.processedAt) < catalogDedupWindow {
			return true
		}
	}
	h.dedup.Store(key, &catalogDedupEntry{processedAt: time.Now()})
	return false
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
	Manufacturer   string   `xml:"Manufacturer"`
	Model          string   `xml:"Model"`
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

func (h *CatalogHandler) ParseCatalogResponse(body []byte) (*CatalogResponse, error) {
	resp := &CatalogResponse{}
	if err := xml.Unmarshal(body, resp); err != nil {
		return nil, fmt.Errorf("parse catalog response failed: %w", err)
	}

	if resp.CmdType != "Catalog" {
		return nil, fmt.Errorf("invalid cmd type: %s", resp.CmdType)
	}

	return resp, nil
}

func (h *CatalogHandler) HandleCatalogResponse(resp *CatalogResponse) error {
	deviceID := resp.DeviceID
	channelCount := len(resp.DeviceList.Items)

	h.logger.Info("Catalog response received",
		zap.String("device_id", deviceID),
		zap.String("sn", resp.SN),
		zap.Int("sum_num", resp.SumNum),
		zap.Int("channel_count", channelCount),
	)

	if h.isDuplicate(deviceID, resp.SN) {
		h.logger.Info("Duplicate catalog response skipped, 200 OK will be sent",
			zap.String("device_id", deviceID),
			zap.String("sn", resp.SN),
		)
		return nil
	}

	for _, item := range resp.DeviceList.Items {
		var existing system.DeviceChannel
		result := global.GVA_DB.Where("device_id = ? AND channel_id = ?", deviceID, item.DeviceID).First(&existing)
		
		if result.Error == nil {
			existing.Name = item.Name
			existing.Manufacturer = item.Manufacturer
			existing.Model = item.Model
			existing.Owner = item.Owner
			existing.CivilCode = item.CivilCode
			existing.Address = item.Address
			existing.Parental = item.Parental
			existing.ParentID = item.ParentID
			existing.SafetyWay = item.SafetyWay
			existing.RegisterWay = item.RegisterWay
			existing.Secrecy = item.Secrecy
			existing.Status = item.Status
			global.GVA_DB.Save(&existing)
		} else {
			channel := system.DeviceChannel{
				DeviceID:     deviceID,
				ChannelID:    item.DeviceID,
				Name:         item.Name,
				Manufacturer: item.Manufacturer,
				Model:        item.Model,
				Owner:        item.Owner,
				CivilCode:    item.CivilCode,
				Address:      item.Address,
				Parental:     item.Parental,
				ParentID:     item.ParentID,
				SafetyWay:    item.SafetyWay,
				RegisterWay:  item.RegisterWay,
				Secrecy:      item.Secrecy,
				Status:       item.Status,
			}
			if err := global.GVA_DB.Create(&channel).Error; err != nil {
				h.logger.Error("create channel failed",
					zap.String("device_id", deviceID),
					zap.String("channel_id", item.DeviceID),
					zap.Error(err),
				)
				continue
			}
		}
	}

	h.logger.Info("Catalog response handled",
		zap.String("device_id", deviceID),
		zap.Int("channels", len(resp.DeviceList.Items)),
	)

	return nil
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