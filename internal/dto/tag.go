package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToTagResponse(tag model.Tag) TagResponse {
	return TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

func ToTagResponses(tags []model.Tag) []TagResponse {
	responses := make([]TagResponse, 0)
	for _, tag := range tags {
		responses = append(responses, ToTagResponse(tag))
	}
	return responses
}
