package review

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary 	Create a review
// @Description Create a review for a quotation.
// @Tags 		reviews
// @Router 		/api/v1/customer/quotations/{quotationId}/review [POST]
// @Security    ApiKeyAuth
// @Param 		quotationId 	path 		uint 					true "Quotation ID"
// @Param 		review 			body 		dto.CreateReviewRequest true "Review details"
// @Success 	204 			"Review created successfully"
// @Failure     400   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
func (h *Handler) HandleCreateReview(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	var req dto.CreateReviewRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid params", err)
	}

	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid body 1", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid body 2", err)
	}

	if err := h.createReview(&req, userID); err != nil {
		return errors.Wrap(err, "failed to create review")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) createReview(req *dto.CreateReviewRequest, userID uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var customer model.User
		if err := h.store.DB.Where("id = ?", userID).First(&customer).Error; err != nil {
			return apperror.NotFound("Customer not found", err)
		}

		var quotation model.Quotation
		if err := h.store.DB.Where("id = ?", req.QuotationID).First(&quotation).Error; err != nil {
			return apperror.NotFound("Quotation not found", err)
		}

		review := model.Review{
			PackageID:   quotation.PackageID,
			CustomerID:  customer.ID,
			QuotationID: quotation.ID,
			Rating:      *req.Rating,
			Comment:     req.Comment,
		}

		if err := h.store.DB.Create(&review).Error; err != nil {
			return errors.Wrap(err, "failed to save Package to DB")
		}

		return nil

	}); err != nil {
		return errors.Wrap(err, "failed to create review")
	}
	return nil
}
