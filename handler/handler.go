package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Data interface{} `json:"data"`
}

//Response 请求响应
type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"msg"`
	Data    interface{} `json:"data"`
} //@name response

//
func SendResponse(c *gin.Context, message interface{}, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, message interface{}, data interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    data,
	})
}
