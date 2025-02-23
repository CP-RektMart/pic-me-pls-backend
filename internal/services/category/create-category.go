package category

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) HandleCreate(c *fiber.Ctx) error {
	var req dto.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid body", err)
	}

	if err := h.store.DB.Where("name = ?", req.Name).First(&model.Category{}).Error; err == nil {
		return apperror.BadRequest("duplicate category name", nil)
	}

	category := model.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := h.store.DB.Create(&category).Error; err != nil {
		return errors.Wrap(err, "failed create category")
	}

	return c.SendStatus(fiber.StatusCreated)
}
