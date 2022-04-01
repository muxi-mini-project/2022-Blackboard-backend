package organization

import (
	"blackboard/handler"
	"blackboard/model"
	"blackboard/pkg/errno"
	"blackboard/services"
	"blackboard/services/connector"
	"blackboard/services/organization"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Detail struct {
	Name string `json:"name" binding:"required"`
	ID   string `json:"id"`
}

// @Summary		查看组织
// @Tags organization
// @Description 查看目前已存在的所有组织
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param limit query int true "limit--偏移量指定开始返回记录之前要跳过的记录数"
// @Param page query int true "page--限定制定要检索的记录数"
// @Success 200 {object} []model.Organization "{"msg":"获取成功"}"
// @Failure 500 {object} errno.Errno "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /organization [get]
func CheckAll(c *gin.Context) {
	var limit, page int
	var err error
	// 翻页
	limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrQuery, nil)
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrQuery, nil)
		return
	}
	org, err := organization.GetAllOrganizations(page*limit, limit)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, "获取成功", org)
}

// @Summary		查看创建
// @Tags organization
// @Description 查看用户创建的组织
// @Accept application/json
// @Produce application/json
// @Param limit query  int true "limit--偏移量指定开始返回记录之前要跳过的记录数 "
// @Param page  query  int true "page--限制指定要检索的记录数 "
// @Param Authorization header string true "token"
// @Success 200 {object} []model.Organization "{"msg":"查询成功"}"
// @Failue 203 {object} errno.Errno	"{"msg":"Fail"}"
// @Failure 400 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /organization/personal/created [get]
func CheckCreated(c *gin.Context) {
	var limit, page int
	var err error
	ID := c.MustGet("student_id").(string)
	limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrQuery, err.Error())
		return
	}
	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrQuery, err.Error())
		return
	}

	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "Fail."})
		return
	}
	created, err := organization.GetCreated(ID, page*limit, limit)
	if err != nil {
		handler.SendError(c, "Fail", errno.ErrDatabase)
		return
	}
	handler.SendResponse(c, "获取成功", created)
}

// @Summary		查看关注
// @Tags organization
// @Description 查看用户关注的组织
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param limit query  int true "limit--偏移量指定开始返回记录之前要跳过的记录数 "
// @Param page  query  int true "page--限制指定要检索的记录数 "
// @Success 200 {object} []model.FollowingOrganization "{"msg":"获取成功"}"
// @Failue 203 {object} errno.Errno	"{"msg":"Fail"}"
// @Router /organization/personal/following [get]
func CheckFollowing(c *gin.Context) {
	var limit, page int
	var err error
	ID := c.MustGet("student_id").(string)
	limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrQuery, err.Error())
		return
	}
	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrQuery, err.Error())
		return
	}
	var following []*model.FollowingOrganization
	following, err = organization.GetFollowing(ID, limit*page, limit)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "Fail."})
		return
	}
	handler.SendResponse(c, "获取成功", following)

}

// @Summary	查看指定组织
// @Tags organization
// @Description 查看某个组织的具体信息
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.Organization "{"msg":"查询成功"}"
// @Failue 203 {object} errno.Errno	"{"msg":"查无此组"}"
// @Failure 400 {object} errno.Errno "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Router /organization/details [get]
func CheckDetails(c *gin.Context) {
	var details Detail
	if e := c.ShouldBindJSON(&details); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
	}

	org, er := model.GetDetails(details.ID, details.Name)
	if er != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "查无此组"})
		return
	}
	handler.SendResponse(c, "获取成功", org)
}

// @Summary	新建组织
// @Tags organization
// @Description 用户新建组织以便发布信息
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param object body model.Organization true "新建组织"
// @Success 200 {object} []model.Organization "{"msg":"新建成功"}"
// @Failure 400 {object} errno.Errno "{"msg":"Lack Param Or Param Not Satisfiable."}"
// @Failure 400 {object} errno.Errno "{"msg":"Fail"}"
// @Router /organization/create [post]
func CreateOne(c *gin.Context) {
	var org model.Organization
	org.FounderID = c.MustGet("student_id").(string)
	if err := c.BindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
		return
	}
	result := model.DB.Create(&org)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "Fail.")
		return
	}
	handler.SendResponse(c, "创建成功", org)
}

// @Summary	修改logo
// @Tags organization
// @Description 修改组织logo
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param file formData file true "文件"
// @Param organization_name path string true "组织名字"
// @Success 200 {object} model.User "{"mgs":"success"}"
// @Failure 400 {object} errno.Errno "{"msg":"Lack Param Or Param Not Satisfiable."}"
// @Failure 400 {object} errno.Errno "{"msg":"不具有资格"}"
// @Failure 400 {object} errno.Errno "{"msg":"上传失败"}"
// @Failure 400 {object} errno.Errno "上传失败,请检查token与其他配置参数是否正确"
// @Router /organization/:organization_name/image [post]
func UploadImage(c *gin.Context) {
	ID := c.MustGet("student_id").(string)
	name := c.Param("organization_name")
	PATH := "Organizations"
	if name == "" {
		handler.SendBadRequest(c, "Lack Param Or Param Not Satisfiable.", nil)
		return
	}
	orgID := model.GetOrgID(name)
	verify := model.JudgeFounder(ID, orgID)
	if !verify {
		handler.SendBadRequest(c, "不具有资格", nil)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		handler.SendBadRequest(c, "上传失败", nil)
		return
	}
	filepath := "./"
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filepath, os.ModePerm)
		}
	}

	fileExt := path.Ext(filepath + file.Filename)

	file.Filename = name + fileExt

	filename := filepath + file.Filename

	if err := c.SaveUploadedFile(file, filename); err != nil {
		handler.SendBadRequest(c, "上传失败", nil)
		return
	}

	// 删除组织原头像

	organization, _ := model.GetDetails("", name)
	if organization.Path != "" && organization.Sha != "" {
		connector.RepoCreate().Del(organization.Path, organization.Sha)
	}

	// 上传新头像
	Base64 := services.ImagesToBase64(filename)
	organization.Avatar, organization.Path, organization.Sha = connector.RepoCreate().Push(PATH, file.Filename, Base64)

	os.Remove(filename)

	e := organization.UpdateOrganization()
	if organization.Avatar == "" || e != nil {
		handler.SendBadRequest(c, "上传失败,请检查token与其他配置参数是否正确", e)
		return
	}

	handler.SendResponse(c, "上传成功", map[string]interface{}{
		"url":  organization.Avatar,
		"sha":  organization.Sha,
		"path": organization.Path,
		"Id":   orgID,
	})
}

// @Summary		关注组织
// @Tags organization
// @Description 用户关注一个已经被创建的组织
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.FollowingOrganization "{"msg":"新建成功"}"
// @Failure 400 {object} errno.Errno "{"message": "Lack Param Or Param Not Satisfiable."}"
// @Failure 400 {object} errno.Errno "{"message":"关注成功"}"
// @Router /organization/follow [post]
func FollowOneOrganization(c *gin.Context) {
	var following model.FollowingOrganization
	following.StudentID = c.MustGet("student_id").(string)
	if err := c.BindJSON(&following); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
	}
	following.OrganizationID = model.GetOrgID(following.OrganizationName)
	e := model.Follow(following)
	if e != nil {
		handler.SendBadRequest(c, nil, e)
	} else {
		handler.SendResponse(c, "关注成功", e)

	}
}
