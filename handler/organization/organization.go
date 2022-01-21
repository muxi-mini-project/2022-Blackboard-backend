package organization

import (
	"blackboard/handler"
	"blackboard/model"
	"blackboard/service/user"

	"github.com/gin-gonic/gin"
)

// @Summary  查看组织
// @Tags organization
// @Description 查看用户创建过的组织
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Organization "{"msg":"查询"}"
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 401 {object} error.Error "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001","message":"Fail."}or {"error_code":"00002","message":"Lack Param or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /organization
func CheckCreated(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid"})
		return
	}
	created, err := model.GetCreated(id)
	if err != nil {
		c.JSON(203, gin.H{"message": "Token Invalid"})
		return
	}
	handler.SendResponse(c, "获取成功", created)
}

func CheckFollowing(c *gin.Context) {

}

func CheckDetails(c *gin.Context) {

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
