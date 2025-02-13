package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToTagResponses(tags []model.Tag) []TagResponse {
	var responses []TagResponse
	for _, tag := range tags {
		responses = append(responses, TagResponse{ID: tag.ID, Name: tag.Name})
	}
	return responses
}
