package message

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
)

var (
	inviteBranchCounter uint64
	inviteSSRCCounter   uint64
)

type InviteHandler struct {
	logger *zap.Logger
}

func NewInviteHandler(logger *zap.Logger) *InviteHandler {
	return &InviteHandler{logger: logger}
}

type InviteRequest struct {
	DeviceID    string
	ChannelID   string
	Domain      string
	Transport   string
	MediaIP     string
	MediaPort   int
	SSRC        string
	Codecs      []string
}

type InviteResponse struct {
	DeviceID    string
	ChannelID   string
	SSRC        string
	MediaIP     string
	MediaPort   int
	Protocol    string
	Codecs      []string
}

func (h *InviteHandler) ParseInviteResponse(body []byte) (*InviteResponse, error) {
	sdp := string(body)
	resp := &InviteResponse{}

	lines := strings.Split(sdp, "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "v=") {
			continue
		} else if strings.HasPrefix(line, "o=") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				resp.SSRC = parts[1]
			}
		} else if strings.HasPrefix(line, "s=") {
			continue
		} else if strings.HasPrefix(line, "c=") {
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				resp.MediaIP = parts[2]
			}
		} else if strings.HasPrefix(line, "m=") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				portParts := strings.Split(parts[1], "/")
				if len(portParts) > 0 {
					fmt.Sscanf(portParts[0], "%d", &resp.MediaPort)
				}
				if len(parts) >= 4 {
					resp.Protocol = parts[3]
				}
			}
		} else if strings.HasPrefix(line, "y=") {
			resp.SSRC = strings.TrimPrefix(line, "y=")
		}
	}

	return resp, nil
}

func (h *InviteHandler) BuildInviteSDP(req *InviteRequest) string {
	sdp := fmt.Sprintf("v=0\r\n")
	sdp += fmt.Sprintf("o=%s %s 0 IN IP4 %s\r\n", req.ChannelID, req.SSRC, req.MediaIP)
	sdp += fmt.Sprintf("s=%s\r\n", "Play")
	sdp += fmt.Sprintf("c=IN IP4 %s\r\n", req.MediaIP)
	sdp += fmt.Sprintf("t=0 0\r\n")

	if req.SSRC != "" {
		sdp += fmt.Sprintf("y=%s\r\n", req.SSRC)
	}

	// GB/T 28181-2016 f= line: media format description
	// V2/4/2 = Video, H.264, Main Profile (standard for embedded DVR/NVR)
	sdp += fmt.Sprintf("f=V2/4/2\r\n")

	for _, codec := range req.Codecs {
		switch codec {
		case "PS":
			sdp += fmt.Sprintf("m=video %d RTP/AVP 96\r\n", req.MediaPort)
			sdp += fmt.Sprintf("a=rtpmap:96 PS/90000\r\n")
		case "H264":
			sdp += fmt.Sprintf("m=video %d RTP/AVP 96\r\n", req.MediaPort)
			sdp += fmt.Sprintf("a=rtpmap:96 H264/90000\r\n")
		case "AAC":
			sdp += fmt.Sprintf("m=audio %d RTP/AVP 97\r\n", req.MediaPort+2)
			sdp += fmt.Sprintf("a=rtpmap:97 MPEG4-GENERIC/8000/1\r\n")
		}
	}

	if len(req.Codecs) == 0 {
		sdp += fmt.Sprintf("m=video %d RTP/AVP 96\r\n", req.MediaPort)
		sdp += fmt.Sprintf("a=rtpmap:96 PS/90000\r\n")
	}

	sdp += fmt.Sprintf("a=recvonly\r\n")

	return sdp
}

func (h *InviteHandler) GenerateSSRC() string {
	n := atomic.AddUint64(&inviteSSRCCounter, 1)
	ts := uint64(time.Now().UnixNano())
	return fmt.Sprintf("%010d", (ts%10000000000)+n)
}

type BYEHandler struct {
	logger *zap.Logger
}

func NewBYEHandler(logger *zap.Logger) *BYEHandler {
	return &BYEHandler{logger: logger}
}

type BYERequest struct {
	CallID  string
	FromTag string
	ToTag   string
	CSeq    int
}

func (h *BYEHandler) BuildBYEMessage(callID, fromTag, toTag string, cseq int) string {
	n := atomic.AddUint64(&inviteBranchCounter, 1)
	msg := fmt.Sprintf("BYE SIP/2.0\r\n")
	msg += fmt.Sprintf("Via: SIP/2.0/UDP 0.0.0.0:5060;branch=z9hG4bK%d-%d\r\n", time.Now().UnixNano(), n)
	msg += fmt.Sprintf("From: <sip:41010500002000000001@4101050000>;tag=%s\r\n", fromTag)
	msg += fmt.Sprintf("To: <sip:%s@4101050000>;tag=%s\r\n", "device", toTag)
	msg += fmt.Sprintf("Call-ID: %s\r\n", callID)
	msg += fmt.Sprintf("CSeq: %d BYE\r\n", cseq)
	msg += fmt.Sprintf("Max-Forwards: 70\r\n")
	msg += fmt.Sprintf("User-Agent: WVP-GB28181/1.0\r\n")
	msg += fmt.Sprintf("\r\n")

	return msg
}

func (h *BYEHandler) ParseBYEMessage(msg []byte) (*BYERequest, error) {
	content := string(msg)
	lines := strings.Split(content, "\r\n")

	req := &BYERequest{}

	for _, line := range lines {
		if strings.HasPrefix(line, "Call-ID:") {
			req.CallID = strings.TrimSpace(strings.TrimPrefix(line, "Call-ID:"))
		} else if strings.HasPrefix(line, "From:") {
			if idx := strings.Index(line, "tag="); idx != -1 {
				req.FromTag = line[idx+4:]
			}
		} else if strings.HasPrefix(line, "To:") {
			if idx := strings.Index(line, "tag="); idx != -1 {
				req.ToTag = line[idx+4:]
			}
		} else if strings.HasPrefix(line, "CSeq:") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				fmt.Sscanf(parts[1], "%d", &req.CSeq)
			}
		}
	}

	return req, nil
}
