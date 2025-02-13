package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type MediaResponse struct {
	ID         uint   `json:"id"`
	PictureURL string `json:"pictureUrl"`
}

func ToMediaResponses(media []model.Media) []MediaResponse {
	var responses []MediaResponse
	for _, m := range media {
		responses = append(responses, MediaResponse{ID: m.ID, PictureURL: m.PictureURL})
	}
	return responses
}
