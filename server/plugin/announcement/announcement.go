package announcement

import "github.com/gin-gonic/gin"

type announcementPlug struct{}

func (a *announcementPlug) Register(group *gin.RouterGroup) {
	group.GET("/list", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "msg": "success", "data": []interface{}{}})
	})
}

func (a *announcementPlug) RouterPath() string {
	return "/announcement"
}

var Plug = new(announcementPlug)
