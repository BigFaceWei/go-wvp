package message

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"

	"wvp-go/server/global"
	"wvp-go/server/model/system"
	"wvp-go/server/pkg/sip"

	"go.uber.org/zap"
)

type RegisterHandler struct {
	logger *zap.Logger
}

func NewRegisterHandler(logger *zap.Logger) *RegisterHandler {
	return &RegisterHandler{logger: logger}
}

type RegisterRequest struct {
	DeviceID    string
	Domain      string
	Expires     int
	Transport   string
	RemoteAddr  string
	Username    string
	Password    string
	Realm       string
	Nonce       string
	Response    string
}

func (h *RegisterHandler) ParseRegister(msg *sip.SIPMessage) (*RegisterRequest, error) {
	if !msg.IsRequest || msg.RequestLine.Method != "REGISTER" {
		return nil, fmt.Errorf("not a REGISTER message")
	}

	req := &RegisterRequest{
		Expires:   3600,
		Transport: "UDP",
	}

	from := msg.GetHeader("From")
	if from != "" {
		parts := strings.Split(from, "@")
		if len(parts) == 2 {
			deviceID := strings.TrimPrefix(parts[0], "<sip:")
			req.DeviceID = deviceID
			domain := strings.TrimSuffix(parts[1], ">")
			if idx := strings.Index(domain, ";"); idx != -1 {
				domain = domain[:idx]
			}
			req.Domain = domain
		}
	}

	if expires := msg.GetHeader("Expires"); expires != "" {
		fmt.Sscanf(expires, "%d", &req.Expires)
	}

	if contact := msg.GetHeader("Contact"); contact != "" {
		if idx := strings.Index(contact, "@"); idx != -1 {
			rest := contact[idx+1:]
			if endIdx := strings.Index(rest, ">"); endIdx != -1 {
				req.RemoteAddr = rest[:endIdx]
			}
		}
	}

	authorization := msg.GetHeader("Authorization")
	if authorization != "" {
		h.parseAuthorization(authorization, req)
	}

	return req, nil
}

func (h *RegisterHandler) parseAuthorization(header string, req *RegisterRequest) {
	header = strings.TrimPrefix(header, "Digest ")
	pairs := strings.Split(header, ",")
	for _, pair := range pairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.Trim(strings.TrimSpace(parts[1]), "\"")

			switch key {
			case "username":
				req.Username = value
			case "realm":
				req.Realm = value
			case "nonce":
				req.Nonce = value
			case "response":
				req.Response = value
			}
		}
	}
}

func (h *RegisterHandler) VerifyDigest(req *RegisterRequest, password string) bool {
	if req.Username == "" || req.Nonce == "" || req.Response == "" {
		return false
	}

	ha1 := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", req.Username, req.Realm, password))))
	ha2 := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("REGISTER:%s", req.Domain))))
	expectedResponse := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", ha1, req.Nonce, ha2))))

	return req.Response == expectedResponse
}

func (h *RegisterHandler) HandleRegister(req *RegisterRequest) error {
	device := &system.Device{}
	result := global.GVA_DB.Where("device_id = ?", req.DeviceID).First(device)
	if result.Error != nil {
		device = &system.Device{
			DeviceID:    req.DeviceID,
			Name:        fmt.Sprintf("Device-%s", req.DeviceID),
			Transport:   req.Transport,
			IP:          req.RemoteAddr,
			Online:      true,
			RegisterTime: time.Now(),
			KeepaliveTime: time.Now(),
		}
		if err := global.GVA_DB.Create(device).Error; err != nil {
			return fmt.Errorf("create device failed: %w", err)
		}
	} else {
		device.IP = req.RemoteAddr
		device.Transport = req.Transport
		device.Online = true
		device.RegisterTime = time.Now()
		device.KeepaliveTime = time.Now()
		if err := global.GVA_DB.Save(device).Error; err != nil {
			return fmt.Errorf("update device failed: %w", err)
		}
	}

	h.logger.Info("Device registered",
		zap.String("device_id", req.DeviceID),
		zap.String("ip", req.RemoteAddr),
		zap.Int("expires", req.Expires),
	)

	return nil
}

func (h *RegisterHandler) GenerateNonce() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (h *RegisterHandler) GenerateChallenge(realm string) string {
	nonce := h.GenerateNonce()
	return fmt.Sprintf("Digest realm=\"%s\", nonce=\"%s\", algorithm=MD5, qop=\"auth\"", realm, nonce)
}