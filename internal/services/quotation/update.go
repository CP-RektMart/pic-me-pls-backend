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
// @Param       id    path      uint                 		true  "Quotation ID"
// @Param       body  body      dto.UpdateQuotationRequest true  "Quotation update details"
// @Success     200
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

	req := new(dto.UpdateQuotationRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.UpdateQuotation(req, userID, uint(quotationID)); err != nil {
		return errors.Wrap(err, "failed to create quotation")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) UpdateQuotation(req *dto.UpdateQuotationRequest, userID uint, quotationID uint) (error) {

	var status model.QuotationStatus = model.QuotationStatus(req.Status)
	if (!status.IsValid()) {
		return errors.Wrap(errors.Errorf("Invalid"), "Status is invalid")
	}

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var quotation  model.Quotation
		if  err := tx.First(&quotation, quotationID).Error; err != nil {
			return errors.Wrap(err, "quotation not found")
		}

		// Find the photographer associated with the user
		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userID).Error; err != nil {
			return errors.Wrap(err, "Photographer not found for user")
		}

		// Check the Photographer is owner 
		if quotation.PhotographerID != photographer.ID {
			return apperror.Forbidden("you do not have permission to update this quotation", errors.Errorf("Not Permission"))
		}

		quotation.CustomerID = req.CustomerID
		quotation.GalleryID = req.GalleryID
		quotation.Description = req.Description
		quotation.Price = req.Price
		quotation.FromDate = req.FromDate
		quotation.ToDate = req.ToDate
		quotation.Status = status
		

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