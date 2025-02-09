package server

import "github.com/gofiber/fiber/v2"

func (s *Server) handleGetMe(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "นี่คือเสียงจากเด็กวัด"})
}
