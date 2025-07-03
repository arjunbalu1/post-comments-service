package models

import (
	"time"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	PostID    uint   `json:"post_id"`
	Content   string `json:"content"`
	CreatedAt time.Time
	Username  string `json:"username"`
}
