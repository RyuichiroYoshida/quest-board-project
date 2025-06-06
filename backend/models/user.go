package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string `gorm:"not null"`
	DiscordId string `gorm:"not null;unique"`
	CreatedAt string `gorm:"not null"`
}
