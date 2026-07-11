package router

import (
	"wvp-go/server/router/gb28181"
	"wvp-go/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System system.RouterGroup
	GB28181 gb28181.Router
}
