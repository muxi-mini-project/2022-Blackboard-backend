package organization

import (
	"blackboard/handler"
	"blackboard/model"
	"blackboard/service/user"

	"github.com/gin-gonic/gin"
)

//@Summary 查看组织
//@Tags organization
//@Description 查看目前已存在的所有组织
//@Accept application/json
//@Produce application/json
//@Param token header string true "token"
// @Success 200 {object} []model.Organization "{"msg":"查询成功"}"
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 401 {object} error.Error "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /organization/personal/created
func CheckAll(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid"})
		return
	}
	org, err := model.GetAllOrganizations("")
	if err != nil {
		c.JSON(203, gin.H{"message": "查询失败."})
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
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 401 {object} error.Error "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /organization/personal/created
func CheckCreated(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid"})
		return
	}
	created, err := model.GetCreated(id)
	if err != nil {
		c.JSON(203, gin.H{"message": "Fail."})
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
// @Success 200 {object} []model.FollowingOrganization "{"msg":"查询成功"}"
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 401 {object} error.Error "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /organization/personal/following
func CheckFollowing(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid"})
		return
	}
	following, err := model.GetFollowing(id)
	if err != nil {
		c.JSON(203, gin.H{"message": "Fail."})
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
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 401 {object} error.Error "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /organization/details
type Detail struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func CheckDetails(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid"})
		return
	}

	var details Detail
	if err := c.BindJSON(&details); err != nil {
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
	}
	org, err := model.GetDetails(details.ID)
	if err != nil {
		c.JSON(203, gin.H{"message": "查无此组"})
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
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 401 {object} error.Error "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /organization/create
func CreateOne(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid"})
		return
	}
	var org model.Organization
	if err := c.BindJSON(&org); err != nil {
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
		return
	}
	result := model.DB.Create(&org)
	if result.Error != nil {
		c.JSON(400, "Fail.")
	}
	handler.SendResponse(c, "创建成功", nil)
}

func FollowOneOrganization(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid"})
		return
	}
	var following model.FollowingOrganization
	if err := c.BindJSON(&following); err != nil {
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
	}

	result := model.DB.Create(&following)
	if result.Error != nil {
		c.JSON(400, "Fail.")
	}
	handler.SendResponse(c, "创建成功", nil)
}
