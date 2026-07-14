package sip

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"
)

var (
	branchCounter uint64
	tagCounter    uint64
	callIDCounter uint64
)

type Builder struct {
	headers    map[string][]string
	body       []byte
	listenAddr string // Via header sent-by address, e.g. "192.168.1.100:5060"
}

func NewBuilder() *Builder {
	return &Builder{
		headers:    make(map[string][]string),
		listenAddr: "0.0.0.0:5060", // fallback default
	}
}

func (b *Builder) SetListenAddr(addr string) *Builder {
	b.listenAddr = addr
	return b
}

func (b *Builder) SetMethod(method, requestURI string) *RequestLine {
	return &RequestLine{
		Method:     method,
		RequestURI: requestURI,
		SIPVersion: "SIP/2.0",
	}
}

func (b *Builder) SetStatus(code int, reason string) *StatusLine {
	return &StatusLine{
		SIPVersion:   "SIP/2.0",
		StatusCode:   code,
		ReasonPhrase: reason,
	}
}

func (b *Builder) SetHeader(name, value string) *Builder {
	b.headers[name] = []string{value}
	return b
}

func (b *Builder) AddHeader(name, value string) *Builder {
	b.headers[name] = append(b.headers[name], value)
	return b
}

func (b *Builder) SetBody(body []byte) *Builder {
	b.body = body
	return b
}

func (b *Builder) BuildRequest(method, requestURI, from, to, callID, cseq string) *SIPMessage {
	msg := &SIPMessage{
		IsRequest: true,
		RequestLine: &RequestLine{
			Method:     method,
			RequestURI: requestURI,
			SIPVersion: "SIP/2.0",
		},
		Headers: make(map[string][]string),
		Body:    b.body,
	}

	for key, values := range b.headers {
		msg.Headers[key] = values
	}

	msg.SetHeader("Via", fmt.Sprintf("SIP/2.0/UDP %s;branch=%s", b.listenAddr, generateBranch()))
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Call-ID", callID)
	msg.SetHeader("CSeq", cseq)
	msg.SetHeader("Max-Forwards", "70")
	msg.SetHeader("User-Agent", "WVP-GB28181/1.0")

	if len(b.body) > 0 {
		msg.SetHeader("Content-Type", "application/sdp")
		msg.SetHeader("Content-Length", fmt.Sprintf("%d", len(b.body)))
	}

	return msg
}

func (b *Builder) BuildResponse(statusCode int, reason, branch, from, to, callID, cseq string) *SIPMessage {
	msg := &SIPMessage{
		IsRequest: false,
		StatusLine: &StatusLine{
			SIPVersion:   "SIP/2.0",
			StatusCode:   statusCode,
			ReasonPhrase: reason,
		},
		Headers: make(map[string][]string),
		Body:    b.body,
	}

	for key, values := range b.headers {
		msg.Headers[key] = values
	}

	msg.SetHeader("Via", fmt.Sprintf("SIP/2.0/UDP %s;branch=%s", b.listenAddr, branch))
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Call-ID", callID)
	msg.SetHeader("CSeq", cseq)

	if len(b.body) > 0 {
		msg.SetHeader("Content-Type", "application/sdp")
		msg.SetHeader("Content-Length", fmt.Sprintf("%d", len(b.body)))
	}

	return msg
}

func generateBranch() string {
	n := atomic.AddUint64(&branchCounter, 1)
	return fmt.Sprintf("z9hG4bK%d-%d", time.Now().UnixNano(), n)
}

func BuildRegisterRequest(domain, serverIP string, expires int) *SIPMessage {
	builder := NewBuilder()
	builder.SetListenAddr(fmt.Sprintf("%s:%d", serverIP, 5060))
	from := fmt.Sprintf("<sip:%s@%s>;tag=%s", "41010500002000000001", domain, generateTag())
	to := fmt.Sprintf("<sip:%s@%s>", "41010500002000000001", domain)
	callID := generateCallID()
	cseq := "1 REGISTER"

	msg := builder.BuildRequest("REGISTER", fmt.Sprintf("sip:%s", domain), from, to, callID, cseq)
	msg.SetHeader("Expires", fmt.Sprintf("%d", expires))
	msg.SetHeader("Contact", fmt.Sprintf("<sip:%s@%s:%d>", "41010500002000000001", serverIP, 5060))

	return msg
}

func BuildKeepaliveResponse(request *SIPMessage) *SIPMessage {
	branch := request.GetHeader("Via")
	branch = ExtractBranch(branch)
	from := request.GetHeader("From")
	to := request.GetHeader("To")
	callID := request.GetHeader("Call-ID")
	cseq := request.GetHeader("CSeq")

	builder := NewBuilder()
	return builder.BuildResponse(200, "OK", branch, from, to, callID, cseq)
}

func BuildCatalogResponse(request *SIPMessage, deviceID string) *SIPMessage {
	branch := request.GetHeader("Via")
	branch = ExtractBranch(branch)
	from := request.GetHeader("From")
	to := request.GetHeader("To")
	callID := request.GetHeader("Call-ID")
	cseq := request.GetHeader("CSeq")

	xmlBody := fmt.Sprintf(`<?xml version="1.0" encoding="GB2312"?>
<Response>
    <CmdType>Catalog</CmdType>
    <SN>1</SN>
    <DeviceID>%s</DeviceID>
    <SumNum>1</SumNum>
    <DeviceList Num="1">
        <Item>
            <DeviceID>%s</DeviceID>
            <Name>Camera01</Name>
            <Manufacturer>Hikvision</Manufacturer>
            <Model>DS-2CD2T25FD</Model>
            <Owner></Owner>
            <CivilCode>410105</CivilCode>
            <Address></Address>
            <Parental>1</Parental>
            <ParentID>%s</ParentID>
            <SafetyWay>0</SafetyWay>
            <RegisterWay>1</RegisterWay>
            <Secrecy>0</Secrecy>
            <Status>ON</Status>
        </Item>
    </DeviceList>
</Response>`, deviceID, deviceID, deviceID)

	builder := NewBuilder()
	msg := builder.BuildResponse(200, "OK", branch, from, to, callID, cseq)
	msg.SetBody([]byte(xmlBody))
	msg.SetHeader("Content-Type", "Application/MANSCDP+xml")

	return msg
}

func ExtractBranch(via string) string {
	if strings.Contains(via, "branch=") {
		parts := strings.Split(via, "branch=")
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1])
		}
	}
	return generateBranch()
}

func generateTag() string {
	n := atomic.AddUint64(&tagCounter, 1)
	return fmt.Sprintf("tag%d-%d", time.Now().UnixNano(), n)
}

func generateCallID() string {
	n := atomic.AddUint64(&callIDCounter, 1)
	return fmt.Sprintf("%d-%d@wvp-gb28181", time.Now().UnixNano(), n)
}