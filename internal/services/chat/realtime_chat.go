package chat

import (
	"sync"

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

	ch := h.chatSystem.Register(jwtEntity.ID)

	var wg sync.WaitGroup
	wg.Add(2)

	go h.receiveRealtimeMessage(&wg, c, jwtEntity.ID)
	go h.sendRealtimeMessage(&wg, c, ch)

	wg.Wait()
	h.chatSystem.Logout(jwtEntity.ID)
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
			h.chatSystem.SendMessage(userID, string(msg))
		}
	}
}

func (h *Handler) sendRealtimeMessage(wg *sync.WaitGroup, c *websocket.Conn, ch chan string) {
	defer wg.Done()

	for {
		msg := <-ch

		if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			logger.Error("failed sending message", err)
			logger.Info("closing connection...")
			break
		}
	}
}
