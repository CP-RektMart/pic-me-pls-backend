package photographer

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *Handler) HandleGetAllPhotographer(c *fiber.Ctx) error {
	var photographers []model.Photographer

	page := c.QueryInt("page", 1)
	limit := 5
	offset := (page - 1) * limit

	if err := h.store.DB.Limit(limit).Offset(offset).Find(&photographers).Error; err != nil {
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
	h.store.DB.Model(&model.Photographer{}).Count(&totalCount)

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[[]dto.PhotographerResponse]{
		Result: photographerResponses,
	})
}
