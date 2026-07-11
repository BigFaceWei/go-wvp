package sip

import (
	"fmt"

	psip "wvp-go/server/pkg/sip"
)

type SIPMessage = psip.SIPMessage
type RequestLine = psip.RequestLine
type StatusLine = psip.StatusLine

var ParseMessage = psip.ParseMessage
var NewBuilder = psip.NewBuilder

func generateTag() string {
	return fmt.Sprintf("tag%d", 1000000000+999999999)
}

func generateCallID() string {
	return fmt.Sprintf("%d@wvp-gb28181", 1000000000+999999999)
}
