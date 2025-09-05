package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go_task4/models"
	"github.com/go_task4/utils"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

type PostsController struct{}

func NewPostsController() PostsController {
	return PostsController{}
}

func (obj PostsController) CreatePosts(c *gin.Context) {
	user_id := c.GetUint("user_id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body missing title"})
		return
	}

	if content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body missing content"})
		return
	}

	posts := models.Posts{
		Title:   title,
		Content: content,
		UserID:  user_id,
	}

	if err := models.GetDB().Create(&posts).Error; err != nil {
		utils.GetLogger().Error(
			"CreatePosts failed",
			zap.String("error", err.Error()),
			zap.Any("posts", posts),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create posts"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posts create successfully"})
}

func (obj PostsController) ReadPostsList(c *gin.Context) {
	var pl []models.Posts
	if err := models.GetDB().Find(&pl).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts list"})
		return
	}

	type PostsIntro struct {
		PostsID uint   `json:"posts_id"`
		Title   string `json:"title"`
	}

	var list = make([]PostsIntro, len(pl))
	for i, v := range pl {
		list[i].PostsID = v.ID
		list[i].Title = v.Title
	}

	c.JSON(http.StatusOK, gin.H{"posts_list": list})
}

func (obj PostsController) ReadPosts(c *gin.Context) {
	idStr := c.Query("posts_id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request params missing posts_id"})
		return
	}

	postsID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid posts_id"})
		return
	}

	posts := models.Posts{}
	posts.ID = uint(postsID)

	if err := models.GetDB().Preload("Comments").Find(&posts).Error; err != nil {
		utils.GetLogger().Info(
			"ReadPosts failed",
			zap.Uint("postsID", posts.ID),
		)
		errStr := fmt.Sprintf("Failed to get posts, id = %v", postsID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errStr})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (obj PostsController) UpdatePosts(c *gin.Context) {
	idStr := c.PostForm("posts_id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body missing posts_id"})
		return
	}

	postsID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid posts_id"})
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")
	if title == "" && content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body missing title or content"})
		return
	}

	user_id := c.GetUint("user_id")
	posts := models.Posts{}
	posts.ID = uint(postsID)
	if err := models.GetDB().First(&posts).Error; err != nil {
		utils.GetLogger().Info(
			"UpdatePosts failed -- no posts record",
			zap.Uint("posts_id", posts.ID),
		)
		errStr := fmt.Sprintf("Failed to get posts, id = %v", postsID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errStr})
		return
	}

	if posts.UserID != user_id {
		utils.GetLogger().Warn(
			"UpdatePosts not allowed",
			zap.Uint("user_id", user_id),
			zap.Uint("belong_userID", posts.UserID),
		)
		errStr := fmt.Sprintf("Can't update other's posts, {my: %v, other: %v}", user_id, posts.UserID)
		c.JSON(http.StatusForbidden, gin.H{"error": errStr})
		return
	}

	posts.Title = title
	posts.Content = content
	if err := models.GetDB().Updates(&posts).Error; err != nil {
		utils.GetLogger().Error(
			"UpdatePosts failed",
			zap.String("error", err.Error()),
			zap.Any("posts", posts),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Posts update successfully"})
}

func (obj PostsController) DeletePosts(c *gin.Context) {
	idStr := c.Query("posts_id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request params missing posts_id"})
		return
	}

	postsID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid posts_id"})
		return
	}

	user_id := c.GetUint("user_id")
	posts := models.Posts{}
	posts.ID = uint(postsID)
	if err := models.GetDB().First(&posts).Error; err != nil {
		utils.GetLogger().Info(
			"DeletePosts failed -- no posts record",
			zap.Uint("posts_id", posts.ID),
		)
		errStr := fmt.Sprintf("Failed to get posts, id = %v", postsID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errStr})
		return
	}

	if posts.UserID != user_id {
		utils.GetLogger().Warn(
			"DeletePosts not allowed",
			zap.Uint("user_id", user_id),
			zap.Uint("belong_userID", posts.UserID),
		)
		errStr := fmt.Sprintf("Can't delete other's posts, {my: %v, other: %v}", user_id, posts.UserID)
		c.JSON(http.StatusForbidden, gin.H{"error": errStr})
		return
	}

	if err := models.GetDB().Select(clause.Associations).Delete(&posts).Error; err != nil {
		utils.GetLogger().Error(
			"DeletePosts failed",
			zap.String("error", err.Error()),
			zap.Any("posts", posts),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Posts delete successfully"})
}
