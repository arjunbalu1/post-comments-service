package models

import "time"

type User struct {
	Username     string `gorm:"primaryKey;unique;not null" json:"username"`
	PasswordHash string `json:"-"` // Never expose password hash in JSON
	CreatedAt    time.Time
	Posts        []Post    `json:"posts"`
	Comments     []Comment `json:"comments"`
}
