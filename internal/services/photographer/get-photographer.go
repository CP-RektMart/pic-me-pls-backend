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

	if err := h.store.DB.Find(&photographers).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("No photographers found", err)
		}
		return errors.Wrap(err, "Error retrieving photographers")
	}

	var photograperResponses []dto.PhotographerResponse
	for _, photographer := range photographers {
		photograperResponses = append(photograperResponses, dto.ToPhotographerResponse(photographer))
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[[]dto.PhotographerResponse]{
		Result: photograperResponses,
	})
}
