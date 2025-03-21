package model

import (
	"gorm.io/gorm"
)

type MessageType string

const (
	MessageTypeText      MessageType = "TEXT"
	MessageTypeImage     MessageType = "IMAGE"
	MessageTypeQuotation MessageType = "QUOTATION"
	MessageTypePreview   MessageType = "PREVIEW"
)

type Message struct {
	gorm.Model
	Type    MessageType `gorm:"not null"`
	Content string      `gorm:"not null"`

	SenderID   uint `gorm:"not null"`
	ReceiverID uint `gorm:"not null"`
	Sender     User `gorm:"foreignKey:SenderID"`
	Receiver   User `gorm:"foreignKey:ReceiverID"`
}

func ValidateMessageType(msgType string) bool {
	switch MessageType(msgType) {
	case MessageTypeText, MessageTypeImage, MessageTypeQuotation, MessageTypePreview:
		return true
	default:
		return false
	}
}
