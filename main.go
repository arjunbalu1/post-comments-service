package main

import (
	"post-comments-service/database"
	"post-comments-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := handlers.AuthMiddleware()

	// Post routes
	r.POST("/posts", auth, handlers.CreatePost)
	r.GET("/posts", handlers.ListPosts)
	r.GET("/posts/:id", handlers.GetPost)

	// Comment routes
	r.POST("/posts/:id/comments", auth, handlers.AddComment)

	r.Run(":8080")
}
