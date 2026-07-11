package initialize

import (
	"wvp-go/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]

	gb28181Router := router.RouterGroupApp.GB28181
	gb28181Router.InitGB28181Router(privateGroup)

	holder(publicGroup, privateGroup)
}
