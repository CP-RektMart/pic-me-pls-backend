package message

import (
	"github.com/cockroachdb/errors"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

const jwtEntityKey = "jwtEntityKey"

// @Summary      connect to websocket
// @Description  Establish a WebSocket connection for real-time communication
// @Tags         message
// @Router       /api/v1/messages/ws [GET]
// @Security	 ApiKeyAuth
// @Success      101    "Switching Protocols"
// @Failure      400
func (h *Handler) HandleWebsocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		jwtEntity, err := h.authentication.GetJWTEntityFromContext(c.UserContext())
		if err != nil {
			return errors.Wrap(err, "failed getting jwtEntity from context")
		}
		c.Locals(jwtEntityKey, jwtEntity)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
