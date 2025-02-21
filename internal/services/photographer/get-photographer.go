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
// @Success			200	{object}	dto.PaginationResponse
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetAllPhotographer(c *fiber.Ctx) error {
	var photographers []model.Photographer

	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 5)
	nameFilter := c.Query("name", "")
	offset := (page - 1) * pageSize

	if page < 1 || pageSize <= 0 || pageSize > 20 {
		return apperror.BadRequest("Invalid page or pageSize parameter", nil)
	}

	query := h.store.DB.Preload("User").Where("is_verified = ? AND active_status = ?", true, true)

	if nameFilter != "" {
		query = query.Joins("JOIN users ON users.id = photographers.user_id").
			Where("users.name ILIKE ?", "%"+nameFilter+"%")
	}

	if err := query.Limit(pageSize).Offset(offset).Find(&photographers).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("No photographers found", err)
		}
		return errors.Wrap(err, "Error retrieving photographers")
	}

	var photographerResponses []dto.PhotographerResponse

	for _, photographer := range photographers {
		photographerResponses = append(photographerResponses, dto.ToPhotographerResponse(photographer))
	}

	var totalCount int64
	query.Model(&model.Photographer{}).Count(&totalCount)
	totalPage := int((totalCount + int64(pageSize) - 1) / int64(pageSize))

	return c.Status(fiber.StatusOK).JSON(dto.PaginationResponse{
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
		Data:      photographerResponses,
	})
}
