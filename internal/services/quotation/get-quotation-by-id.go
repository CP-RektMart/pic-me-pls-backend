package quotation

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Get Quotation By ID
// @Description		Get Quotation By ID
// @Tags			quotations
// @Router			/api/v1/quotations/{id} [GET]
// @Security		ApiKeyAuth
// @Param 			id 	path 	uint 	true 	"quotaion id"
// @Success			200	{object}	dto.HttpResponse[dto.QuotationResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetQuotationByID(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.GetQuotationRequest)
	if err := c.ParamsParser(req); err != nil {
		return apperror.BadRequest("invalid params", err)
	}

	var quotation *model.Quotation
	if err := h.store.DB.
		Preload("Package.Photographer").
		Preload("Package.Media").
		Preload("Customer").
		Preload("Photographer.User").
		First(&quotation, req.QuotationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("quotation not found", err)
		}
		return errors.Wrap(err, "failed find quotation")
	}

	if quotation.CustomerID != userID && quotation.Photographer.UserID != userID {
		return apperror.Forbidden("user not have permission", nil)
	}

	var packages []model.Package

	photographerID := quotation.PhotographerID
	if err := h.store.DB.Where("photographer_id = ?", photographerID).Find(&packages).Error; err != nil {
		return apperror.NotFound("package not found", err)
	}

	quotation.Photographer.Packages = packages

	response := dto.ToQuotationResponse(*quotation)

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.QuotationResponse]{
		Result: response,
	})

}
