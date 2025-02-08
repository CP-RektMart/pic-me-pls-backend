package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name              string `gorm:"not null"`
	Email             string `gorm:"not null;unique"`
	PhoneNumber       string `gorm:"size:10"`
	ProfilePictureURL string
	Role              string      `gorm:"not null"`
	SenderMessages    []Message   `gorm:"foreignKey:SenderID"`
	ReceiverMessages  []Message   `gorm:"foreignKey:ReceiverID"`
	Quotations        []Quotation `gorm:"foreignKey:CustomerID"`
	Reviews           []Review    `gorm:"foreignKey:CustomerID"`
}
