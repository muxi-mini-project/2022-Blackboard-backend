package middleware

import (
	"blackboard/handler"
	"blackboard/pkg/auth"
	"blackboard/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	ctx, err := auth.ParseRequest(c)
	if err != nil {
		handler.SendResponse(c, errno.ErrTokenInvalid, err.Error())
		//终止函数运行
		c.Abort()
		return
	}
	//跨越中间件取直
	c.Set("userID", ctx.ID)
	c.Set("expiresAt", ctx.ExpiresAt)
}
