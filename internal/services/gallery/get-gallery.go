package gallery

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Question: Does it required to be shown to the customer?
// or its for photograoher as well?
func (h *Handler) HandleGetAllGallery(c *fiber.Ctx) error {
	var galleries []model.Gallery
	if err := h.store.DB.
		Preload("Photographer.User").
		Preload("Tags").
		Preload("Media").
		Preload("Reviews.Customer").
		Preload("Categories").
		Preload("Quotations.Customer").
		Find(&galleries).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("No galleries found", err)
		}
		return errors.Wrap(err, "Error retrieving galleries")
	}

	var galleryResponses []dto.GalleryResponse
	for _, gallery := range galleries {
		galleryResponses = append(galleryResponses, dto.ToGalleryResponse(gallery))
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result: galleryResponses,
	})
}
