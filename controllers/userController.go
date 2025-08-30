package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

func (obj UserController) CreateUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "创建用户")
}

func (obj UserController) GetUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "获取用户")
}

func (obj UserController) UpdateUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "更新用户")
}

func (obj UserController) DeleteUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "删除用户")
}

func (obj UserController) LoginUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "登录用户")
}
