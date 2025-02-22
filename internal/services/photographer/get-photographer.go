package photographer

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Get All Photographers
// @Description		Retrieve a paginated list of photographers, optionally filtered by name.
// @Tags			photographer
// @Router			/api/v1/photographers [GET]
// @Param			page		query		int	false	"Page number for pagination (default: 1)"
// @Param			pageSize	query		int	false	"Number of records per page (default: 5, max: 20)"
// @Param			name		query		string	false	"Filter by photographer's name (case-insensitive)"
// @Success      	200 {object} 	dto.PaginationResponse[dto.PhotographerResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetAllPhotographers(c *fiber.Ctx) error {
	var photographers []model.Photographer
	var params dto.PhotographerRequest

	// Parse query parameters
	if err := c.QueryParser(&params); err != nil {
		return apperror.BadRequest("Invalid query parameters", err)
	}

	// Set default values
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 5
	}

	offset := (params.Page - 1) * params.PageSize

	// Query photographers
	query := h.store.DB.Preload("User").
		Joins("User").
		Where("\"User\".name ILIKE ?", "%"+params.Name+"%").
		Where("is_verified = ? AND active_status = ?", true, true)

	// Calculate total page
	var totalCount int64
	if err := query.Model(&model.Photographer{}).Count(&totalCount).Error; err != nil {
		return errors.Wrap(err, "Error counting photographers")
	}
	totalPage := (int(totalCount) + params.PageSize - 1) / params.PageSize

	// Retrieve photographers
	if err := query.Limit(params.PageSize).Offset(offset).Find(&photographers).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("No photographers found", err)
		}
		return errors.Wrap(err, "Error retrieving photographers")
	}

	// Convert to response
	photographerResponses := make([]dto.PhotographerResponse, 0, len(photographers))
	for _, photographer := range photographers {
		photographerResponses = append(photographerResponses, dto.ToPhotographerResponse(photographer))
	}

	return c.Status(fiber.StatusOK).JSON(dto.PaginationResponse[dto.PhotographerResponse]{
		Page:      params.Page,
		PageSize:  params.PageSize,
		TotalPage: totalPage,
		Data:      photographerResponses,
	})
}
