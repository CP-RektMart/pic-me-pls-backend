package category

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) HandleDeleteCategory(c *fiber.Ctx) error {
	var req dto.DeleteCategoryRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid params", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid body", err)
	}

	if err := h.store.DB.Delete(&model.Category{}, req.ID).Error; err != nil {
		return errors.Wrap(err, "failed delete category")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
