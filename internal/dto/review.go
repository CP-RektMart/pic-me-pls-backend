package dto

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/samber/lo"
)

type ReviewResponse struct {
	ID       uint             `json:"id"`
	Rating   float64          `json:"rating"`
	Comment  string           `json:"comment"`
	Customer CustomerResponse `json:"customer"`
}

func ToReviewResponse(review model.Review) ReviewResponse {
	return ReviewResponse{
		ID:       review.ID,
		Rating:   review.Rating,
		Comment:  review.Comment,
		Customer: ToCustomerResponse(review.Customer),
	}
}

func ToReviewResponses(reviews []model.Review) []ReviewResponse {
	return lo.Map(reviews, func(review model.Review, _ int) ReviewResponse {
		return ToReviewResponse(review)
	})

}
