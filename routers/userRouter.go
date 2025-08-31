package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_task4/controllers"
)

// 用户消息
func RegisterUserRouter(r *gin.Engine) {
	group := r.Group("/user")

	obj := controllers.NewUserController()
	group.POST("/register", obj.RegisterUser)
	group.POST("/login", obj.LoginUser)
	group.GET("/info", obj.GetUser)
	group.PUT("/update", obj.UpdateUser)
	group.DELETE("/delete", obj.DeleteUser)
}
