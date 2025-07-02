package handlers

import (
	"net/http"

	"post-comments-service/database"
	"post-comments-service/models"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

// CreatePost handles POST /posts
func CreatePost(c *gin.Context) {
	var post models.Post
	// Bind JSON body to post struct
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save post to database
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

// ListPosts handles GET /posts
func ListPosts(c *gin.Context) {
	var posts []models.Post
	// Preload comments for each post
	if err := database.DB.Preload("Comments").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPost handles GET /posts/:id
func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := database.DB.Preload("Comments").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	// Render Markdown to HTML for each comment
	for i, comment := range post.Comments {
		renderer := html.NewRenderer(html.RendererOptions{})
		htmlContent := markdown.ToHTML([]byte(comment.Content), nil, renderer)
		post.Comments[i].Content = string(htmlContent)
	}
	c.JSON(http.StatusOK, post)
}
