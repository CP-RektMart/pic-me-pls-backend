package quotation

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			cancelled quotation
// @Description			cancelled quotation
// @Tags			quotations
// @Router			/api/v1/customer/quotations/{id}/cancel [PATCH]
// @Security			ApiKeyAuth
// @Param 			quotation id 	path 	uint 	true 	"quotation id"
// @Success			204
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandlerCancelQuotation(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed get user id from context")
	}

	var req dto.CancelQuotationRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid params", err)
	}

	var quotation *model.Quotation
	if err := h.store.DB.First(&quotation, req.QuotationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("quotation not found", err)
		}
		return errors.Wrap(err, "failed find quotation")
	}

	if quotation.CustomerID != userID {
		return apperror.Forbidden("user not have permission", nil)
	}

	quotation.Status = model.QuotationCancelled
	if err := h.store.DB.Save(&quotation).Error; err != nil {
		return errors.Wrap(err, "failed confirm quotation")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
