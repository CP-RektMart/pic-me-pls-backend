package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (h *Handler) HandleGetAllUsers(c *fiber.Ctx) error {
	var users []model.User

	if err := h.store.DB.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("users not found", err)
		}

		return errors.Wrap(err, "failed fetch users")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[[]dto.PublicUserResponse]{
		Result: dto.ToPublicUserResponses(users),
	})
}
