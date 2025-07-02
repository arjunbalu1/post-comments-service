package main

import (
	"post-comments-service/database"
	"post-comments-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	// Post routes
	r.POST("/posts", handlers.CreatePost)
	r.GET("/posts", handlers.ListPosts)
	r.GET("/posts/:id", handlers.GetPost)

	// Comment routes
	r.POST("/posts/:id/comments", handlers.AddComment)

	r.Run(":8080")
}
