package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToCategoryResponses(categories []model.Category) []CategoryResponse {
	var responses []CategoryResponse
	for _, category := range categories {
		responses = append(responses, CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return responses
}

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateCategoryRequest struct {
	ID          uint   `params:"id" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
