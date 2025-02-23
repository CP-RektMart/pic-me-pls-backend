package quotation

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary     Update a quotation
// @Description Updates an existing quotation
// @Tags        quotation
// @Router      /api/v1/quotations/{id} [PATCH]
// @Security    ApiKeyAuth
// @Param       id    path      uint                 true  "Quotation ID"
// @Param       body  body      dto.QuotationRequest true  "Quotation update details"
// @Success     200   {object}  dto.HttpResponse
// @Failure     400   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
func (h *Handler) HandleUpdate(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	// Use param to query
	quotationID, err := c.ParamsInt("id")
	if err != nil {
		return apperror.BadRequest("invalid quotation ID", err)
	}

	req := new(dto.QuotationRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.UpdateQuotation(req, userID, uint(quotationID)); err != nil {
		return errors.Wrap(err, "failed to create quotation")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) UpdateQuotation(req *dto.QuotationRequest, userID uint, quotationID uint) (error) {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		quotation := &model.Quotation{
			PhotographerID: userID,
			CustomerID: req.CustomerID,
			GalleryID: req.GalleryID,
			Description:    req.Description,
			Price:          req.Price,
			FromDate: req.FromDate,
			ToDate:  req.ToDate,
		}

		// Find the quotation
		if err := tx.First(&quotation, quotationID).Error; err != nil {
			return errors.Wrap(err, "quotation not found")
		}

		// Check the Photographer is owner 
		if quotation.PhotographerID != userID {
			return apperror.Forbidden("you do not have permission to update this quotation", errors.Errorf("Not Permission"))
		}

		// Check CustomerID and GalleryID existed in database
		var customer model.User
		if err := tx.First(&customer, req.CustomerID).Error; err != nil {
			return errors.Wrap(err, "customer not found")
		}
		var gallery model.Gallery
		if err := tx.First(&gallery, req.GalleryID).Error; err != nil {
			return errors.Wrap(err, "gallery not found")
		}

		// Save changes
		if err := tx.Save(&quotation).Error; err != nil {
			return errors.Wrap(err, "failed to update quotation")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to create quotation")
	}
	return nil
} 