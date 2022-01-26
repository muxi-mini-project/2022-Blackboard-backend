package organization

import (
	"blackboard/handler"
	"blackboard/model"
	"blackboard/pkg/errno"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Wishforpeace/My-Tool/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 查看组织
// @Tags organization
// @Description 查看目前已存在的所有组织
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Organization "{"msg":"获取成功"}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /api/v1/organization/personal/created [get]
func CheckAll(c *gin.Context) {
	var (
		org []model.Organization
		err error
	)
	org, err = model.GetAllOrganizations("查看所有组织")
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, "获取成功", org)
}

// @Summary  查看创建
// @Tags organization
// @Description 查看用户创建的组织
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Organization "{"msg":"查询成功"}"
// @Failure 400 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/organization/personal/created [get]
func CheckCreated(c *gin.Context) {
	id, ok := c.Get("student_id")
	if !ok {
		handler.SendBadRequest(c, "未输入ID", nil)
	}
	ID := id.(string)
	created, err := model.GetCreated(ID)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "Fail."})
		return
	}
	handler.SendResponse(c, "获取成功", created)
}

// @Summary  查看关注
// @Tags organization
// @Description 查看用户关注的组织
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.FollowingOrganization "{"msg":"获取成功"}"
// @Failure 400 {object} errno.Errno "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /api/v1/organization/personal/following [get]
func CheckFollowing(c *gin.Context) {
	id, ok := c.Get("student_id")
	if !ok {
		handler.SendBadRequest(c, "未输入身份", "null")
	}
	ID := id.(string)
	following, err := model.GetFollowing(ID)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "Fail."})
		return
	}
	handler.SendResponse(c, "获取成功", following)

}

// @Summary  查看指定组织
// @Tags organization
// @Description 查看某个组织的具体信息
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Organization "{"msg":"查询成功"}"
// @Failure 400 {object} errno.Errno "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /api/v1/organization/details [get]
type Detail struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func CheckDetails(c *gin.Context) {
	var details Detail
	if e := c.ShouldBindJSON(&details); e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
	}

	org, er := model.GetDetails(details.ID, details.Name)
	if er != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": er,
		})
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "查无此组"})
		return
	}
	handler.SendResponse(c, "获取成功", org)
}

// @Summary  新建组织
// @Tags organization
// @Description 用户新建组织以便发布信息
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Organization "{"msg":"新建成功"}"
// @Failure 400 {object} errno.Errno "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Router /api/v1/organization/create [post]
func CreateOne(c *gin.Context) {
	var org model.Organization
	if err := c.BindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
		return
	}
	result := model.DB.Create(&org)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "Fail.")
	}
	handler.SendResponse(c, "创建成功", nil)
	file, e := c.FormFile("avatar_file")
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
		log.Panicln("无法创建文件", e.Error())
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
	org.Avatar = sql.NullString{String: avatarUrl}
	e = org.UpdateOrganization(org.ID)
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("数据无法更新", e.Error())
	}
	c.Redirect(http.StatusMovedPermanently, "/organization/details"+org.OrganizationName)
}

// @Summary  关注组织
// @Tags organization
// @Description 用户关注一个已经被创建的组织
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.FollowingOrganization "{"msg":"新建成功"}"
// @Failure 400 {object} errno.Errno "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Failure 400 {object} errno.Errno
// @Router /api/v1/organization/personal/follow [post]
func FollowOneOrganization(c *gin.Context) {
	var following model.FollowingOrganization
	if err := c.BindJSON(&following); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
	}
	result := model.DB.Create(&following)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "Fail.")
	}
	handler.SendResponse(c, "关注成功", nil)
}
