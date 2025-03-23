package quotation

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			submit preview photo
// @Description		submit preview photo and set status to submitted
// @Tags			quotations
// @Router			/api/v1/photographer/quotations/{id}/preview [PATCH]
// @Security		ApiKeyAuth
// @Param 			quotation id 	path 	uint 	true 	"quotaion id"
// @Param 			link 		body 	string 	true 	"link"
// @Success			204
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleCreatePreviewPhoto(c *fiber.Ctx) error {
	var req dto.CreatePreviewPhotoRequest
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	quotation, err := h.CreatePreviewPhoto(&req)
	if err != nil {
		errors.Wrap(err, "failed to create preview photo")
	}

	if err := h.SetStatusSubmitted(quotation); err != nil {
		errors.Wrap(err, "failed to set status to submitted")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) CreatePreviewPhoto(req *dto.CreatePreviewPhotoRequest) (*model.Quotation, error) {
	var quotation model.Quotation
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		link := req.Link
		quotationID := req.QuotationID

		preview := model.Preview{
			Link:        link,
			QuotationID: quotationID,
		}

		if err := tx.Where("id = ?", quotationID).First(&quotation).Error; err != nil {
			return apperror.NotFound("quotation not found", err)
		}

		if err := tx.Save(&preview).Error; err != nil {
			return errors.Wrap(err, "failed to save preview")

		}

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "failed to create preview photo")
	}

	return &quotation, nil
}

func (h *Handler) SetStatusSubmitted(quotation *model.Quotation) error {
	quotation.Status = model.QuotationPreviewPhotoSubmitted
	if err := h.store.DB.Save(&quotation).Error; err != nil {
		return errors.Wrap(err, "failed to set status to submitted")
	}

	return nil
}
