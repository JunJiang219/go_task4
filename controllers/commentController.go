package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go_task4/models"
)

type CommentController struct{}

func NewCommentController() CommentController {
	return CommentController{}
}

func (obj CommentController) CreateComment(c *gin.Context) {
	user_id := c.GetUint("user_id")
	content := c.PostForm("content")
	if content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body missing content"})
		return
	}

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

	comment := models.Comment{
		Content: content,
		UserID:  user_id,
		PostsID: uint(postsID),
	}

	if err := models.GetDB().Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment create successfully"})
}

func (obj CommentController) ReadCommentList(c *gin.Context) {
	idStr := c.Query("posts_id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body missing posts_id"})
		return
	}

	postsID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid posts_id"})
		return
	}

	var cs []models.Comment
	if err := models.GetDB().Where("posts_id = ?", postsID).Find(&cs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts list"})
		return
	}

	c.JSON(http.StatusOK, cs)
}
