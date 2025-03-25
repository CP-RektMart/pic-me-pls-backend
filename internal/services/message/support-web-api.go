package message

import "github.com/gofiber/fiber/v2"

func (h *Handler) HandleSupportWebAPI(c *fiber.Ctx) error {
	header := c.Get("Sec-WebSocket-Protocol")
	if header != "" {
		c.Request().Header.Set("Authorization", "Bearer "+header)
	}
	println(c.Get("Authorization"))
	return c.Next()
}
