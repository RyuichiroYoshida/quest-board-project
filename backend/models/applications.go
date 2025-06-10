package models

import (
	"time"

	"gorm.io/gorm"
)

type Applications struct {
	gorm.Model
	Id        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserId    string    `gorm:"column:user_id;not null"`
	User      User      `gorm:"foreignKey:UserId;references:Id"`
	PostsId   string    `gorm:"column:posts_id;not null"`
	Posts     Posts     `gorm:"foreignKey:PostsId;references:Id"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
