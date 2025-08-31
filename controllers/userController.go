package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_task4/models"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

func (obj UserController) RegisterUser(c *gin.Context) {
	c.String(http.StatusOK, "注册用户")
	user := models.User{}
	c.ShouldBind(&user)
}

func (obj UserController) GetUser(c *gin.Context) {
	c.String(http.StatusOK, "获取用户")
}

func (obj UserController) UpdateUser(c *gin.Context) {
	c.String(http.StatusOK, "更新用户")
}

func (obj UserController) DeleteUser(c *gin.Context) {
	c.String(http.StatusOK, "删除用户")
}

func (obj UserController) LoginUser(c *gin.Context) {
	c.String(http.StatusOK, "登录用户")
}
