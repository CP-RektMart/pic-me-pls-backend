package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type ReviewResponse struct {
	ID       uint    `json:"id"`
	Rating   float64 `json:"rating"`
	Comment  string  `json:"comment,omitempty"`
	Customer string  `json:"customer"`
}

func ToReviewResponses(reviews []model.Review) []ReviewResponse {
	var responses []ReviewResponse
	for _, review := range reviews {
		responses = append(responses, ReviewResponse{
			ID:       review.ID,
			Rating:   review.Rating,
			Comment:  review.Comment,
			Customer: review.Customer.Name,
		})
	}
	return responses
}
