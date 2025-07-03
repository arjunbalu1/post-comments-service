package handlers

import (
	"fmt"
	"net/http"

	"post-comments-service/database"
	"post-comments-service/models"

	"github.com/gin-gonic/gin"
)

// AddComment handles POST /posts/:id/comments
// It adds a new comment to a specific post. The comment content is saved as Markdown.
func AddComment(c *gin.Context) {
	var comment models.Comment
	postID := c.Param("id") // Get the post ID from the URL as a string

	// Bind the JSON body to the comment struct
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Input validation: require non-empty content
	if comment.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content is required"})
		return
	}

	// Check if the post exists in the database
	if err := database.DB.First(&models.Post{}, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Convert the postID from string to uint, because Comment.PostID is a uint
	// This is necessary because URL parameters are always strings
	comment.PostID = parseUint(postID)

	// Set the username from the authenticated user
	comment.Username = c.GetString("username")

	// Save the comment to the database
	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created comment as JSON
	c.JSON(http.StatusCreated, comment)
}

// parseUint is a helper to convert a string to uint
// Used to convert the postID from the URL (which is a string) to uint for the database
func parseUint(s string) uint {
	var n uint
	fmt.Sscanf(s, "%d", &n)
	return n
}
