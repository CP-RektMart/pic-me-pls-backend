package model

import "time"

type Message struct {
	ID       uint      `gorm:"primaryKey"`
	ChatID   uint      `gorm:"not null"`
	SenderID uint      `gorm:"not null"`
	Content  string    `gorm:"not null"`
	SendTime time.Time `gorm:"autoCreateTime"`
	Chat     Chat      `gorm:"foreignKey:ChatID"`
}
