package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_task4/controllers"
)

// 文章路由
func RegisterCommentRouter(r *gin.Engine) {
	group := r.Group("/comment")

	obj := controllers.NewCommentController()
	group.POST("/create", obj.CreateComment)
	group.GET("/info-list", obj.ReadCommentList)
}
