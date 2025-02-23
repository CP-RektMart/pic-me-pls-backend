package category

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			update category
// @Description			update category
// @Tags			category
// @Router			/api/v1/categories/{id} [PATCH]
// @Security			ApiKeyAuth
// @Param 			id	 	path 	uint			 	true 	"category id"
// @Param 			RequestBody 	body 	dto.UpdateCategoryRequest 	true 	"request body (don't need to include id)"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdateCategory(c *fiber.Ctx) error {
	var req dto.UpdateCategoryRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid params", err)
	}

	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid body", err)
	}

	category := model.Category{
		Model:       gorm.Model{ID: req.ID},
		Name:        req.Name,
		Description: req.Description,
	}
	if err := h.store.DB.Updates(&category).Error; err != nil {
		return errors.Wrap(err, "failed create category")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
