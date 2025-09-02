package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_task4/middlewares"
	"github.com/go_task4/models"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

func (obj UserController) RegisterUser(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := models.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (obj UserController) LoginUser(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storedUser := models.User{}
	if err := models.GetDB().Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	tokenStr, err := middlewares.NewJWTAuth(storedUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

func (obj UserController) ReadUser(c *gin.Context) {
	c.String(http.StatusOK, "读取用户")
}

func (obj UserController) UpdateUser(c *gin.Context) {
	c.String(http.StatusOK, "更新用户")
}

func (obj UserController) DeleteUser(c *gin.Context) {
	c.String(http.StatusOK, "删除用户")
}
