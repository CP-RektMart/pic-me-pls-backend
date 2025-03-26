package message

import "github.com/gofiber/fiber/v2"

func (h *Handler) HandleSupportWebAPI(c *fiber.Ctx) error {
	header := c.Query("accessToken")
	if header != "" {
		c.Request().Header.Set("Authorization", "Bearer "+header)
	}
	return c.Next()
}
