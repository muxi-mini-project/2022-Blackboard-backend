package user

import (
	"blackboard/handler"
	"blackboard/model"
	"database/sql"

	// "encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Wishforpeace/My-Tool/utils"
	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
)

func Check(c *gin.Context) {
	handler.SendResponse(c, "成功登录", nil)
}

// @Summary  用户信息
// @Tags user
// @Description 获取用户信息
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} model.User "{"msg":"获取成功"}"
// @Failure 400 {object} errno.Errno "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/Info [get]

func UserInfo(c *gin.Context) {
	ID := c.MustGet("student_id").(string)

	u, e := model.GetUserInfo(ID)
	var Info model.Info
	if e != nil {
		handler.SendBadRequest(c, "获取失败", nil)
	} else {
		Info = model.Info{
			Model:     u.Model,
			StudentID: u.StudentID,
			NickName:  u.NickName,
			Avatar:    u.Avatar,
		}
		handler.SendResponse(c, "获取成功", Info)

	}
}

// @Summary  修改用户名
// @Tags user
// @Description 接收新的User结构体来修改用户名
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param User body model.User true "需要修改的用户信息"
// @Success 200 {object} handler.Response "{"msg":"修改成功"}"
// @Failure 400 {object} errno.Errno "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/changename [put]
func ChangeUserName(c *gin.Context) {
	var Info model.Info
	Info.StudentID = c.MustGet("student_id").(string)
	if err := c.BindJSON(&Info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if e := model.ChangeName(Info); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "更改失败"})
		return
	}
	handler.SendResponse(c, "修改成功", nil)
	c.Redirect(http.StatusMovedPermanently, "/user/Info"+Info.StudentID)
}

// @Summary  查看用户收藏
// @Tags user
// @Description 查看用户收藏的通告
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Collection "{"msg":"获取成功"}"
// @Failure 400 {object} errno.Errno "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Router /user/colletion [get]
func CheckCollections(c *gin.Context) {
	id, ok := c.Get("student_id")
	if !ok {
		handler.SendBadRequest(c, "没有ID", ok)
	}
	ID := id.(string)
	collect, err := model.GetCollection(ID)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": "Fail.",
		})
	}
	handler.SendResponse(c, "获取成功", collect)
}

// @Summary  通告
// @Tags user
// @Description 查看用户发布过的通告
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Announcement "{"msg":"获取成功"}"
// @Failure 400 {object} errno.Errno "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Router /user/publisher [get]
func UserPublished(c *gin.Context) {
	id, ok := c.Get("student_id")
	if !ok {
		handler.SendBadRequest(c, "没有ID", ok)
	}
	ID := id.(string)
	published, err := model.GetPublished(ID)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": "Token Invalid",
		})
		return
	}
	handler.SendResponse(c, "获取成功", published)
}

// @Summary 修改头像
// @Tags user
// @Description 修改用户头像
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param file formData file true "文件"
// @Success 200 {object} model.User "{"mgs":"success"}"
// @Failure 200 {object} errno.Errno "绑定发生错误"
// @Failure 200 {object} errno.Errno "文件上传错误"
// @Failure 200 {object} errno.Errno "无法创建文件夹"
// @Failure 200 {object} errno.Errno "无法保存文件"
// @Failure 200 {object} errno.Errno "数据无法更新"
// @Failure 404 "该用户不存在"
// @Router /user/profile [post]
func UpdateUserProfile(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		log.Panicln("绑定发生错误 ", err.Error())
	}
	file, e := c.FormFile("avatar-file")
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("文件上传错误", e.Error())
	}
	path := utils.RootPath()
	path = filepath.Join(path, "avatar")
	fmt.Println("path =>", path)
	e = os.MkdirAll(path, os.ModePerm)
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	e = c.SaveUploadedFile(file, filepath.Join(path, fileName))
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}
	avatarUrl := "/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	e = user.UpdateUser(user.StudentID)
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("数据无法更新", e.Error())
	}
	c.Redirect(http.StatusMovedPermanently, "/user/profile?id="+user.StudentID)
}
