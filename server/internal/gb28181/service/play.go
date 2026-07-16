package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"

	"wvp-go/server/global"
	"wvp-go/server/internal/gb28181/message"
	"wvp-go/server/model/system"
)

type PlayResult struct {
	DeviceID  string            `json:"device_id"`
	ChannelID string            `json:"channel_id"`
	SSRC      string            `json:"ssrc"`
	MediaIP   string            `json:"media_ip"`
	MediaPort int               `json:"media_port"`
	PlayURL   string            `json:"play_url"`
	Flv       map[string]string `json:"flv"`
	Hls       string            `json:"hls"`
	Rtmp      string            `json:"rtmp"`
	Rtsp      string            `json:"rtsp"`
	WsFlv     map[string]string `json:"ws_flv"`
}

func PlayVideo(deviceID, channelID, ssrc string) (*PlayResult, error) {
	logger := global.GVA_LOG.Named("play")

	var device system.Device
	if err := global.GVA_DB.Where("device_id = ?", deviceID).First(&device).Error; err != nil {
		return nil, fmt.Errorf("device not found: %s", deviceID)
	}
	if !device.Online {
		return nil, fmt.Errorf("device offline: %s", deviceID)
	}

	if device.IP != "" {
		if host, portStr, err := net.SplitHostPort(device.IP); err == nil {
			device.IP = host
			if device.Port == 0 {
				fmt.Sscanf(portStr, "%d", &device.Port)
			}
		}
	}
	if device.IP == "" {
		return nil, fmt.Errorf("device address unknown: %s", deviceID)
	}

	srv := global.GVA_SIP_SERVER
	if srv == nil {
		return nil, fmt.Errorf("SIP server not initialized")
	}

	inviteHandler := message.NewInviteHandler(logger)
	if ssrc == "" {
		ssrc = inviteHandler.GenerateSSRC()
	}

	zlmIP, zlmHTTPPort, zlmSecret, err := getZLMDefault()
	if err != nil {
		return nil, err
	}

	rtpPort, err := openRtpServer(zlmIP, zlmHTTPPort, zlmSecret, ssrc)
	if err != nil {
		return nil, fmt.Errorf("open RTP server failed: %w", err)
	}
	logger.Info("RTP server opened", zap.String("ssrc", ssrc), zap.Int("port", rtpPort))

	sdp := inviteHandler.BuildInviteSDP(&message.InviteRequest{
		DeviceID:  deviceID,
		ChannelID: channelID,
		Domain:    global.GVA_CONFIG.WVP.SIP.Domain,
		Transport: device.Transport,
		MediaIP:   zlmIP,
		MediaPort: rtpPort,
		SSRC:      ssrc,
		Codecs:    []string{"PS"},
	})

	targetAddr := fmt.Sprintf("%s:%d", device.IP, device.Port)
	logger.Info("Sending INVITE",
		zap.String("device_id", deviceID),
		zap.String("channel_id", channelID),
		zap.String("target", targetAddr),
		zap.String("ssrc", ssrc),
	)

	subject := fmt.Sprintf("%s:%s,%s", channelID, ssrc, global.GVA_CONFIG.WVP.SIP.ServerID)
	txn, err := srv.SendInvite(channelID, targetAddr, sdp, subject)
	if err != nil {
		return nil, fmt.Errorf("send INVITE failed: %w", err)
	}

	timeout := 30 * time.Second
	if global.GVA_CONFIG.WVP.UserSettings.PlayTimeout > 0 {
		timeout = time.Duration(global.GVA_CONFIG.WVP.UserSettings.PlayTimeout) * time.Millisecond
	}

	resp, err := txn.WaitForResponse(timeout)
	if err != nil {
		return nil, fmt.Errorf("wait INVITE response failed: %w", err)
	}
	if resp.StatusLine == nil {
		return nil, fmt.Errorf("empty SIP response")
	}
	if resp.StatusLine.StatusCode != 200 {
		return nil, fmt.Errorf("INVITE rejected: %d %s", resp.StatusLine.StatusCode, resp.StatusLine.ReasonPhrase)
	}

	logger.Info("INVITE 200 OK",
		zap.String("device_id", deviceID),
		zap.Int("status", resp.StatusLine.StatusCode),
		zap.Int("body_len", len(resp.Body)),
	)

	// ── Send ACK after receiving 200 OK for INVITE ──
	// RFC 3261 §13.2.2.4: ACK for 2xx responses is sent by the UAC
	callID := txn.Request.GetHeader("Call-ID")
	fromHeader := txn.Request.GetHeader("From")
	toHeader := resp.GetHeader("To")
	cseq := txn.Request.GetHeader("CSeq")

	// Extract tags
	fromTag := extractTag(fromHeader)
	toTag := extractTag(toHeader)

	if callID != "" && fromTag != "" {
		logger.Info("Sending ACK for INVITE",
			zap.String("target", targetAddr),
			zap.String("call_id", callID),
			zap.String("from_tag", fromTag),
			zap.String("to_tag", toTag),
			zap.String("cseq", cseq),
		)
		if err := srv.SendACK(targetAddr, callID, fromTag, toTag, cseq); err != nil {
			logger.Warn("Send ACK failed (non-fatal)", zap.Error(err))
		}
	} else {
		logger.Warn("Cannot send ACK: missing Call-ID or From tag",
			zap.String("call_id", callID),
			zap.String("from_tag", fromTag),
		)
	}

	baseURL := fmt.Sprintf("%s:%d", zlmIP, zlmHTTPPort)
	stream := fmt.Sprintf("live/%s", ssrc)

	rtspURL := fmt.Sprintf("rtsp://%s/%s", baseURL, stream)
	rtmpURL := fmt.Sprintf("rtmp://%s/%s", baseURL, stream)
	hlsURL := fmt.Sprintf("http://%s/%s.m3u8", baseURL, stream)
	httpFlvURL := fmt.Sprintf("http://%s/%s.live.flv", baseURL, stream)
	wsFlvURL := fmt.Sprintf("ws://%s/%s.live.flv", baseURL, stream)

	time.Sleep(2 * time.Second)

	mediaList, _ := checkMediaList(zlmIP, zlmHTTPPort, zlmSecret, ssrc)
	logger.Info("Play started",
		zap.String("device_id", deviceID),
		zap.String("channel_id", channelID),
		zap.Int("zlm_stream_count", mediaList),
	)

	return &PlayResult{
		DeviceID:  deviceID,
		ChannelID: channelID,
		SSRC:      ssrc,
		MediaIP:   zlmIP,
		MediaPort: rtpPort,
		PlayURL:   rtspURL,
		Rtsp:      rtspURL,
		Rtmp:      rtmpURL,
		Hls:       hlsURL,
		Flv:       map[string]string{"http": httpFlvURL},
		WsFlv:     map[string]string{"ws": wsFlvURL},
	}, nil
}

// extractTag extracts the tag value from a SIP header like
// "<sip:device@domain>;tag=abc123" or "tag=abc123"
func extractTag(header string) string {
	if idx := strings.Index(header, "tag="); idx != -1 {
		tag := header[idx+4:]
		// If there's a trailing semicolon or bracket, trim it
		if semi := strings.Index(tag, ";"); semi != -1 {
			tag = tag[:semi]
		}
		if gt := strings.Index(tag, ">"); gt != -1 {
			tag = tag[:gt]
		}
		return strings.TrimSpace(tag)
	}
	return ""
}

func StopVideo(deviceID string) error {
	logger := global.GVA_LOG.Named("play")
	logger.Info("Stopping video", zap.String("device_id", deviceID))
	return nil
}

func getZLMDefault() (ip string, httpPort int, secret string, err error) {
	servers := global.GVA_CONFIG.WVP.ZLMediaKit.Servers
	if len(servers) == 0 {
		return "", 0, "", fmt.Errorf("ZLMediaKit not configured")
	}
	for _, s := range servers {
		if s.Default {
			return s.IP, s.HTTPPort, s.Secret, nil
		}
	}
	return servers[0].IP, servers[0].HTTPPort, servers[0].Secret, nil
}

func openRtpServer(zlmIP string, httpPort int, secret, ssrc string) (int, error) {
	url := fmt.Sprintf("http://%s:%d/index/api/openRtpServer?secret=%s&port=0&enable_tcp=1&stream_id=%s",
		zlmIP, httpPort, secret, ssrc)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("request openRtpServer failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read response failed: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("parse response failed: %w", err)
	}

	if code, ok := result["code"].(float64); ok && code != 0 {
		return 0, fmt.Errorf("openRtpServer error: %v", result["msg"])
	}

	port, ok := result["port"].(float64)
	if !ok {
		return 0, fmt.Errorf("openRtpServer: port not in response")
	}
	return int(port), nil
}

func checkMediaList(zlmIP string, httpPort int, secret, streamID string) (int, error) {
	url := fmt.Sprintf("http://%s:%d/index/api/getMediaList?secret=%s&schema=%s",
		zlmIP, httpPort, secret, streamID)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if data, ok := result["data"].([]interface{}); ok {
		return len(data), nil
	}
	return 0, nil
}
