package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ChatID   uint   `gorm:"not null"`
	SenderID uint   `gorm:"not null"`
	Content  string `gorm:"not null"`
	Chat     Chat   `gorm:"foreignKey:ChatID"`
}
