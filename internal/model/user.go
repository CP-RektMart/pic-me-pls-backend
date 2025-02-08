package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name              string `gorm:"not null"`
	PhoneNumber       string `gorm:"size:10"`
	ProfilePictureURL string
	Role              string `gorm:"not null"`
}
