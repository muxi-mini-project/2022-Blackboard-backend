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
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
} //@name response

//
func SendResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}
