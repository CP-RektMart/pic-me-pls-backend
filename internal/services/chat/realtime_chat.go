package chat

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/chatsystem"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/gofiber/contrib/websocket"
)

func (h *Handler) HandleRealTimeChat(c *websocket.Conn) {
	defer c.Close()

	jwtEntity, ok := c.Locals(jwtEntityKey).(jwt.JWTentity)
	if !ok {
		logger.Error("failed receive userID from jwtEntity")
		return
	}

	client := h.chatSystem.Register(jwtEntity.ID)

	go h.receiveRealtimeMessage(c, jwtEntity.ID)
	go h.sendRealtimeMessage(c, jwtEntity.ID, client)
}

func (h *Handler) receiveRealtimeMessage(c *websocket.Conn, userID uint) {
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			logger.Error("failed receiving message", err)
			logger.Info("closing connection...")
			h.chatSystem.Logout(userID)
			break
		}

		if mt == websocket.TextMessage {
			h.chatSystem.SendMessage(userID, string(msg))
		}
	}
}

func (h *Handler) sendRealtimeMessage(c *websocket.Conn, userID uint, client *chatsystem.Client) {
	select {
	case <-client.Terminate:
		break
	case msg := <-client.Message:
		if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			logger.Error("failed sending message", err)
			logger.Info("closing connection...")
			h.chatSystem.Logout(userID)
			break
		}
	}
}
