package chat

import (
	"sync"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
)

type ChatSystem struct {
	clients map[uint]chan model.Message
	store   *database.Store
}

var (
	instance *ChatSystem
	once     sync.Once
)

func NewChatSystem(store *database.Store) *ChatSystem {
	once.Do(func() {
		instance = &ChatSystem{
			clients: make(map[uint]chan model.Message),
			store:   store,
		}
	})

	return instance
}

func (c *ChatSystem) Register(userID uint) chan model.Message {
	channel := make(chan model.Message)
	c.clients[userID] = channel
	return channel
}

func (c *ChatSystem) Logout(userID uint) {
	if c.IsUserExist(userID) {
		delete(c.clients, userID)
	}
}

func (c *ChatSystem) SendMessage(message model.Message) {
	receiverID := message.ReceiverID

	if c.IsUserExist(receiverID) {
		c.clients[receiverID] <- message
	}
}

func (c *ChatSystem) IsUserExist(userID uint) bool {
	_, ok := c.clients[userID]
	return ok
}
