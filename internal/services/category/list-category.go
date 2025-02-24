package category

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			list category
// @Description			list category
// @Tags			category
// @Router			/api/v1/categories [GET]
// @Param			page		query		int	false	"Page number for pagination (default: 1)"
// @Param			pageSize	query		int	false	"Number of records per page (default: 20)"
// @Success			200	{object}	dto.HttpResponse[PaginationResponse[dto.CategoryResponse]]
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleListCategory(c *fiber.Ctx) error {
	var req dto.PaginationRequest
	categories := make([]model.Category, 0)

	if err := c.QueryParser(&req); err != nil {
		return apperror.BadRequest("invalid query", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid query", err)
	}

	page, pageSize, offset := dto.GetPaginationData(req, 1, 20)
	if err := h.store.DB.Offset(offset).Limit(pageSize).Find(&categories).Error; err != nil {
		return errors.Wrap(err, "failed list categories")
	}

	var count int64
	if err := h.store.DB.Model(&model.Category{}).Count(&count).Error; err != nil {
		return errors.Wrap(err, "failed count categories")
	}
	totalPage := (int(count) + pageSize - 1) / pageSize

	result := dto.PaginationResponse[dto.CategoryResponse]{
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
		Data:      dto.ToCategoryResponses(categories),
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.PaginationResponse[dto.CategoryResponse]]{
		Result: result,
	})
}
