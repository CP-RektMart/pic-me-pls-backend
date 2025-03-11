package quotation

import (
	"fmt"

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
// @Param       body  body      dto.CreateQuotationRequest       true  "Quotation update details"
// @Success     200
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

	body := new(dto.CreateQuotationRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	fmt.Println(body.CustomerID)
	fmt.Println(body.PackageID)

	if err := h.validate.Struct(body); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.UpdateQuotation(body, userID, req.QuotationID); err != nil {
		return errors.Wrap(err, "failed to create quotation")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) UpdateQuotation(req *dto.CreateQuotationRequest, userID uint, quotationID string) (error) {

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var quotation  model.Quotation
		if  err := tx.First(&quotation, quotationID).Error; err != nil {
			return apperror.NotFound("quotation not found", err)
		}

		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userID).Error; err != nil {
			return apperror.NotFound( "Photographer not found for user", err)
		}

		if quotation.PhotographerID != photographer.ID {
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

		if err := tx.Model(&quotation).Where("id = ?", quotationID).Updates(model.Quotation{
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