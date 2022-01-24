package user

import (
	"blackboard/handler"
	"blackboard/model"
	"blackboard/service/user"
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
)

//@Summary  用户信息
// @Tags user
// @Description 获取用户信息
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} model.User "{"msg":"获取成功"}"
// @Failure 203 {object} error.Error "{"error_code":"20001", "message":"Fail."}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user [get]
func UserInfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{"message": "Token Invalid"})
		return
	}
	u, err := model.GetUserInfo(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(203, gin.H{"message": "Fail."})
		return
	}
	handler.SendResponse(c, "获取成功", u)

}

// @Summary  修改用户名
// @Tags user
// @Description 接收新的User结构体来修改用户名
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param User body model.User true "需要修改的用户信息"
// @Success 200 {object} handler.Response "{"msg":"修改成功"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/changename [put]
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
		c.JSON(400, gin.H{"message": "更改失败"})
		return
	}
	handler.SendResponse(c, "修改成功", nil)
}

//@Summary  查看用户收藏
//@Tags user
//@Description 查看用户收藏的通告
//@Accept application/json
//@Produce application/json
//@Param token header string true "token"
//@Success 200 {object} []model.Collection "{"msg":"获取成功"}"
//@Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
//@Failure 401 {object} error.Error "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
//@Failure 400 {object} error.Error "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
//@Router /user/colletion
func CheckCollections(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}
	collect, err := model.GetCollection(id)
	if err != nil {
		c.JSON(203, gin.H{
			"message": "Fail.",
		})
	}
	handler.SendResponse(c, "获取成功", collect)
}

//@Summary  通告
//@Tags user
//@Description 查看用户发布过的通告
//@Accept application/json
//@Produce application/json
//@Param token header string true "token"
//@Success 200 {object} []model.Announcement "{"msg":"获取成功"}"
//@Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
//@Failure 401 {object} error.Error "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
//@Failure 400 {object} error.Error "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
//@Router /user/publisher
func UserPublished(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{
			"message": "Token Invalid",
		})
		return
	}
	published, err := model.GetPublished(id)
	if err != nil {
		c.JSON(203, gin.H{
			"message": "Token Invalid",
		})
		return
	}
	handler.SendResponse(c, "获取成功", published)
}
