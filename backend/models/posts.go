package models

import "gorm.io/gorm"

type Posts struct {
	gorm.Model
	Id                 string   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	PostsUserId        string   `gorm:"column:posts_user_id;not null"`
	User               User     `gorm:"foreignKey:PostsUserId;references:Id"`
	ApplicationsUserId []string `gorm:"column:applications_user_id;not null"`
	UserApplications   []User   `gorm:"many2many:posts_applications;foreignKey:Id;joinForeignKey:PostsId;References:Id;joinReferences:UserId"`
	Title              string   `gorm:"not null"`
	Description        string   `gorm:"not null"`
	MaxApplications    int      `gorm:"not null"`
	Deadline           string   `gorm:"not null"`
	Status             string   `gorm:"not null;default:'open'"` // 'open', 'closed', 'coming_soon'
	CreatedAt          string   `gorm:"not null"`
	UpdatedAt          string   `gorm:"not null"`
}
