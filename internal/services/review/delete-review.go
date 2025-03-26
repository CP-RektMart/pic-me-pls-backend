package review

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary 		Delete a review
// @Description 	Delete a review from a quotation.
// @Tags 			reviews
// @Router 			/api/v1/customer/quotations/{quotationId}/review/{id} [DELETE]
// @Security    	ApiKeyAuth
// @Param 			quotationId 	path 	string 	true 	"Quotation ID"
// @Param 			id 				path 	uint 	true 	"ID"
// @Success 		204 			"Review Deleted successfully"
// @Failure     	400 {object}  	dto.HttpError
// @Failure     	500 {object} 	dto.HttpError
func (h *Handler) HandleDeleteReview(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	var req dto.DeleteReviewRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid params", err)
	}

	if err := h.DeleteReview(&req, userID); err != nil {
		return errors.Wrap(err, "failed to Delete review")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) DeleteReview(req *dto.DeleteReviewRequest, userID uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var customer model.User
		if err := h.store.DB.Where("id = ?", userID).First(&customer).Error; err != nil {
			return apperror.NotFound("Customer not found", err)
		}

		var quotation model.Quotation
		if err := h.store.DB.Where("id = ?", req.QuotationID).First(&quotation).Error; err != nil {
			return apperror.NotFound("Quotaion not found", err)
		}

		var review model.Review
		if err := h.store.DB.Where("id = ?", req.ID).First(&review).Error; err != nil {
			return apperror.NotFound("Review not found", err)
		}

		if review.CustomerID != userID {
			return apperror.Forbidden("You are not allowed to delete this review", errors.New("unauthorized"))
		}

		if err := h.store.DB.Delete(&review).Error; err != nil {
			return errors.Wrap(err, "failed to delete review")
		}

		return nil

	}); err != nil {
		return errors.Wrap(err, "failed to delete review")
	}
	return nil
}
