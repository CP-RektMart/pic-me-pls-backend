package quotation

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *Handler) HandleCreatePreview(c *fiber.Ctx) error {
	var req dto.CreatePreviewRequest
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.CreatePreviewPhoto(req); err != nil {
		return errors.Wrap(err, "failed to create preview photo")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) CreatePreviewPhoto(req dto.CreatePreviewRequest) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		link := req.Link
		quotationID := req.QuotationID

		preview := model.Preview{
			Link:        link,
			QuotationID: quotationID,
		}

		var quotation model.Quotation
		if err := tx.Where("id = ?", quotationID).First(&quotation).Error; err != nil {
			return apperror.NotFound("quotation not found", err)
		}

		if err := tx.Save(&preview).Error; err != nil {
			return errors.Wrap(err, "failed to save preview")

		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to create preview photo")
	}

	return nil
}
