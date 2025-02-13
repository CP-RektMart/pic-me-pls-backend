package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToCategoryResponses(categories []model.Category) []CategoryResponse {
	var responses []CategoryResponse
	for _, category := range categories {
		responses = append(responses, CategoryResponse{ID: category.ID, Name: category.Name})
	}
	return responses
}
