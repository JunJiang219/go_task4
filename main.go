package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_task4/pkg/errcode"
	"github.com/go_task4/routers"
)

func main() {
	fmt.Println(errcode.ErrPostsNotFound)

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    errcode.Success,
			"message": "欢迎访问首页",
		})
	})
	routers.RegisterUserRouter(r)
	r.Run()
}
