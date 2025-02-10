package user

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store    *database.Store
	validate *validator.Validate
}

func NewHandler(store *database.Store, validate *validator.Validate) *Handler {
	return &Handler{
		store:    store,
		validate: validate,
	}
}

func (h *Handler) HandleGetMe(c *fiber.Ctx) error {
	userDto, ok := c.Locals("user").(*model.UserDto)

	if !ok {
		return apperror.BadRequest("no user profile found in context", nil)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"profile": userDto,
	})
}
