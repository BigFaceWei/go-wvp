package sip

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	psip "wvp-go/server/pkg/sip"
)

var (
	internalTagCounter    uint64
	internalCallIDCounter uint64
)

type SIPMessage = psip.SIPMessage
type RequestLine = psip.RequestLine
type StatusLine = psip.StatusLine

var ParseMessage = psip.ParseMessage
var NewBuilder = psip.NewBuilder
var ExtractBranch = psip.ExtractBranch

func generateTag() string {
	n := atomic.AddUint64(&internalTagCounter, 1)
	return fmt.Sprintf("tag%d-%d", time.Now().UnixNano(), n)
}

func generateBranch() string {
	n := atomic.AddUint64(&internalTagCounter, 1) // reuse counter
	return fmt.Sprintf("z9hG4bK%d-%d", time.Now().UnixNano(), n)
}

func generateCallID() string {
	n := atomic.AddUint64(&internalCallIDCounter, 1)
	return fmt.Sprintf("%d-%d@wvp-gb28181", time.Now().UnixNano(), n)
}

// ExtractBranchParam extracts the branch parameter value from a Via header.
// Input: "SIP/2.0/UDP 192.168.1.1:5060;branch=z9hG4bK12345"
// Output: "z9hG4bK12345"
func ExtractBranchParam(via string) string {
	prefix := "branch="
	if idx := strings.Index(via, prefix); idx != -1 {
		branch := via[idx+len(prefix):]
		// Trim any trailing whitespace or additional parameters
		if semi := strings.Index(branch, ";"); semi != -1 {
			branch = branch[:semi]
		}
		return strings.TrimSpace(branch)
	}
	return ""
}
