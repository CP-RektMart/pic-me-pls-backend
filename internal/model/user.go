package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID                uint   `gorm:"primaryKey"`
	Name              string `gorm:"not null"`
	PhoneNumber       string `gorm:"size:10"`
	ProfilePictureURL string
	Role              string `gorm:"not null"`
}
