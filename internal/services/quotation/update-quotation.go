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
// @Tags        quotations
// @Router      /api/v1/photographer/quotations/{id} [PATCH]
// @Security    ApiKeyAuth
// @Param       id    path      uint                 		true  "Quotation ID"
// @Param       body  body      dto.UpdateQuotationRequest       true  "Quotation update details"
// @Success     204
// @Failure     400   {object}  dto.HttpError
// @Failure     403   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
func (h *Handler) HandleUpdateQuotation(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.UpdateQuotationRequest)
	if err := c.ParamsParser(req); err != nil {
		return apperror.BadRequest("invalid params", err)
	}

	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.UpdateQuotation(req, userID); err != nil {
		return errors.Wrap(err, "failed to create quotation")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) UpdateQuotation(req *dto.UpdateQuotationRequest, userID uint) (error) {

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var quotation  model.Quotation
		if  err := tx.First(&quotation, req.QuotationID).Error; err != nil {
			return apperror.NotFound("quotation not found", err)
		}

		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userID).Error; err != nil {
			return apperror.NotFound( "Photographer not found for user", err)
		}

		if quotation.PhotographerID != photographer.UserID {
			return apperror.Forbidden("you do not have permission to update this quotation", errors.Errorf("Not Permission"))
		}

		var customer model.User
		if err := tx.First(&customer, req.CustomerID).Error; err != nil {
			return apperror.NotFound("customer not found", err)
		}

		var targetPackage model.Package
		if err := tx.First(&targetPackage, req.PackageID).Error; err != nil {
			return apperror.NotFound("package not found", err)
		}

		if err := tx.Model(&quotation).Where("id = ?", req.QuotationID).Updates(model.Quotation{
			CustomerID:   req.CustomerID,
			PackageID:    req.PackageID,
			Description:  req.Description,
			Price:        req.Price,
			FromDate:     req.FromDate,
			ToDate:       req.ToDate,
		}).Error; err != nil {
			return errors.Wrap(err, "failed to update quotation")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to update quotation")
	}
	return nil
} 