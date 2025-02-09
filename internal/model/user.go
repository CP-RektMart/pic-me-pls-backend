package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name              string `gorm:"not null"`
	PhoneNumber       string `gorm:"size:10"`
	ProfilePictureURL string
	Role              string      `gorm:"not null"`
	SenderMessages    []Message   `gorm:"foreignKey:SenderID"`
	ReceiverMessages  []Message   `gorm:"foreignKey:ReceiverID"`
	Quotations        []Quotation `gorm:"foreignKey:CustomerID"`
	Reviews           []Review    `gorm:"foreignKey:CustomerID"`
}

type UserDto struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	PhoneNumber       string `json:"phone_number"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Role              string `json:"role"`
}
