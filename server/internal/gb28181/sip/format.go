package sip

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// FormatIncoming formats an incoming SIP message as a readable multi-line string.
func FormatIncoming(msg *SIPMessage, addr string) string {
	var buf strings.Builder

	if msg.IsRequest && msg.RequestLine != nil {
		buf.WriteString(fmt.Sprintf("←── SIP %s from %s\n", msg.RequestLine.Method, addr))
		buf.WriteString(fmt.Sprintf("    %s %s %s\n", msg.RequestLine.Method, msg.RequestLine.RequestURI, msg.RequestLine.SIPVersion))
	} else if !msg.IsRequest && msg.StatusLine != nil {
		buf.WriteString(fmt.Sprintf("←── SIP %d %s from %s\n", msg.StatusLine.StatusCode, msg.StatusLine.ReasonPhrase, addr))
		buf.WriteString(fmt.Sprintf("    %s %d %s\n", msg.StatusLine.SIPVersion, msg.StatusLine.StatusCode, msg.StatusLine.ReasonPhrase))
	} else {
		buf.WriteString(fmt.Sprintf("←── SIP Message from %s\n", addr))
	}

	for _, h := range orderedHeaders(msg) {
		buf.WriteString(fmt.Sprintf("    %s: %s\n", h.key, h.value))
	}

	if len(msg.Body) > 0 {
		buf.WriteString(fmt.Sprintf("\n    Body (%d bytes):\n", len(msg.Body)))
		pretty := PrettyXML(msg.Body)
		for _, line := range strings.Split(pretty, "\n") {
			buf.WriteString(fmt.Sprintf("      %s\n", line))
		}
	}

	return strings.TrimRight(buf.String(), "\n")
}

// FormatOutgoing formats an outgoing SIP message as a readable multi-line string.
func FormatOutgoing(msg *SIPMessage, target string) string {
	var buf strings.Builder

	if msg.IsRequest && msg.RequestLine != nil {
		buf.WriteString(fmt.Sprintf("──→ SIP %s to %s\n", msg.RequestLine.Method, target))
		buf.WriteString(fmt.Sprintf("    %s %s %s\n", msg.RequestLine.Method, msg.RequestLine.RequestURI, msg.RequestLine.SIPVersion))
	} else if !msg.IsRequest && msg.StatusLine != nil {
		buf.WriteString(fmt.Sprintf("──→ SIP %d %s to %s\n", msg.StatusLine.StatusCode, msg.StatusLine.ReasonPhrase, target))
		buf.WriteString(fmt.Sprintf("    %s %d %s\n", msg.StatusLine.SIPVersion, msg.StatusLine.StatusCode, msg.StatusLine.ReasonPhrase))
	} else {
		buf.WriteString(fmt.Sprintf("──→ SIP Message to %s\n", target))
	}

	for _, h := range orderedHeaders(msg) {
		buf.WriteString(fmt.Sprintf("    %s: %s\n", h.key, h.value))
	}

	if len(msg.Body) > 0 {
		buf.WriteString(fmt.Sprintf("\n    Body (%d bytes):\n", len(msg.Body)))
		pretty := PrettyXML(msg.Body)
		for _, line := range strings.Split(pretty, "\n") {
			buf.WriteString(fmt.Sprintf("      %s\n", line))
		}
	}

	return strings.TrimRight(buf.String(), "\n")
}

type headerEntry struct {
	key   string
	value string
}

// orderedHeaders returns headers in a consistent display order.
func orderedHeaders(msg *SIPMessage) []headerEntry {
	order := []string{"Via", "From", "To", "Call-ID", "CSeq", "Contact", "User-Agent", "Content-Type", "Content-Length", "Max-Forwards"}
	seen := make(map[string]bool)
	var result []headerEntry

	for _, name := range order {
		if v := msg.GetHeader(name); v != "" {
			result = append(result, headerEntry{name, v})
			seen[name] = true
		}
	}

	// Any remaining headers not in the order list
	for name, values := range msg.Headers {
		if !seen[name] && len(values) > 0 {
			result = append(result, headerEntry{name, values[0]})
			seen[name] = true
		}
	}

	return result
}

// PrettyXML attempts to pretty-print XML bytes with proper indentation.
// Returns the original data as-is if XML parsing fails (not valid XML).
func PrettyXML(data []byte) string {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return ""
	}

	var buf bytes.Buffer
	d := xml.NewDecoder(bytes.NewReader(trimmed))
	// Accept any charset encoding declaration, pass raw bytes through
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return input, nil
	}
	e := xml.NewEncoder(&buf)
	e.Indent("", "  ")

	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return string(trimmed)
		}
		if err := e.EncodeToken(t); err != nil {
			return string(trimmed)
		}
	}
	if err := e.Flush(); err != nil {
		return string(trimmed)
	}

	return buf.String()
}
