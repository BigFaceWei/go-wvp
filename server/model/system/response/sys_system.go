package response

import "wvp-go/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
