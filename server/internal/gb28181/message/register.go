package message

import (
	"crypto/md5"
	"fmt"
	"net"
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
	URI         string // request URI from Authorization header (used for HA2 computation)
	Cnonce      string // client nonce, required when qop="auth"
	NC          string // nonce count, required when qop="auth"
	QOP         string // quality of protection ("auth" or "auth-int")
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
			case "uri":
				req.URI = value
			case "cnonce":
				req.Cnonce = value
			case "nc":
				req.NC = value
			case "qop":
				req.QOP = value
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
	method := "REGISTER"
	// Use the URI from the Authorization header as the device used to compute HA2.
	// Per RFC 3261, this is the Request-URI of the REGISTER request.
	// Fall back to req.Domain-based URI for backward compatibility.
	uri := req.URI
	if uri == "" {
		uri = req.Domain
		if !strings.HasPrefix(uri, "sip:") {
			uri = "sip:" + uri
		}
	}
	ha2 := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s", method, uri))))
	// Per RFC 7616 Digest Authentication:
	//   No qop:           response = MD5(HA1:nonce:HA2)
	//   qop="auth"/'auth-int': response = MD5(HA1:nonce:nc:cnonce:qop:HA2)
	var expectedResponse string
	if req.QOP == "auth" || req.QOP == "auth-int" {
		expectedResponse = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, req.Nonce, req.NC, req.Cnonce, req.QOP, ha2))))
	} else {
		expectedResponse = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", ha1, req.Nonce, ha2))))
	}

	return req.Response == expectedResponse
}

// HasAuthorization returns true if the request contains an Authorization header.
func (h *RegisterHandler) HasAuthorization(msg *sip.SIPMessage) bool {
	return msg.GetHeader("Authorization") != ""
}

// HandleRegister processes a validated register request.
// Returns (registered, isNew, error).
// isNew is true when a new device record was created (first-time registration).
func (h *RegisterHandler) HandleRegister(req *RegisterRequest) (bool, bool, error) {
	// Handle unregistration (Expires: 0)
	if req.Expires <= 0 {
		device := &system.Device{}
		result := global.GVA_DB.Where("device_id = ?", req.DeviceID).First(device)
		if result.Error == nil {
			device.Online = false
			global.GVA_DB.Save(device)
			h.logger.Info("Device unregistered",
				zap.String("device_id", req.DeviceID),
			)
		}
		return false, false, nil
	}

	ip, port := parseHostPort(req.RemoteAddr)

	device := &system.Device{}
	result := global.GVA_DB.Where("device_id = ?", req.DeviceID).First(device)
	isNew := result.Error != nil
	if isNew {
		device = &system.Device{
			DeviceID:      req.DeviceID,
			Name:          fmt.Sprintf("Device-%s", req.DeviceID),
			Transport:     req.Transport,
			IP:            ip,
			Port:          port,
			Online:        true,
			RegisterTime:  time.Now(),
			KeepaliveTime: time.Now(),
		}
		if err := global.GVA_DB.Create(device).Error; err != nil {
			return false, false, fmt.Errorf("create device failed: %w", err)
		}
	} else {
		device.IP = ip
		device.Port = port
		device.Transport = req.Transport
		device.Online = true
		device.RegisterTime = time.Now()
		device.KeepaliveTime = time.Now()
		if err := global.GVA_DB.Save(device).Error; err != nil {
			return false, false, fmt.Errorf("update device failed: %w", err)
		}
	}

	h.logger.Info("Device registered",
		zap.String("device_id", req.DeviceID),
		zap.String("ip", req.RemoteAddr),
		zap.Int("expires", req.Expires),
	)

	return true, isNew, nil
}

// GetDevicePassword retrieves the device password from the database.
// Returns empty string if device not found or no password configured.
func (h *RegisterHandler) GetDevicePassword(deviceID string) (string, error) {
	device := &system.Device{}
	result := global.GVA_DB.Where("device_id = ?", deviceID).First(device)
	if result.Error != nil {
		return "", result.Error
	}
	return device.Password, nil
}

func (h *RegisterHandler) GenerateNonce() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (h *RegisterHandler) GenerateChallenge(realm string) string {
	nonce := h.GenerateNonce()
	return fmt.Sprintf("Digest realm=\"%s\", nonce=\"%s\", algorithm=MD5, qop=\"auth\"", realm, nonce)
}

func parseHostPort(addr string) (string, int) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return addr, 5060
	}
	var port int
	fmt.Sscanf(portStr, "%d", &port)
	if port == 0 {
		port = 5060
	}
	return host, port
}
