package auto

import "github.com/gin-gonic/gin"

type autoPlug struct{}

func (a *autoPlug) Register(group *gin.RouterGroup) {
	group.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "msg": "success", "data": gin.H{"status": "running"}})
	})
}

func (a *autoPlug) RouterPath() string {
	return "/auto"
}

var Plug = new(autoPlug)
