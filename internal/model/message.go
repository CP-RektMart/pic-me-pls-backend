package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	SenderID   uint   `gorm:"not null"`
	Sender     User   `gorm:"foreignKey:SenderID"`
	ReceiverID uint   `gorm:"not null"`
	Receiver   User   `gorm:"foreignKey:ReceiverID"`
	Content    string `gorm:"not null"`
}
