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

	srv.RegisterHandler("REGISTER", func(msg *gbsip.SIPMessage, addr string) {
		req, err := registerHandler.ParseRegister(msg)
		if err != nil {
			logger.Error("parse REGISTER failed", zap.Error(err))
			return
		}
		if err := registerHandler.HandleRegister(req); err != nil {
			logger.Error("handle REGISTER failed", zap.Error(err))
		}
	})

	srv.RegisterHandler("MESSAGE", func(msg *gbsip.SIPMessage, addr string) {
		body := msg.Body
		from := msg.GetHeader("From")
		
		if strings.Contains(from, global.GVA_CONFIG.WVP.SIP.ServerID) {
			logger.Debug("Skip MESSAGE from self")
			return
		}

		if strings.Contains(string(body), "<CmdType>Catalog</CmdType>") && strings.Contains(string(body), "<Response>") {
			resp, err := catalogHandler.ParseCatalogResponse(body)
			if err != nil {
				logger.Error("parse catalog response failed", zap.Error(err))
				return
			}
			if err := catalogHandler.HandleCatalogResponse(resp); err != nil {
				logger.Error("handle catalog response failed", zap.Error(err))
			}
			return
		}

		req, err := keepaliveHandler.ParseKeepalive(msg)
		if err != nil {
			logger.Error("parse MESSAGE failed", zap.Error(err))
			return
		}
		if err := keepaliveHandler.HandleKeepalive(req); err != nil {
			logger.Error("handle keepalive failed", zap.Error(err))
		}
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
