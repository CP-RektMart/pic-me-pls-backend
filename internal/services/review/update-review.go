package review

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary 		Update a review
// @Description 	Update a review for a quotation.
// @Tags 			reviews
// @Router 			/api/v1/customer/quotations/{quotationId}/review/{id} [PATCH]
// @Security    	ApiKeyAuth
// @Param 			quotationId path string true "Quotation ID"
// @Param 			id 			path uint true "ID"
// @Param 			review 		body 		dto.UpdateReviewRequest true "Review details"
// @Success 		204 			"Review Updated successfully"
// @Failure     	400   {object}  dto.HttpError
// @Failure     	500   {object}  dto.HttpError
func (h *Handler) HandleUpdateReview(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	var req dto.UpdateReviewRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid params", err)
	}

	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid body 1", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid body 2", err)
	}

	if err := h.UpdateReview(&req, userID, req.ID); err != nil {
		return errors.Wrap(err, "failed to Update review")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) UpdateReview(req *dto.UpdateReviewRequest, userID uint, reviewID uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var customer model.User
		if err := h.store.DB.Where("id = ?", userID).First(&customer).Error; err != nil {
			return apperror.NotFound("Customer not found", err)
		}

		var quotation model.Quotation
		if err := h.store.DB.Where("id = ?", req.QuotationID).First(&quotation).Error; err != nil {
			return apperror.NotFound("Quotation not found", err)
		}

		var targetReview model.Review
		if err := h.store.DB.Where("id = ?", req.ID).First(&targetReview).Error; err != nil {
			return apperror.NotFound("Review not found", err)
		}

		if targetReview.CustomerID != userID {
			return apperror.Forbidden("You are not allowed to update this review", errors.New("unauthorized"))
		}

		if err := h.store.DB.Model(&targetReview).Updates(model.Review{
			PackageID:   quotation.PackageID,
			CustomerID:  customer.ID,
			QuotationID: quotation.ID,
			Rating:      *req.Rating,
			Comment:     req.Comment,
			IsEdited:    true,
		}).Error; err != nil {
			return errors.Wrap(err, "Failed to update review")
		}

		return nil

	}); err != nil {
		return errors.Wrap(err, "failed to update review")
	}
	return nil
}
