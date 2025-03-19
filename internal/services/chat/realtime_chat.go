package chat

import (
	"strconv"
	"strings"
	"sync"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
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

	go h.receiveRealtimeMessage(&wg, c)
	go h.sendRealtimeMessage(&wg, ch, c)

	wg.Wait()
	c.Close()
}

func (h *Handler) receiveRealtimeMessage(wg *sync.WaitGroup, c *websocket.Conn) {
	defer wg.Done()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			logger.Error("failed receiving message: %s", err)
			logger.Info("closing connection...")
			break
		}

		if mt == websocket.TextMessage {
			messages := strings.Split(string(msg), " ")
			recevierID, err := strconv.Atoi(messages[0])
			if err != nil {
				logger.Error("failed convert receiver id", err)
				continue
			}
			h.chatSystem.SendMessage(uint(recevierID), messages[1])
		}
	}
}

func (h *Handler) sendRealtimeMessage(wg *sync.WaitGroup, ch chan string, c *websocket.Conn) {
	defer wg.Done()

	for {
		msg := <-ch
		if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			logger.Error("failed sending message: %s", err)
			logger.Info("closing connection...")
			break
		}
	}
}
