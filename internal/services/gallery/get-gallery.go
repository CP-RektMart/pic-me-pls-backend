package gallery

import (
	"strconv"

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

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return apperror.BadRequest("Invalid page number", err)
	}

	limit, err := strconv.Atoi(c.Query("limit", "20"))
	if err != nil {
		return apperror.BadRequest("Invalid page limit number", err)
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var total int64
	if err := h.store.DB.Model(&model.Gallery{}).Count(&total).Error; err != nil {
		return errors.Wrap(err, "Error counting galleries")
	}

	if err := h.store.DB.
		Preload("Photographer.User").
		Preload("Tags").
		Preload("Media").
		Preload("Reviews.Customer").
		Preload("Categories").
		Preload("Quotations.Customer").
		Limit(limit).
		Offset(offset).
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

	// Pagination response
	pagination := dto.PaginationResponse{
		Page:        page,
		Total:       total,
		Limit:       limit,
		TotalPages:  int((total + int64(limit) - 1) / int64(limit)),
		HasNextPage: int64(offset+limit) < total,
		HasPrevPage: page > 1,
	}

	result := map[string]interface{}{
		"response":   galleryResponses,
		"pagination": pagination,
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result: result,
	})
}
