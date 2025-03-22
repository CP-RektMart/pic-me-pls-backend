package dto

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
)

type CreateReviewRequest struct {
	ID      string   `params:"id" validate:"required"`
	Rating  *float64 `json:"rating" validate:"validRating"`
	Comment string   `json:"comment" validate:"required"`
}

type GetReviewsByPackageIDRequest struct {
	PackageID uint `params:"packageId" validate:"required"`
}

type ReviewResponse struct {
	ID       uint             `json:"id"`
	Rating   *float64         `json:"rating"`
	Comment  string           `json:"comment"`
	Customer CustomerResponse `json:"customer"`
}

func validRating(fl validator.FieldLevel) bool {
	ratingPtr := fl.Field().Float()

	rating := ratingPtr

	allowedValues := map[float64]bool{
		0.0: true, 0.5: true, 1.0: true, 1.5: true, 2.0: true,
		2.5: true, 3.0: true, 3.5: true, 4.0: true, 4.5: true, 5.0: true,
	}
	return allowedValues[rating]
}

func RegisterValidations(validate *validator.Validate) {
	if err := validate.RegisterValidation("validRating", validRating); err != nil {
		panic("failed to register validation: " + err.Error())
	}
}

func ToReviewResponse(review model.Review) ReviewResponse {
	return ReviewResponse{
		ID:       review.ID,
		Rating:   &review.Rating,
		Comment:  review.Comment,
		Customer: ToCustomerResponse(review.Customer),
	}
}

func ToReviewResponses(reviews []model.Review) []ReviewResponse {
	return lo.Map(reviews, func(review model.Review, _ int) ReviewResponse {
		return ToReviewResponse(review)
	})

}
