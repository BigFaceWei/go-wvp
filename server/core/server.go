package core

import (
	"fmt"
	"strings"
	"time"

	"wvp-go/server/global"
	gbsip "wvp-go/server/internal/gb28181/sip"
	"wvp-go/server/initialize"
	"wvp-go/server/internal/gb28181/message"
	mcpTool "wvp-go/server/mcp"
	"wvp-go/server/service/system"
	"go.uber.org/zap"
)

func RunServer() {
	if global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
		if global.GVA_CONFIG.System.UseMultipoint {
			initialize.RedisList()
		}
	}

	if global.GVA_CONFIG.System.UseMongo {
		if err := initialize.Mongo.Initialization(); err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}

	if global.GVA_DB != nil {
		system.LoadAll()
	}

	go startSIPServer()

	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	mcpBaseURL := mcpTool.ResolveMCPServiceURL()

	fmt.Printf(`
	欢迎使用 gin-vue-admin
	当前版本:%s
	项目地址:https://github.com/flipped-aurora/gin-vue-admin
	插件市场:https://plugin.gin-vue-admin.com
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	MCP 独立服务请手动启动: go run ./cmd/mcp -config ./cmd/mcp/config.yaml
	默认MCP StreamHTTP地址:%s
	默认前端文件运行地址:http://127.0.0.1:8080
	GB28181 SIP 监听地址:%s:%d
`, global.Version, address, mcpBaseURL,
		global.GVA_CONFIG.WVP.SIP.ListenIP,
		global.GVA_CONFIG.WVP.SIP.ListenPort)

	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}

func startSIPServer() {
	sipCfg := global.GVA_CONFIG.WVP.SIP
	if sipCfg.ListenPort == 0 {
		return
	}

	logger := zap.L().Named("sip-server")

	srvConfig := &gbsip.ServerConfig{
		ListenIP:   sipCfg.ListenIP,
		ListenPort: sipCfg.ListenPort,
		Domain:     sipCfg.Domain,
		ServerID:   sipCfg.ServerID,
		Transport:  sipCfg.Transport,
		SIPLog:     sipCfg.SIPLog,
	}

	srv := gbsip.NewServer(srvConfig, logger)
	global.GVA_SIP_SERVER = srv

	registerHandler := message.NewRegisterHandler(logger)
	keepaliveHandler := message.NewKeepaliveHandler(logger)
	catalogHandler := message.NewCatalogHandler(logger)

	// Register handler — implements GB28181 Digest authentication
	srv.RegisterHandler("REGISTER", func(msg *gbsip.SIPMessage, addr string, txn *gbsip.Transaction) {
		req, err := registerHandler.ParseRegister(msg)
		if err != nil {
			logger.Error("parse REGISTER failed", zap.Error(err))
			return
		}

		// Handle unregistration (Expires: 0)
		if req.Expires <= 0 {
			registered, err := registerHandler.HandleRegister(req)
			if err != nil {
				logger.Error("handle unregister failed", zap.Error(err))
				srv.SendOK(txn) // Still respond OK so device knows we processed it
			} else if registered {
				logger.Warn("unexpected registered=true for expires=0")
			}
			srv.SendOK(txn)
			return
		}

		realm := sipCfg.Domain
		if realm == "" {
			realm = sipCfg.ServerID
		}

		// Check if this is an authenticated registration (with Authorization header)
		if !registerHandler.HasAuthorization(msg) {
			// Send 401 challenge — no Authorization header present
			nonce := registerHandler.GenerateNonce()
			logger.Info("Sending 401 digest challenge for device registration",
				zap.String("device_id", req.DeviceID),
				zap.String("nonce", nonce),
			)
			srv.SendUnauthorized(txn, realm, nonce)
			return
		}

		// Verify digest auth
		password, err := registerHandler.GetDevicePassword(req.DeviceID)
		if err != nil {
			logger.Warn("Device not found in database, using default password",
				zap.String("device_id", req.DeviceID),
			)
			// Use configured default password if available, otherwise fail
			password = global.GVA_CONFIG.WVP.SIP.Password
		}

		if password == "" {
			logger.Warn("No password configured for device, using empty password",
				zap.String("device_id", req.DeviceID),
			)
		}

		if !registerHandler.VerifyDigest(req, password) {
			logger.Warn("Digest authentication failed",
				zap.String("device_id", req.DeviceID),
				zap.String("username", req.Username),
			)
			// Send 401 with new challenge
			nonce := registerHandler.GenerateNonce()
			srv.SendUnauthorized(txn, realm, nonce)
			return
		}

		// Auth successful — complete registration
		registered, err := registerHandler.HandleRegister(req)
		if err != nil {
			logger.Error("handle REGISTER failed", zap.Error(err))
			// Still send OK so device knows we received the registration
			srv.SendOK(txn)
			return
		}
		if registered {
			logger.Info("Device registered successfully with digest auth",
				zap.String("device_id", req.DeviceID),
			)
		}
		srv.SendOK(txn)
	})

	// MESSAGE handler — handles Catalog responses, Keepalive, etc.
	srv.RegisterHandler("MESSAGE", func(msg *gbsip.SIPMessage, addr string, txn *gbsip.Transaction) {
		body := msg.Body
		from := msg.GetHeader("From")

		if strings.Contains(from, global.GVA_CONFIG.WVP.SIP.ServerID) {
			logger.Debug("Skip MESSAGE from self")
			srv.SendOK(txn)
			return
		}

		if strings.Contains(string(body), "<CmdType>Catalog</CmdType>") && strings.Contains(string(body), "<Response>") {
			resp, err := catalogHandler.ParseCatalogResponse(body)
			if err != nil {
				logger.Error("parse catalog response failed", zap.Error(err))
				srv.SendOK(txn)
				return
			}
			if err := catalogHandler.HandleCatalogResponse(resp); err != nil {
				logger.Error("handle catalog response failed", zap.Error(err))
			}
			srv.SendOK(txn)
			return
		}

		req, err := keepaliveHandler.ParseKeepalive(msg)
		if err != nil {
			logger.Error("parse MESSAGE failed", zap.Error(err))
			srv.SendOK(txn)
			return
		}
		if err := keepaliveHandler.HandleKeepalive(req); err != nil {
			logger.Error("handle keepalive failed", zap.Error(err))
		}
		srv.SendOK(txn)
	})

	// Handle INVITE responses (200 OK with SDP) — send ACK
	srv.RegisterHandler("INVITE", func(msg *gbsip.SIPMessage, addr string, txn *gbsip.Transaction) {
		// INVITE requests from devices are handled by this handler.
		// Currently the server initiates INVITE as a client (for PlayVideo).
		// Handle incoming INVITE from upper platform (cascade) with 200 OK response.
		logger.Info("Received INVITE request",
			zap.String("from", addr),
			zap.String("uri", msg.RequestLine.RequestURI),
		)

		// Build 200 OK with SDP for the incoming INVITE
		sdpBody := `v=0
o=` + sipCfg.ServerID + ` 0 0 IN IP4 ` + sipCfg.ListenIP + `
s=Play
c=IN IP4 ` + sipCfg.ListenIP + `
t=0 0
m=video 0 RTP/AVP 96
a=rtpmap:96 PS/90000
a=sendonly
`
		response := &gbsip.SIPMessage{
			IsRequest: false,
			StatusLine: &gbsip.StatusLine{
				SIPVersion:   "SIP/2.0",
				StatusCode:   200,
				ReasonPhrase: "OK",
			},
			Headers: make(map[string][]string),
		}

		// Copy all Via headers in order
		if vias, ok := msg.Headers["Via"]; ok {
			for _, v := range vias {
				response.AddHeader("Via", v)
			}
		}
		response.SetHeader("From", msg.GetHeader("From"))
		response.SetHeader("To", msg.GetHeader("To"))
		response.SetHeader("Call-ID", msg.GetHeader("Call-ID"))
		response.SetHeader("CSeq", msg.GetHeader("CSeq"))
		response.SetHeader("Contact", fmt.Sprintf("<sip:%s@%s:%d>", sipCfg.ServerID, sipCfg.ListenIP, sipCfg.ListenPort))
		response.SetHeader("Content-Type", "application/sdp")
		response.SetBody([]byte(sdpBody))

		txn.SendResponse(response)
	})

	if err := srv.Start(); err != nil {
		logger.Error("SIP server start failed", zap.Error(err))
		return
	}

	logger.Info("SIP server started",
		zap.String("listen", fmt.Sprintf("%s:%d", sipCfg.ListenIP, sipCfg.ListenPort)),
	)

	select {}
}
