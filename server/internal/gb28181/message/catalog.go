package message

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"wvp-go/server/global"
	"wvp-go/server/model/system"

	"go.uber.org/zap"
	"golang.org/x/text/encoding/simplifiedchinese"
)

const (
	pendingQueryTimeout = 30 * time.Second
)

// CatalogQueryResult is sent back to the API handler via the pending query channel.
type CatalogQueryResult struct {
	DeviceID string
	Success  bool   // true when received >= SumNum
	Received int    // actual items accumulated
	SumNum   int    // expected total from NVR
	Items    []CatalogItem
	Err      error
}

type pendingCatalogQuery struct {
	mu        sync.Mutex
	DeviceID  string
	SN        string
	SumNum    int
	Items     []CatalogItem
	ResultCh  chan *CatalogQueryResult
	CreatedAt time.Time
}

type CatalogHandler struct {
	logger         *zap.Logger
	pendingQueries sync.Map
}

// GlobalCatalogHandler is a package-level reference set by NewCatalogHandler,
// used by API handlers to register pending queries and wait for results.
var GlobalCatalogHandler *CatalogHandler

func NewCatalogHandler(logger *zap.Logger) *CatalogHandler {
	h := &CatalogHandler{logger: logger}
	GlobalCatalogHandler = h
	go h.cleanupLoop()
	return h
}

func (h *CatalogHandler) cleanupLoop() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		h.pendingQueries.Range(func(key, value any) bool {
			pq := value.(*pendingCatalogQuery)
			pq.mu.Lock()
			stale := time.Since(pq.CreatedAt) > pendingQueryTimeout
			pq.mu.Unlock()
			if stale {
				h.logger.Warn("Pending catalog query timed out, cleaning up",
					zap.String("device_id", pq.DeviceID),
					zap.Int("received", len(pq.Items)),
					zap.Int("sum_num", pq.SumNum),
				)
				h.pendingQueries.Delete(key)
			}
			return true
		})
	}
}

// RegisterPendingQuery registers a pending catalog query for the API handler to wait on.
func (h *CatalogHandler) RegisterPendingQuery(deviceID, sn string) *pendingCatalogQuery {
	pq := &pendingCatalogQuery{
		DeviceID:  deviceID,
		SN:        sn,
		Items:     make([]CatalogItem, 0),
		ResultCh:  make(chan *CatalogQueryResult, 1),
		CreatedAt: time.Now(),
	}
	h.pendingQueries.Store(deviceID, pq)
	return pq
}

// RemovePendingQuery removes a pending query entry.
func (h *CatalogHandler) RemovePendingQuery(deviceID string) {
	h.pendingQueries.Delete(deviceID)
}

// GetPendingQuery returns the pending query for a device, if any.
func (h *CatalogHandler) GetPendingQuery(deviceID string) (*pendingCatalogQuery, bool) {
	v, ok := h.pendingQueries.Load(deviceID)
	if !ok {
		return nil, false
	}
	return v.(*pendingCatalogQuery), true
}

// Progress returns the current received count and SumNum safely (with lock).
func (pq *pendingCatalogQuery) Progress() (received, sumNum int) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return len(pq.Items), pq.SumNum
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
	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(charset) {
		case "gb2312", "gbk":
			return simplifiedchinese.GBK.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unsupported charset: %s", charset)
		}
	}
	if err := decoder.Decode(resp); err != nil {
		return nil, fmt.Errorf("parse catalog response failed: %w", err)
	}

	if resp.CmdType != "Catalog" {
		return nil, fmt.Errorf("invalid cmd type: %s", resp.CmdType)
	}

	return resp, nil
}

// HandleCatalogResponse accumulates items into a pending query.
// It does NOT write to the database — the API handler that initiated the query
// is responsible for deleting old channels and batch-inserting after all items arrive.
func (h *CatalogHandler) HandleCatalogResponse(resp *CatalogResponse) error {
	deviceID := resp.DeviceID
	channelCount := len(resp.DeviceList.Items)

	h.logger.Info("Catalog response received",
		zap.String("device_id", deviceID),
		zap.String("sn", resp.SN),
		zap.Int("sum_num", resp.SumNum),
		zap.Int("channel_count", channelCount),
	)

	if channelCount == 0 {
		return nil
	}

	// Look up the pending query for this device
	pq, ok := h.GetPendingQuery(deviceID)
	if !ok {
		h.logger.Warn("No pending catalog query for device, ignoring response",
			zap.String("device_id", deviceID),
		)
		return nil
	}

	pq.mu.Lock()
	defer pq.mu.Unlock()

	// Store SumNum from first response
	if pq.SumNum == 0 {
		pq.SumNum = resp.SumNum
	}

	// Accumulate items
	pq.Items = append(pq.Items, resp.DeviceList.Items...)
	h.logger.Info("Catalog accumulated",
		zap.String("device_id", deviceID),
		zap.Int("received", len(pq.Items)),
		zap.Int("total", pq.SumNum),
	)

	// If all items received, signal the waiter
	if pq.SumNum > 0 && len(pq.Items) >= pq.SumNum {
		result := &CatalogQueryResult{
			DeviceID: deviceID,
			Success:  true,
			Received: len(pq.Items),
			SumNum:   pq.SumNum,
			Items:    make([]CatalogItem, len(pq.Items)),
		}
		copy(result.Items, pq.Items)
		pq.ResultCh <- result
	}

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
