package server

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) handleGetMe(c *fiber.Ctx) error {
	userDto, ok := c.Locals("user").(*model.UserDto)

	if !ok {
		return apperror.BadRequest("no user profile found in context", nil)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"profile": userDto,
	})
}
