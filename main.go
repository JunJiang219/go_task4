package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_task4/middlewares"
	"github.com/go_task4/models"
	"github.com/go_task4/routers"
)

func main() {
	models.AutoMigrate()

	r := gin.Default()
	r.Use(middlewares.JWTAuthMiddleware)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "欢迎访问首页"})
	})
	routers.RegisterUserRouter(r)
	r.Run()
}
