package gb28181

import (
	"fmt"
	"net/http"
	"strconv"

	"wvp-go/server/global"
	"wvp-go/server/internal/gb28181/service"
	"wvp-go/server/model/system"
	"wvp-go/server/utils/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var streamProxyService = service.NewStreamProxyService(nil)

func GetStreamProxyList(c *gin.Context) {
	name := c.Query("name")
	stream := c.Query("stream")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	proxies, total, err := streamProxyService.GetList(name, stream, page, pageSize)
	if err != nil {
		global.GVA_LOG.Error("get stream proxy list failed", zap.Error(err))
		response.Fail(c, response.ERROR, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"list":  proxies,
			"total": total,
		},
	})
}

func GetStreamProxy(c *gin.Context) {
	id := c.Param("id")
	proxy, err := streamProxyService.Get(id)
	if err != nil {
		response.Fail(c, response.ERROR, fmt.Sprintf("proxy not found: %v", err))
		return
	}
	response.Success(c, proxy)
}

func CreateStreamProxy(c *gin.Context) {
	var req system.StreamProxy
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := streamProxyService.Create(&req); err != nil {
		global.GVA_LOG.Error("create stream proxy failed", zap.Error(err))
		response.Fail(c, response.ERROR, "create failed")
		return
	}
	response.Success(c, req)
}

func UpdateStreamProxy(c *gin.Context) {
	id := c.Param("id")
	proxy, err := streamProxyService.Get(id)
	if err != nil {
		response.Fail(c, response.ERROR, "proxy not found")
		return
	}
	if err := c.ShouldBindJSON(proxy); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := streamProxyService.Update(proxy); err != nil {
		global.GVA_LOG.Error("update stream proxy failed", zap.Error(err))
		response.Fail(c, response.ERROR, "update failed")
		return
	}
	response.Success(c, nil)
}

func DeleteStreamProxy(c *gin.Context) {
	id := c.Param("id")
	if err := streamProxyService.Delete(id); err != nil {
		global.GVA_LOG.Error("delete stream proxy failed", zap.Error(err))
		response.Fail(c, response.ERROR, "delete failed")
		return
	}
	response.Success(c, nil)
}

func StartStreamProxy(c *gin.Context) {
	id := c.Param("id")
	if err := streamProxyService.Start(id); err != nil {
		global.GVA_LOG.Error("start stream proxy failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

func StopStreamProxy(c *gin.Context) {
	id := c.Param("id")
	if err := streamProxyService.Stop(id); err != nil {
		global.GVA_LOG.Error("stop stream proxy failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}
