package router

import (
	"blackboard/handler/announcement"
	"blackboard/handler/organization"
	"blackboard/handler/user"
	"blackboard/router/middleware"
	"net/http"
	"path/filepath"

	"github.com/Wishforpeace/My-Tool/utils"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.Static("/statics", "./statics/")
	r.StaticFS("/avatar", http.Dir(filepath.Join(utils.RootPath(), "avatar")))
	//登录
	auth := r.Group("/api/v1/login")
	{
		auth.POST("", user.Login)
	}

	//user:
	g1 := r.Group("/api/v1/user").Use(middleware.Auth())
	{
		//查看信息
		g1.GET("/info", user.UserInfo)

		//修改用户名
		g1.PUT("/changename", user.ChangeUserName)

		//修改用户头像
		g1.POST("/update", user.UpdateUserProfile)

		//用户收藏的通告
		g1.GET("/collection", user.CheckCollections)

		//用户发布的通告
		g1.GET("/published", user.UserPublished)

	}

	//organizations
	g2 := r.Group("/api/v1/organization").Use(middleware.Auth())
	{
		//查看所有组织
		g2.GET("", organization.CheckAll)

		//查看创建的组织信息
		g2.GET("/personal/created", organization.CheckCreated)

		//查看关注的组织
		g2.GET("/personal/following", organization.CheckFollowing)

		//查看指定组织信息
		g2.GET("/details", organization.CheckDetails)

		//创建新的组织
		g2.POST("/create", organization.CreateOne)

		//上传组织头像
		g2.POST("/:organization_name/image", organization.UploadImage)

		//关注新的组织
		g2.POST("/follow", organization.FollowOneOrganization)
	}
	//announcement
	g3 := r.Group("/api/v1/announcement").Use(middleware.Auth())
	{
		//查看最新通知
		g3.GET("", announcement.CheckAllPubilshed)

		//新建分组
		g3.POST("/group", announcement.CreateGroup)

		//发布通知
		g3.POST("/publish", announcement.PublishNews)

		//删除通知
		g3.DELETE("/delete/:announcement_id", announcement.DeletePublished)

		//收藏通知
		g3.POST("/collect", announcement.Collect)

		//取消收藏
		g3.DELETE("/collect/cancel/:collect_id", announcement.CancelCollect)
	}
}
