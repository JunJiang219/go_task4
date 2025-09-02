package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_task4/controllers"
)

// 用户路由
func RegisterUserRouter(r *gin.Engine) {
	group := r.Group("/user")

	obj := controllers.NewUserController()
	group.POST("/register", obj.RegisterUser)
	group.POST("/login", obj.LoginUser)
	group.GET("/info", obj.ReadUser)
	group.PUT("/update", obj.UpdateUser)
	group.DELETE("/delete", obj.DeleteUser)
}
