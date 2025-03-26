package chat

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
)

type EventType string

const (
	EventError   EventType = "ERROR"
	EventMessage EventType = "MESSAGE"
)

type Client struct {
	Message   chan string
	Terminate chan bool
}

type Server struct {
	clients  map[uint]*Client
	store    *database.Store
	validate *validator.Validate
}

var (
	instance *Server
	once     sync.Once
)

func NewServer(store *database.Store, validate *validator.Validate) *Server {
	once.Do(func() {
		instance = &Server{
			store:    store,
			validate: validate,
			clients:  make(map[uint]*Client),
		}
	})

	return instance
}

func (c *Server) Register(userID uint) *Client {
	client := Client{
		Message:   make(chan string),
		Terminate: make(chan bool),
	}
	c.clients[userID] = &client
	return &client
}

func (c *Server) Logout(userID uint) {
	if c.isUserExist(userID) {
		c.clients[userID].Terminate <- true
		delete(c.clients, userID)
	}
}

func (c *Server) SendRawString(senderID uint, msg string) {
	var msgReq dto.RealTimeMessageRequest
	if err := json.Unmarshal([]byte(msg), &msgReq); err != nil {
		logger.Error("Failed Unmarshal json", slog.Any("error", err))
		c.sendMessage(EventError, senderID, "invalid message")
		return
	}

	if err := c.validate.Struct(msgReq); err != nil {
		logger.Error("Failed validate message request", slog.Any("error", err))
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

func (c *Server) SendMessageModel(msg model.Message) error {
	if err := c.store.DB.Create(&msg).Error; err != nil {
		return errors.Wrap(err, "failed save massage record to database")
	}

	json, err := json.Marshal(dto.ToRealTimeMessageResponse(msg))
	if err != nil {
		return errors.Wrap(err, "failed Marshal to json")
	}

	c.sendMessage(EventMessage, msg.ReceiverID, string(json))

	return nil
}

func (c *Server) sendMessage(event EventType, receiverID uint, msg string) {
	msg = fmt.Sprintf("%s %s", event, msg)
	if c.isUserExist(receiverID) {
		c.clients[receiverID].Message <- msg
	}
}

func (c *Server) isUserExist(userID uint) bool {
	_, ok := c.clients[userID]
	return ok
}
