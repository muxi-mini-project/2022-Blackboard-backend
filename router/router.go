package router

import (
	// "blackboard/handler/organization"
	// "blackboard/handler"
	"blackboard/handler/user"
	"github.com/gin-gonic/gin"
	"net/http"
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

		//修改用户名
		g1.PUT("/changename", user.ChangeUserName)

		// //修改用户头像
		// g1.PUT("", user.ChangeHeadportrait)

		// //用户收藏
		// g1.GET("/collection/:id", user.UserCollection)

		// //用户发布
		// g1.GET("/published/:id", user.UserPublish)

	}

	// //organizations
	// g2 := r.Group("/api/v1/organization")
	// {
	// 	//查看创建的组织
	// 	g2.GET("/personal/:user_id/created", organization.CheckCreated)
	// 	//查看关注的组织
	// 	g2.GET("/personal/:user_id/following", organization.CheckFollowing)
	// 	//查看特定组织信息
	// 	g2.GET("/personal/:user_id/created/:organization_id", organization.CheckDetails)
	// 	g2.GET("/personal/:user_id/following/:organization_id", organization.CheckDetails)
	// }
	// //announcement
	// g3 := r.Group("/api/v1/announcement")
	// {
	// 	//查看最新通知
	// 	g3.GET("", announcement.CheckAllPubilshed)
	// 	//发布通知
	// 	g3.POST("/:user_id/publish/:organization_id", announcement.PublishNews)
	// 	//删除通知
	// 	g3.DELETE("/:user_id/delete/:organization_id", announcement.DeletePublished)
	// 	//收藏通知
	// 	g3.POST("/:user_id/collect/:organization_id/:announcement_id", announcement.Collect)
	// }
}
