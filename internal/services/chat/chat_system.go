package chat

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
)

type EventType string

const (
	EventError   EventType = "ERROR"
	EventMessage EventType = "MESSAGE"
)

type ChatSystem struct {
	clients map[uint]chan string
	store   *database.Store
}

var (
	instance *ChatSystem
	once     sync.Once
)

func NewChatSystem(store *database.Store) *ChatSystem {
	once.Do(func() {
		instance = &ChatSystem{
			clients: make(map[uint]chan string),
			store:   store,
		}
	})

	return instance
}

func (c *ChatSystem) Register(userID uint) chan string {
	channel := make(chan string)
	c.clients[userID] = channel
	return channel
}

func (c *ChatSystem) Logout(userID uint) {
	if c.IsUserExist(userID) {
		delete(c.clients, userID)
	}
}

func (c *ChatSystem) SendMessage(senderID uint, msg string) {
	var msgReq dto.RealTimeMessageRequest
	if err := json.Unmarshal([]byte(msg), &msgReq); err != nil {
		logger.Error("Failed Unmarshal json", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "invalid message")
		return
	}

	msgModel := dto.ToMessageModel(senderID, msgReq)
	if msgModel.ReceiverID == senderID {
		logger.Error("cannot send message to yourself")
		c.sendMessage(EventError, senderID, "cannot send message to yourself")
		return
	}

	if err := c.store.DB.Create(&msgModel).Error; err != nil {
		logger.Error("failed inserting message to database", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "internal error")
		return
	}

	json, err := json.Marshal(dto.ToRealTimeMessageResponse(msgModel))
	if err != nil {
		logger.Error("failed Marshal realtime message response to json", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "internal error")
		return
	}

	c.sendMessage(EventMessage, msgModel.ReceiverID, string(json))
}

func (c *ChatSystem) sendMessage(event EventType, receiverID uint, msg string) {
	msg = fmt.Sprintf("%s %s", event, msg)
	if c.IsUserExist(receiverID) {
		c.clients[receiverID] <- msg
	}
}

func (c *ChatSystem) IsUserExist(userID uint) bool {
	_, ok := c.clients[userID]
	return ok
}
