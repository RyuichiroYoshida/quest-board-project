package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Id                 string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	PostsUserId        string         `gorm:"column:posts_user_id;not null"`
	User               User           `gorm:"foreignKey:PostsUserId;references:Id"`
	ApplicationsUserId pq.StringArray `gorm:"type:text[];not null"`
	UserApplications   []User         `gorm:"many2many:posts_applications;foreignKey:Id;joinForeignKey:PostsId;References:Id;joinReferences:UserId"`
	Title              string         `gorm:"not null"`
	Description        string         `gorm:"not null"`
	MaxApplications    int            `gorm:"not null"`
	Deadline           string         `gorm:"not null"`
	Status             string         `gorm:"not null;default:'open'"` // 'open', 'closed', 'coming_soon'
	CreatedAt          time.Time      `gorm:"not null"`
	UpdatedAt          time.Time      `gorm:"not null"`
}
