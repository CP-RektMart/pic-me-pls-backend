package chat

import (
	"encoding/json"
	"sync"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/contrib/websocket"
)

func (h *Handler) HandleRealTimeChat(c *websocket.Conn) {
	jwtEntity, ok := c.Locals(jwtEntityKey).(jwt.JWTentity)
	if !ok {
		logger.Error("failed receive userID from jwtEntity")
		c.Close()
	}

	ch := h.chatSystem.Register(jwtEntity.ID)

	var wg sync.WaitGroup
	wg.Add(2)

	go h.receiveRealtimeMessage(&wg, c, jwtEntity.ID)
	go h.sendRealtimeMessage(&wg, ch, c)

	wg.Wait()

	h.chatSystem.Logout(jwtEntity.ID)
	c.Close()
}

func (h *Handler) receiveRealtimeMessage(wg *sync.WaitGroup, c *websocket.Conn, userID uint) {
	defer wg.Done()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			logger.Error("failed receiving message", err)
			logger.Info("closing connection...")
			break
		}

		if mt == websocket.TextMessage {
			var message dto.RealTimeMessageRequest
			if err := json.Unmarshal(msg, &message); err != nil {
				logger.Error("invalid message", err)
				if err := h.sendErrorMessage(c, "invalid message"); err != nil {
					logger.Error("failed sending error", err)
				}
				continue
			}
			if err := h.validate.Struct(message); err != nil {
				logger.Error("invalid message", err)
				if err := h.sendErrorMessage(c, "invalid message"); err != nil {
					logger.Error("failed sending error", err)
				}
				continue
			}
			if userID == message.ReceiverID {
				logger.Error("receiverID is same as userID")
				if err := h.sendErrorMessage(c, "receiverID is same as userID"); err != nil {
					logger.Error("failed sending error", err)
				}
				continue
			}

			if err := h.chatSystem.SendMessage(dto.ToMessageModel(userID, message)); err != nil {
				logger.Error("failed sending message", err)
				if err := h.sendErrorMessage(c, "internal error"); err != nil {
					logger.Error("failed sending error", err)
				}
				continue
			}
		}
	}
}

func (h *Handler) sendRealtimeMessage(wg *sync.WaitGroup, ch chan model.Message, c *websocket.Conn) {
	defer wg.Done()

	for {
		msg := <-ch

		json, err := json.Marshal(dto.ToRealTimeMessageResponse(msg))
		if err != nil {
			logger.Error("failed convert message to json")
			if err := h.sendErrorMessage(c, "internal error"); err != nil {
				logger.Error("failed sending error", err)
			}
		}

		if err := c.WriteMessage(websocket.TextMessage, json); err != nil {
			logger.Error("failed sending message", err)
			logger.Info("closing connection...")
			break
		}
	}
}

func (h *Handler) sendErrorMessage(c *websocket.Conn, msg string) error {
	response := dto.ToRealTimeMessageResponse(
		model.Message{
			Type:    model.MessageTypeError,
			Content: msg,
		},
	)

	json, err := json.Marshal(&response)
	if err != nil {
		return errors.Wrap(err, "failed marshal json")
	}

	if err := c.WriteMessage(websocket.TextMessage, json); err != nil {
		return errors.Wrap(err, "failed sending error message")
	}

	return nil
}
