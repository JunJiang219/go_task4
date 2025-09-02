package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_task4/controllers"
)

// 文章路由
func RegisterPostsRouter(r *gin.Engine) {
	group := r.Group("/posts")

	obj := controllers.NewPostsController()
	group.POST("/create", obj.CreatePosts)
	group.GET("/info-list", obj.ReadPostsList)
	group.GET("/info", obj.ReadPosts)
	group.PUT("/update", obj.UpdatePosts)
	group.DELETE("/delete", obj.DeletePosts)
}
