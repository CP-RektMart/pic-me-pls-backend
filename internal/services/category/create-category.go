package category

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			create category
// @Description			create category
// @Tags			categories
// @Router			/api/v1/admin/categories [POST]
// @Security			ApiKeyAuth
// @Param 			RequestBody 	body 	dto.CreateCategoryRequest 	true 	"request body"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleCreateCategory(c *fiber.Ctx) error {
	var req dto.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid body", err)
	}

	if err := h.validate.Struct(req); err != nil {
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

	return c.SendStatus(fiber.StatusNoContent)
}

func ValidateCreateCategoryRequest(req dto.CreateCategoryRequest) error {
	if req.Name == "" {
		return apperror.BadRequest("name is required", nil)
	}
	if req.Description == "" {
		return apperror.BadRequest("description is required", nil)
	}
	return nil
}
