package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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

type DeleteCategoryRequest struct {
	ID uint `params:"id" validate:"required"`
}

func ToCategoryResponse(category model.Category) CategoryResponse {
	return CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func ToCategoryResponses(categories []model.Category) []CategoryResponse {
	responses := make([]CategoryResponse, 0)
	for _, category := range categories {
		responses = append(responses, ToCategoryResponse(category))
	}
	return responses
}
