package announcement

import (
	"blackboard/handler"
	"blackboard/model"
	"blackboard/pkg/errno"

	"github.com/gin-gonic/gin"
)

// @Summary 查看通知
// @Tags announcement
// @Description 用户查看已经发布过的通知
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Announcement "{"msg":"查看"}"
// @Failure 203 {object} errno.Errno "{"error_code":"20001","message":"Fail."}"
// @Failure 401 {object} errno.Errno "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
// @Router /api/v1/announcement [get]
func CheckAllPubilshed(c *gin.Context) {
	announcement, err := model.GetAnnouncements(" ")
	if err != nil {
		handler.SendError(c, errno.ErrQuery, nil)
		return
	}
	handler.SendResponse(c, "查看成功", announcement)
}

// @Summary 发布通知
// @Tags announcement
// @Description 仅组织创建者可发布新的通知
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param announcement body model.Announcement true "组织创建者发布的新通知"
// @Success 200 {object} []model.Announcement "{"msg":"创建成功"}"
// @Failure 400 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Failure 412 {object} errno.Errno "{"msg":"身份认证失败"}"
// @Router /api/v1/announcement/publish [post]
func PublishNews(c *gin.Context) {
	id, ok := c.Get("student_id")
	if !ok {
		handler.SendBadRequest(c, "未输入身份", nil)
	}
	ID := id.(string)
	var announcement model.Announcement
	err := c.BindJSON(&announcement)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil)
		return
	}
	verify := model.JudgeFounder(ID, announcement.OrganizationID)
	if !verify {
		result := model.DB.Create(&announcement)
		if result.Error != nil {
			handler.SendError(c, errno.ErrDatabase, nil)
			return
		}
	} else {
		c.JSON(412, gin.H{
			"message": "身份认证错误",
		})
	}
	handler.SendResponse(c, "创建成功", announcement)
}

//@Summary 创建分组
//@Tags announcement
//@Description 仅组织创建者可新建通告分组
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param group body model.Group ture "新建分组"
// @Success 200 {object} []model.Group "{"msg":"创建成功"}"
// @Failure 400 {object} errno.Errno
// @Failure 412 {object} errno.Errno "{"msg":"身份认证失败"}"
// @Router /api/v1/announcement/create_group [post]
func CreateGroup(c *gin.Context) {
	id, ok := c.Get("student_id")
	if !ok {
		handler.SendBadRequest(c, "未输入身份", "null")
	}
	ID := id.(string)
	var group model.Group
	err := c.BindJSON(&group)
	if err != nil || group.OrganizationID == "" {
		handler.SendBadRequest(c, errno.ErrBind, nil)
		return
	}
	verfify := model.JudgeFounder(ID, group.OrganizationID)
	if verfify {
		result := model.DB.Create(&group)
		if result.Error != nil {
			handler.SendError(c, errno.ErrDatabase, nil)
		}
	} else {
		c.JSON(412, gin.H{
			"message": "身份认证失败",
		})
	}
	handler.SendResponse(c, "创建成功", group)
}

// @Summary 删除通知
// @Tags announcement
// @Description 仅组织创建者可删除通告
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param announcement_id path string true "通知ID"
// @Success 200  "{"msg":"删除成功"}"
// @Failure 400 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Failure 203 {object} errno.Errno "{"error_code":"20001","message":"Fail."}"
// @Failure 412 {object} errno.Errno "{"msg":"身份认证失败"}"
// @Router /api/v1/announcement/delete [delete]
func DeletePublished(c *gin.Context) {
	id, ok := c.Get("student_id")
	if !ok {
		handler.SendBadRequest(c, "未输入身份", "null")
	}
	ID := id.(string)
	AnnoucementID := c.Param("announcement_id")
	verfify := model.JudgePublisher(ID, AnnoucementID)
	if verfify {
		err := model.DeleteAnnoucement(AnnoucementID)
		if err != nil {
			handler.SendError(c, errno.ErrDatabase, nil)
			return
		}
	} else {
		c.JSON(412, gin.H{
			"message": "身份认证错误",
		})
		return
	}
	handler.SendResponse(c, "删除成功", nil)
}

// @Summary 收藏通知
// @Tags announcement
// @Description 用户将通知加入自己的收藏
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Collection "{"msg":"收藏成功"}"
// @Failure 203 {object} errno.Errno  "{"error_code":"20001","message":"Fail."}"
// @Failure 400 {object} errno.Errno
// @Failure 412 {object} errno.Errno  "{"msg":"身份认证失败"}"
// @Router /api/v1/announcement/collect [post]
func Collect(c *gin.Context) {
	id, _ := c.Get("student_id")
	ID := id.(string)
	var collect model.Collection
	collect.StudentID = ID
	if err := c.BindJSON(&collect); err != nil && collect.AnnouncementID == "" {
		handler.SendBadRequest(c, errno.ErrBind, nil)
		return
	}
	result := model.DB.Create(&collect)
	if result.Error != nil {
		c.JSON(400, gin.H{"message": "Fail"})
		return
	}
	handler.SendResponse(c, "关注成功", nil)
}

// @Summary 取消收藏
// @Tags announcement
// @Descrip 用户删除之前收藏的通知
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param collect_id query string true "collect_id"
// @Success 200 {object} []model.Collection "{"msg":"取消成功"}"
// @Failure 200 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/announcement/collect/cancel [delete]
func CancelCollect(c *gin.Context) {
	CollectID := c.Query("collect_id")
	err := model.CancelCollect(CollectID)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, "取消成功", nil)
}
