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

var streamPushService = service.NewStreamPushService(nil)

func GetStreamPushList(c *gin.Context) {
	name := c.Query("name")
	stream := c.Query("stream")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	pushes, total, err := streamPushService.GetList(name, stream, page, pageSize)
	if err != nil {
		global.GVA_LOG.Error("get stream push list failed", zap.Error(err))
		response.Fail(c, response.ERROR, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"list":  pushes,
			"total": total,
		},
	})
}

func GetStreamPush(c *gin.Context) {
	id := c.Param("id")
	push, err := streamPushService.Get(id)
	if err != nil {
		response.Fail(c, response.ERROR, fmt.Sprintf("push not found: %v", err))
		return
	}
	response.Success(c, push)
}

func CreateStreamPush(c *gin.Context) {
	var req system.StreamPush
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := streamPushService.Create(&req); err != nil {
		global.GVA_LOG.Error("create stream push failed", zap.Error(err))
		response.Fail(c, response.ERROR, "create failed")
		return
	}
	response.Success(c, req)
}

func UpdateStreamPush(c *gin.Context) {
	id := c.Param("id")
	push, err := streamPushService.Get(id)
	if err != nil {
		response.Fail(c, response.ERROR, "push not found")
		return
	}
	if err := c.ShouldBindJSON(push); err != nil {
		response.Fail(c, response.INVALID_PARAMS, nil)
		return
	}
	if err := streamPushService.Update(push); err != nil {
		global.GVA_LOG.Error("update stream push failed", zap.Error(err))
		response.Fail(c, response.ERROR, "update failed")
		return
	}
	response.Success(c, nil)
}

func DeleteStreamPush(c *gin.Context) {
	id := c.Param("id")
	if err := streamPushService.Delete(id); err != nil {
		global.GVA_LOG.Error("delete stream push failed", zap.Error(err))
		response.Fail(c, response.ERROR, "delete failed")
		return
	}
	response.Success(c, nil)
}

func StartStreamPush(c *gin.Context) {
	id := c.Param("id")
	if err := streamPushService.Start(id); err != nil {
		global.GVA_LOG.Error("start stream push failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

func StopStreamPush(c *gin.Context) {
	id := c.Param("id")
	if err := streamPushService.Stop(id); err != nil {
		global.GVA_LOG.Error("stop stream push failed", zap.Error(err))
		response.Fail(c, response.ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}
