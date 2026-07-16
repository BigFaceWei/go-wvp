package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	"wvp-go/server/global"
	"wvp-go/server/internal/gb28181/message"
	"wvp-go/server/model/system"
)

// sipSession stores SIP transaction info needed to send BYE when stopping a stream.
type sipSession struct {
	CallID     string
	FromTag    string
	ToTag      string
	TargetAddr string
	CSeq       int
	ChannelID  string
}

// activeSessions maps streamID → session info for active play sessions.
var activeSessions sync.Map

type PlayResult struct {
	DeviceID  string            `json:"device_id"`
	ChannelID string            `json:"channel_id"`
	SSRC      string            `json:"ssrc"`
	MediaIP   string            `json:"media_ip"`
	MediaPort int               `json:"media_port"`
	PlayURL   string            `json:"play_url"`
	Stream    string            `json:"stream"`   // matches Java wvp streamInfo.stream
	App       string            `json:"app"`       // matches Java wvp streamInfo.app
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

	// stream_id format matches wvp-GB28181-pro: deviceId_channelId
	streamID := fmt.Sprintf("%s_%s", deviceID, channelID)

	// Check if RTP stream already exists (matching wvp-GB28181-pro's getRtpInfo)
	rtpExists, rtpPort, err := checkRtpInfo(zlmIP, zlmHTTPPort, zlmSecret, streamID)
	if err != nil {
		logger.Warn("Failed to check RTP info, will try to open", zap.Error(err))
	}

	if rtpExists {
		logger.Info("RTP stream already exists, returning existing stream URLs",
			zap.String("stream_id", streamID),
			zap.Int("media_port", rtpPort),
		)
	} else {
		rtpPort, err = openRtpServer(zlmIP, zlmHTTPPort, zlmSecret, streamID)
		if err != nil {
			return nil, fmt.Errorf("open RTP server failed: %w", err)
		}
		logger.Info("RTP server opened", zap.String("ssrc", ssrc), zap.Int("port", rtpPort), zap.String("stream_id", streamID))

		// Only send INVITE if we opened a new RTP server
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

		// Format matches wvp-GB28181-pro: {channelId}:{ssrc},{sipId}:0
		subject := fmt.Sprintf("%s:%s,%s:0", channelID, ssrc, global.GVA_CONFIG.WVP.SIP.ServerID)
		txn, err := srv.SendInvite(channelID, targetAddr, sdp, subject)
		if err != nil {
			closeRtpServer(zlmIP, zlmHTTPPort, zlmSecret, streamID)
			return nil, fmt.Errorf("send INVITE failed: %w", err)
		}

		playTimeout := 30 * time.Second
		if global.GVA_CONFIG.WVP.UserSettings.PlayTimeout > 0 {
			playTimeout = time.Duration(global.GVA_CONFIG.WVP.UserSettings.PlayTimeout) * time.Millisecond
		}

		resp, err := txn.WaitForResponse(playTimeout)
		if err != nil {
			closeRtpServer(zlmIP, zlmHTTPPort, zlmSecret, streamID)
			return nil, fmt.Errorf("wait INVITE response failed: %w", err)
		}
		if resp.StatusLine == nil {
			closeRtpServer(zlmIP, zlmHTTPPort, zlmSecret, streamID)
			return nil, fmt.Errorf("empty SIP response")
		}
		if resp.StatusLine.StatusCode != 200 {
			closeRtpServer(zlmIP, zlmHTTPPort, zlmSecret, streamID)
			return nil, fmt.Errorf("INVITE rejected: %d %s", resp.StatusLine.StatusCode, resp.StatusLine.ReasonPhrase)
		}

		logger.Info("INVITE 200 OK",
			zap.String("device_id", deviceID),
			zap.Int("status", resp.StatusLine.StatusCode),
			zap.Int("body_len", len(resp.Body)),
		)

		// ── Send ACK after receiving 200 OK for INVITE ──
		callID := txn.Request.GetHeader("Call-ID")
		fromHeader := txn.Request.GetHeader("From")
		toHeader := resp.GetHeader("To")
		cseqStr := txn.Request.GetHeader("CSeq")

		fromTag := extractTag(fromHeader)
		toTag := extractTag(toHeader)

		// Extract CSeq number for BYE later
		cseqNum := 0
		if idx := strings.Index(cseqStr, " "); idx > 0 {
			fmt.Sscanf(cseqStr[:idx], "%d", &cseqNum)
		}

		if callID != "" && fromTag != "" {
			logger.Info("Sending ACK for INVITE",
				zap.String("target", targetAddr),
				zap.String("call_id", callID),
				zap.String("from_tag", fromTag),
				zap.String("to_tag", toTag),
				zap.String("cseq", cseqStr),
			)
			if err := srv.SendACK(targetAddr, callID, fromTag, toTag, cseqStr); err != nil {
				logger.Warn("Send ACK failed (non-fatal)", zap.Error(err))
			}
		} else {
			logger.Warn("Cannot send ACK: missing Call-ID or From tag")
		}

		// Store session info for BYE on stop/auto-close
		activeSessions.Store(streamID, &sipSession{
			CallID:     callID,
			FromTag:    fromTag,
			ToTag:      toTag,
			TargetAddr: targetAddr,
			CSeq:       cseqNum + 1, // next CSeq for BYE
			ChannelID:  channelID,
		})
	}

	// Stream URLs match wvp-GB28181-pro: app=rtp, stream=deviceId_channelId
	// ZLMediaKit serves different protocols on different ports:
	//   HTTP-FLV/WS-FLV/HLS on HTTP port, RTSP on 554, RTMP on 1935
	rtpAppStream := fmt.Sprintf("rtp/%s", streamID)

	rtspURL := fmt.Sprintf("rtsp://%s:554/%s", zlmIP, rtpAppStream)
	rtmpURL := fmt.Sprintf("rtmp://%s:1935/%s", zlmIP, rtpAppStream)
	hlsURL := fmt.Sprintf("http://%s:%d/%s/hls.m3u8", zlmIP, zlmHTTPPort, rtpAppStream)
	httpFlvURL := fmt.Sprintf("http://%s:%d/%s.live.flv", zlmIP, zlmHTTPPort, rtpAppStream)
	wsFlvURL := fmt.Sprintf("ws://%s:%d/%s.live.flv", zlmIP, zlmHTTPPort, rtpAppStream)

	time.Sleep(2 * time.Second)

	mediaList, _ := checkMediaList(zlmIP, zlmHTTPPort, zlmSecret, streamID)
	logger.Info("Play started",
		zap.String("device_id", deviceID),
		zap.String("channel_id", channelID),
		zap.String("stream_id", streamID),
		zap.Int("zlm_stream_count", mediaList),
	)

	return &PlayResult{
		DeviceID:  deviceID,
		ChannelID: channelID,
		SSRC:      ssrc,
		MediaIP:   zlmIP,
		MediaPort: rtpPort,
		PlayURL:   rtspURL,
		Stream:    streamID, // matches Java wvp: deviceId_channelId
		App:       "rtp",    // matches Java wvp: rtp app from openRtpServer
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

func StopVideo(deviceID, channelID string) error {
	logger := global.GVA_LOG.Named("play")
	logger.Info("Stopping video",
		zap.String("device_id", deviceID),
		zap.String("channel_id", channelID),
	)

	// stream_id matches the format used in openRtpServer: deviceId_channelId
	streamID := fmt.Sprintf("%s_%s", deviceID, channelID)

	// Send BYE to device if we have an active session (matches wvp-GB28181-pro's stop flow)
	if session, ok := activeSessions.LoadAndDelete(streamID); ok {
		s := session.(*sipSession)
		srv := global.GVA_SIP_SERVER
		if srv != nil {
			logger.Info("Sending BYE to device",
				zap.String("call_id", s.CallID),
				zap.String("target", s.TargetAddr),
			)
			if err := sendBye(s, logger); err != nil {
				logger.Warn("Send BYE failed (non-fatal)", zap.Error(err))
			}
		}
	} else {
		logger.Info("No active SIP session found, skipping BYE",
			zap.String("stream_id", streamID),
		)
	}

	zlmIP, httpPort, zlmSecret, err := getZLMDefault()
	if err != nil {
		return fmt.Errorf("get ZLM config failed: %w", err)
	}

	// Close RTP server (matches wvp-GB28181-pro's closeRtpServer)
	apiPort := httpPort // ZLM serves API on HTTP port too
	if err := closeRtpServer(zlmIP, apiPort, zlmSecret, streamID); err != nil {
		logger.Warn("Close RTP server failed", zap.Error(err))
	}

	// Close all streams for this device/channel (matches wvp-GB28181-pro's close_streams)
	if err := closeStreams(zlmIP, apiPort, zlmSecret, "rtp", streamID); err != nil {
		logger.Warn("Close streams failed", zap.Error(err))
	}

	// Update channel: clear stream_id (matches wvp-GB28181-pro's stopPlay)
	if err := global.GVA_DB.Model(&system.DeviceChannel{}).
		Where("device_id = ? AND channel_id = ?", deviceID, channelID).
		Update("stream_id", "").Error; err != nil {
		logger.Warn("Update channel stream_id failed", zap.Error(err))
	}

	logger.Info("Video stopped successfully",
		zap.String("device_id", deviceID),
		zap.String("channel_id", channelID),
	)
	return nil
}

// sendBye sends a BYE request to the device to end the SIP session.
func sendBye(session *sipSession, logger *zap.Logger) error {
	srv := global.GVA_SIP_SERVER
	if srv == nil {
		return fmt.Errorf("SIP server not available")
	}

	contact := fmt.Sprintf("<sip:%s@%s:%d>",
		global.GVA_CONFIG.WVP.SIP.ServerID,
		global.GVA_CONFIG.WVP.SIP.ListenIP,
		global.GVA_CONFIG.WVP.SIP.ListenPort,
	)
	fromHeader := fmt.Sprintf("<sip:%s@%s>;tag=%s",
		global.GVA_CONFIG.WVP.SIP.ServerID,
		global.GVA_CONFIG.WVP.SIP.Domain,
		session.FromTag,
	)
	toHeader := fmt.Sprintf("<sip:%s@%s>;tag=%s",
		session.ChannelID,
		session.TargetAddr,
		session.ToTag,
	)

	requestURI := fmt.Sprintf("sip:%s@%s", session.ChannelID, session.TargetAddr)
	headers := map[string]string{
		"To":         toHeader,
		"From":       fromHeader,
		"Call-ID":    session.CallID,
		"Contact":    contact,
		"User-Agent": "WVP-GB28181/1.0",
		"CSeq":       fmt.Sprintf("%d BYE", session.CSeq),
	}

	_, err := srv.SendRequestTo("BYE", requestURI, session.TargetAddr, headers, nil)
	if err != nil {
		return fmt.Errorf("send BYE failed: %w", err)
	}
	logger.Info("BYE sent to device",
		zap.String("target", session.TargetAddr),
		zap.String("call_id", session.CallID),
	)
	return nil
}

func closeRtpServer(zlmIP string, apiPort int, secret, streamID string) error {
	url := fmt.Sprintf("http://%s:%d/index/api/closeRtpServer?secret=%s&stream_id=%s",
		zlmIP, apiPort, secret, streamID)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("request closeRtpServer failed: %w", err)
	}
	defer resp.Body.Close()
	return nil
}

func closeStreams(zlmIP string, apiPort int, secret, app, stream string) error {
	// Matches wvp-GB28181-pro: POST /index/api/close_streams
	// with params: app=rtp, vhost=__defaultVhost__, stream=xxx, force=1
	url := fmt.Sprintf("http://%s:%d/index/api/close_streams?secret=%s&app=%s&vhost=__defaultVhost__&stream=%s&force=1",
		zlmIP, apiPort, secret, app, stream)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("request close_streams failed: %w", err)
	}
	defer resp.Body.Close()
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

// checkRtpInfo checks whether an RTP stream already exists on ZLM.
// Returns (exists, port, error). Matching wvp-GB28181-pro's getRtpInfo.
func checkRtpInfo(zlmIP string, httpPort int, secret, streamID string) (bool, int, error) {
	url := fmt.Sprintf("http://%s:%d/index/api/getRtpInfo?secret=%s&stream_id=%s",
		zlmIP, httpPort, secret, streamID)

	resp, err := http.Get(url)
	if err != nil {
		return false, 0, fmt.Errorf("request getRtpInfo failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, 0, fmt.Errorf("read getRtpInfo response failed: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return false, 0, fmt.Errorf("parse getRtpInfo response failed: %w", err)
	}

	exist, _ := result["exist"].(bool)
	port := 0
	if p, ok := result["local_port"].(float64); ok {
		port = int(p)
	}
	return exist, port, nil
}

func openRtpServer(zlmIP string, httpPort int, secret, streamID string) (int, error) {
	// Parameters match wvp-GB28181-pro:
	//   secret, re_use_port=0, stream_id=deviceId_channelId, port=0, only_audio=0, tcp_mode=0
	url := fmt.Sprintf("http://%s:%d/index/api/openRtpServer?secret=%s&port=0&tcp_mode=0&re_use_port=0&only_audio=0&stream_id=%s",
		zlmIP, httpPort, secret, streamID)

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
