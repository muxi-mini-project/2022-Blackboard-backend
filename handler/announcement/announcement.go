package announcement

import (
	"blackboard/handler"
	"blackboard/model"
	"blackboard/service/user"

	"github.com/gin-gonic/gin"
)

//@Summary 查看通知
//@Tags announcement
//@Description 用户查看已经发布过的通知
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Announcement "{"msg":"查看"}"
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 401 {object} error.Error "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
// @Router /announcement
func CheckAllPubilshed(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "认证失败"})
		return
	}
	announcement, err := model.GetAnnouncements("")
	if err != nil {
		c.JSON(203, gin.H{"message": "Fail."})
		return
	}
	handler.SendResponse(c, "查看成功", announcement)

}

//@Summary 发布通知
//@Tags announcement
//@Description 仅组织创建者可发布新的通知
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Announcement "{"msg":"创建成功"}"
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 412 {object}  "{"msg":"身份认证失败"}"
// @Router /announcement/publish
func PublishNews(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "Token Invalid"})
	}
	var announcement model.Announcement
	err = c.BindJSON(&announcement)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
		return
	}
	verfi := model.JudgeFounder(id, announcement.OrganizationID)
	if verfi {
		result := model.DB.Create(&announcement)
		if result.Error != nil {
			c.JSON(400, gin.H{
				"message": "Fail.",
			})
		}
	} else {
		c.JSON(412, gin.H{
			"message": "身份认证错误",
		})
		return
	}
	handler.SendResponse(c, "创建成功", announcement)
}

func CreateGroup(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "Token Invalid"})
	}
	var announcement model.Announcement
	err = c.BindJSON(&announcement)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
		return
	}
	verfi := model.JudgeFounder(id, announcement.OrganizationID)
	if verfi {
		result := model.DB.Create(&announcement)
		if result.Error != nil {
			c.JSON(400, gin.H{
				"message": "Fail.",
			})
		}
	} else {
		c.JSON(412, gin.H{
			"message": "身份认证错误",
		})
		return
	}
	handler.SendResponse(c, "创建成功", announcement)
}
