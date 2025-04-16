package admin

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) HandleGetAllUsers(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "ok1",
	})
}
