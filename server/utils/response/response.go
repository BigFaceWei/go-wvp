package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

func Fail(c *gin.Context, code ErrorCode, data interface{}) {
	c.JSON(code.HTTPStatus(), Response{
		Code:    int(code),
		Message: code.Message(),
		Data:    data,
	})
}

func FailWithMessage(c *gin.Context, code ErrorCode, message string) {
	c.JSON(code.HTTPStatus(), Response{
		Code:    int(code),
		Message: message,
	})
}

func FailWithDetail(c *gin.Context, code ErrorCode, message string, data interface{}) {
	c.JSON(code.HTTPStatus(), Response{
		Code:    int(code),
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code ErrorCode) {
	c.JSON(code.HTTPStatus(), Response{
		Code:    int(code),
		Message: code.Message(),
	})
}