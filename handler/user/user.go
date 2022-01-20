package user

import (
	"blackboard/handler"
	"blackboard/model"
	"blackboard/service/user"
	"encoding/base64"
	"fmt"
	"go/token"

	"github.com/gin-gonic/gin"
)

//@Summary 用户信息
//@Tags user
//@Description 获取用户信息
func ChangeUserName(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}

	user.StudentID = id
	if user.PassWord != "" {
		user.PassWord = base64.StdEncoding.EncodeToString([]byte(user.PassWord))
	}
	if err := model.ChangeName(user); err != nil {
		c.JSON(400, gin.H{"message": "更新失败"})
		return
	}
	handler.SendResponse(c, "修改成功", nil)
}
