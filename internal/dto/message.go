package dto

import (
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
)

type RealTimeMessageRequest struct {
	Type       model.MessageType `json:"type" validate:"required,messageType"`
	Content    string            `json:"content" validate:"required"`
	ReceiverID uint              `json:"receiverId" validate:"required"`
}

type RealTimeMessageResponse struct {
	ID         uint              `json:"id"`
	Type       model.MessageType `json:"type"`
	Content    string            `json:"content"`
	ReceiverID uint              `json:"receiverId"`
	SenderID   uint              `json:"senderId"`
	SendedAt   time.Time         `json:"sendedAt"`
}

func ToMessageModel(senderID uint, message RealTimeMessageRequest) model.Message {
	return model.Message{
		Type:       message.Type,
		Content:    message.Content,
		SenderID:   senderID,
		ReceiverID: message.ReceiverID,
	}
}

func ToRealTimeMessageResponse(message model.Message) RealTimeMessageResponse {
	return RealTimeMessageResponse{
		ID:         message.ID,
		Type:       message.Type,
		Content:    message.Content,
		SenderID:   message.SenderID,
		ReceiverID: message.ReceiverID,
		SendedAt:   message.CreatedAt,
	}
}
