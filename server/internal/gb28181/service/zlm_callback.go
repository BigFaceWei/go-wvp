package service

import "wvp-go/server/internal/media/zlm"

func init() {
	// Register the StopVideo callback for on_stream_none_reader hook.
	// This is called when ZLM detects no one is watching a stream.
	// The service package imports zlm (via stream_proxy/stream_push),
	// so there's no circular dependency.
	zlm.OnStreamNoneReaderCallback = StopVideo
}
