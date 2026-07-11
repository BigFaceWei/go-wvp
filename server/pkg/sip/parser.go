package sip

import (
	"fmt"
	"strings"
)

type SIPMessage struct {
	RequestLine *RequestLine
	StatusLine  *StatusLine
	Headers     map[string][]string
	Body        []byte
	IsRequest   bool
}

type RequestLine struct {
	Method    string
	RequestURI string
	SIPVersion string
}

type StatusLine struct {
	SIPVersion string
	StatusCode int
	ReasonPhrase string
}

func ParseMessage(data []byte) (*SIPMessage, error) {
	msg := &SIPMessage{
		Headers: make(map[string][]string),
	}

	content := string(data)
	lines := strings.Split(content, "\r\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("empty SIP message")
	}

	firstLine := lines[0]
	if strings.HasPrefix(firstLine, "SIP/") {
		msg.IsRequest = false
		statusLine, err := parseStatusLine(firstLine)
		if err != nil {
			return nil, err
		}
		msg.StatusLine = statusLine
	} else {
		msg.IsRequest = true
		requestLine, err := parseRequestLine(firstLine)
		if err != nil {
			return nil, err
		}
		msg.RequestLine = requestLine
	}

	bodyStart := false
	var bodyLines []string
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			bodyStart = true
			continue
		}
		if bodyStart {
			bodyLines = append(bodyLines, line)
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			msg.Headers[key] = append(msg.Headers[key], value)
		}
	}

	if len(bodyLines) > 0 {
		msg.Body = []byte(strings.Join(bodyLines, "\r\n"))
	}

	return msg, nil
}

func parseRequestLine(line string) (*RequestLine, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line: %s", line)
	}

	return &RequestLine{
		Method:     parts[0],
		RequestURI: parts[1],
		SIPVersion: parts[2],
	}, nil
}

func parseStatusLine(line string) (*StatusLine, error) {
	parts := strings.SplitN(line, " ", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid status line: %s", line)
	}

	var code int
	_, err := fmt.Sscanf(parts[1], "%d", &code)
	if err != nil {
		return nil, fmt.Errorf("invalid status code: %s", parts[1])
	}

	return &StatusLine{
		SIPVersion: parts[0],
		StatusCode: code,
		ReasonPhrase: parts[2],
	}, nil
}

func (m *SIPMessage) GetHeader(name string) string {
	if values, ok := m.Headers[name]; ok && len(values) > 0 {
		return values[0]
	}
	return ""
}

func (m *SIPMessage) SetHeader(name, value string) {
	m.Headers[name] = []string{value}
}

func (m *SIPMessage) AddHeader(name, value string) {
	m.Headers[name] = append(m.Headers[name], value)
}

func (m *SIPMessage) SetBody(body []byte) {
	m.Body = body
}

func (m *SIPMessage) String() string {
	var sb strings.Builder

	if m.IsRequest && m.RequestLine != nil {
		sb.WriteString(fmt.Sprintf("%s %s %s\r\n", m.RequestLine.Method, m.RequestLine.RequestURI, m.RequestLine.SIPVersion))
	} else if !m.IsRequest && m.StatusLine != nil {
		sb.WriteString(fmt.Sprintf("%s %d %s\r\n", m.StatusLine.SIPVersion, m.StatusLine.StatusCode, m.StatusLine.ReasonPhrase))
	}

	for key, values := range m.Headers {
		for _, value := range values {
			sb.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
		}
	}

	sb.WriteString("\r\n")

	if len(m.Body) > 0 {
		sb.Write(m.Body)
	}

	return sb.String()
}