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
// @Param Authorization header string true "token"
// @Success 200 {object} []model.Announcement "{"msg":"查看"}"
// @Failure 203 {object} errno.Errno "{"error_code":"20001","message":"Fail."}"
// @Failure 401 {object} errno.Errno "{"error_code":"10001","message":"Token Invalid."} 身份验证失败 重新登录"
// @Router /announcement [get]
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
// @Param Authorization header string true "token"
// @Param announcement body model.Announcement true "组织创建者发布的新通知"
// @Success 200 {object} []model.Announcement "{"msg":"创建成功"}"
// @Failure 400 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Failure 412 {object} errno.Errno "{"msg":"身份认证失败"}"
// @Router /announcement/publish [post]
func PublishNews(c *gin.Context) {
	ID := c.MustGet("student_id").(string)
	var announcement model.Announcement
	err := c.BindJSON(&announcement)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil)
		return
	}
	announcement.PublisherID = ID
	announcement.OrganizationID = model.GetOrgID(announcement.OrganizationName)
	announcement.GroupID = model.GetGroupID(announcement.GroupName, announcement.OrganizationName)
	verify := model.JudgeFounder(ID, announcement.OrganizationID)
	if !verify {
		c.JSON(412, gin.H{
			"message": "身份认证错误",
		})
		return

	} else {
		result := model.DB.Create(&announcement)
		if result.Error != nil {
			handler.SendError(c, errno.ErrDatabase, nil)
			return
		}
	}
	handler.SendResponse(c, "创建成功", announcement)
}

//@Summary 创建分组
//@Tags announcement
//@Description 仅组织创建者可新建通告分组
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param group body model.Grouping ture "新建分组"
// @Success 200 {object} []model.Grouping "{"msg":"创建成功"}"
// @Failure 400 {object} errno.Errno
// @Failure 412 {object} errno.Errno "{"msg":"身份认证失败"}"
// @Router /announcement/group [post]
func CreateGroup(c *gin.Context) {
	ID := c.MustGet("student_id").(string)
	var group model.Grouping
	err := c.BindJSON(&group)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil)
		return
	}
	verfify := model.JudgeFounder(ID, group.OrganizationID)
	if !verfify {
		c.JSON(412, gin.H{
			"message": "身份认证失败",
		})
		return
	}
	group.OrganizationID = model.GetOrgID(group.OrganizationName)
	result := model.DB.Create(&group)
	if result.Error != nil {
		handler.SendBadRequest(c, "Fail", nil)
		return
	}
	handler.SendResponse(c, "创建成功", group)
}

// @Summary 删除通知
// @Tags announcement
// @Description 仅组织创建者可删除通告
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param announcement_id path string true "通知ID"
// @Success 200  "{"msg":"删除成功"}"
// @Failure 400 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Failure 203 {object} errno.Errno "{"error_code":"20001","message":"Fail."}"
// @Failure 412 {object} errno.Errno "{"msg":"身份认证失败"}"
// @Router /announcement/delete/:announcement_id [delete]
func DeletePublished(c *gin.Context) {
	ID := c.MustGet("student_id").(string)
	AnnoucementID := c.Param("announcement_id")
	verfify := model.JudgePublisher(ID, AnnoucementID)
	if verfify {
		err := model.DeleteAnnoucement(AnnoucementID)
		if err != nil {
			handler.SendError(c, errno.ErrDatabase, nil)
			return
		} else {
			handler.SendResponse(c, "删除成功", nil)
		}
	} else {
		c.JSON(412, gin.H{
			"message": "身份认证错误",
		})
		return
	}

}

// @Summary 收藏通知
// @Tags announcement
// @Description 用户将通知加入自己的收藏
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []model.Collection "{"msg":"收藏成功"}"
// @Failure 203 {object} errno.Errno  "{"error_code":"20001","message":"Fail."}"
// @Failure 400 {object} errno.Errno
// @Failure 412 {object} errno.Errno  "{"msg":"身份认证失败"}"
// @Router /announcement/collect [post]
func Collect(c *gin.Context) {
	ID := c.MustGet("student_id").(string)
	var collect model.Collection
	if err := c.BindJSON(&collect); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil)
		return
	}
	collect.StudentID = ID
	collect.Announcement = model.CheckAnnouce(collect.AnnouncementID)
	result := model.DB.Create(&collect)
	if result.Error != nil {
		c.JSON(400, gin.H{"message": "Fail"})
		return
	} else {
		handler.SendResponse(c, "关注成功", collect)
	}

}

// @Summary 取消收藏
// @Tags announcement
// @Descrip 用户删除之前收藏的通知
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param collect_id path string true "collect_id"
// @Success 200 {object} []model.Collection "{"msg":"取消成功"}"
// @Failure 200 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /announcement/collect/cancel/:collect_id [delete]
func CancelCollect(c *gin.Context) {
	CollectID := c.Param("collect_id")
	err := model.CancelCollect(CollectID)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, "取消成功", err)
}
