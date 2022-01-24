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
	verify := model.JudgeFounder(id, announcement.OrganizationID)
	if !verify {
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

//@Summary 创建分组
//@Tags announcement
//@Description 仅组织创建者可新建通告分组
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Group "{"msg":"创建成功"}"
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 412 {object}  "{"msg":"身份认证失败"}"
// @Router /announcement/create_group
func CreateGroup(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "Token Invalid"})
	}
	var group model.Group
	err = c.BindJSON(&group)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
		return
	}
	verfify := model.JudgeFounder(id, group.OrganizationID)
	if verfify {
		result := model.DB.Create(&group)
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
	handler.SendResponse(c, "创建成功", group)
}

//@Summary 删除通知
//@Tags announcement
//@Description 仅组织创建者可删除通告
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Group "{"msg":"删除成功"}"
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 412 {object}  "{"msg":"身份认证失败"}"
// @Router /announcement/delete
func DeletePublished(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "Token Invalid"})
	}
	AnnoucementID := c.Param("announcement_id")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Lack Param Or Param Not Satisfiable.",
		})
		return
	}
	verfify := model.JudgePublisher(id, AnnoucementID)
	if verfify {
		err = model.DeleteAnnoucement(AnnoucementID)
		if err != nil {
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
	handler.SendResponse(c, "删除成功", nil)
}

//@Summary 收藏通知
//@Tags announcement
//@Description 用户将通知加入自己的收藏
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Collection "{"msg":"收藏成功"}"
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 412 {object}  "{"msg":"身份认证失败"}"
// @Router /announcement/collect
func Collect(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "Token Invalid"})
	}
	var collect model.Collection
	if err = c.BindJSON(&collect); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	result := model.DB.Create(&collect)
	if result.Error != nil {
		c.JSON(400, gin.H{"message": "Fail"})
		return
	}
	handler.SendResponse(c, "关注成功", nil)
}

//@Summary 取消收藏
//@Tags announcement
//@Descrip 用户删除之前收藏的通知
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Collection "{"msg":"取消收藏"}"
// @Failure 203 {object} error.Error "{"error_code":"20001","message":"Fail."}"
// @Failure 412 {object}  "{"msg":"身份认证失败"}"
// @Router /announcement/collect/cancel
func CancelCollect(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := user.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "Token Invalid"})
	}
	CollectID := c.Query("collect_id")
	err = model.CancelCollect(CollectID)
	handler.SendResponse(c, "关注成功", nil)

}
