package router

import (
	// "blackboard/handler/organization"
	// "blackboard/handler"
	"blackboard/handler/announcement"
	"blackboard/handler/organization"
	"blackboard/handler/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})
	//user:
	g1 := r.Group("/api/v1/user")
	{
		//登录
		g1.POST("", user.Login)

		//查看信息
		g1.GET("Info", user.UserInfo)

		//修改用户名
		g1.PUT("/changename", user.ChangeUserName)

		// //修改用户头像
		// g1.PUT("", user.ChangeHeadportrait)

		//用户收藏的通告
		g1.GET("/collection", user.CheckCollections)

		//用户发布的通告
		g1.GET("/published", user.UserPublished)

	}

	//organizations
	g2 := r.Group("/api/v1/organization")
	{
		//查看所有组织
		g2.GET("", organization.CheckAll)

		//查看创建的组织
		g2.GET("/personal/created", organization.CheckCreated)

		//查看关注的组织
		g2.GET("/personal/following", organization.CheckFollowing)

		//查看指定组织信息
		g2.GET("/details", organization.CheckDetails)

		//创建新的组织
		g2.POST("/create", organization.CreateOne)

		//关注新的组织
		g2.POST("/following", organization.FollowOneOrganization)
	}
	//announcement
	g3 := r.Group("/api/v1/announcement")
	{
		//查看最新通知
		g3.GET("", announcement.CheckAllPubilshed)
		//新建分组
		g3.POST("/create_group", announcement.CreateGroup)
		//发布通知
		g3.POST("/publish", announcement.PublishNews)
		//删除通知
		g3.DELETE("/:user_id/delete/:organization_id", announcement.DeletePublished)
		//收藏通知
		g3.POST("/:user_id/collect/:organization_id/:announcement_id", announcement.Collect)
	}
}
