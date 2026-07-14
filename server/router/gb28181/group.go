package gb28181

import (
	"wvp-go/server/api/v1/gb28181"

	"github.com/gin-gonic/gin"
)

type GroupRouter struct{}

var groupRouter = new(GroupRouter)

func (g *GroupRouter) InitGroupRouter(Router *gin.RouterGroup) {
	r := Router.Group("/group")
	{
		r.GET("/tree", gb28181.GetGroupTree)
		r.GET("/list", gb28181.GetGroupList)
		r.POST("", gb28181.CreateGroup)
		r.PUT("/:id", gb28181.UpdateGroup)
		r.DELETE("/:id", gb28181.DeleteGroup)
	}
}
