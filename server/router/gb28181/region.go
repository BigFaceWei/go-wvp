package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type RegionRouter struct{}

var regionRouter = new(RegionRouter)

func (r *RegionRouter) InitRegionRouter(Router *gin.RouterGroup) {
	rg := Router.Group("/region")
	{
		rg.GET("/tree", gb28181.GetRegionTree)
		rg.GET("/list", gb28181.GetRegionList)
		rg.POST("", gb28181.CreateRegion)
		rg.PUT("/:code", gb28181.UpdateRegion)
		rg.DELETE("/:code", gb28181.DeleteRegion)
	}
}
