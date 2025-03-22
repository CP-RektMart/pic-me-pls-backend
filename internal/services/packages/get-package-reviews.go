package packages

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) HandleGetPackageReviews(c *fiber.Ctx) error {
	req := new(dto.GetReviewsByPackageIDRequest)
	if err := c.ParamsParser(req); err != nil {
		return apperror.BadRequest("invalid params", err)
	}

	reviews, err := h.getReviewsByPackageByID(req.ID)
	if err != nil {
		return errors.Wrap(err, "failed get reviews")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpListResponse[dto.ReviewResponse]{
		Result: dto.ToReviewResponses(reviews),
	})
}

func (h *Handler) getReviewsByPackageByID(ID uint) ([]model.Review, error) {
	var reviews []model.Review
	if err := h.store.DB.
		Preload("Customer").
		Where("package_id = ?", ID).
		Find(&reviews).Error; err != nil {
		return []model.Review{}, errors.Wrap(err, "failed to get reviews")
	}

	return reviews, nil
}
