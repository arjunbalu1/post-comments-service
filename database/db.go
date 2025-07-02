package database

import (
	"fmt"
	"log"
	"os"

	"post-comments-service/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "host=localhost user=postgres dbname=post_comments_service sslmode=disable password=postgres"
	}
	var err error
	DB, err = gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB.AutoMigrate(&models.Post{}, &models.Comment{})
	fmt.Println("Database connection established and migrated")
}
