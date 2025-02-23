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

type CreateMediaRequest struct {
	PictureURL  string `json:"pictureUrl" validate:"required"`
	Description string `json:"description"`
	PackageID   uint   `json:"packageId" validate:"required,min=1"`
}

type UpdateMediaRequest struct {
	PictureURL  string `json:"pictureUrl"`
	Description string `json:"description"`
	MediaId     uint   `params:"mediaId" validate:"required,min=1"`
}

type UpdateMediaResponse struct {
	PictureURL  string `json:"pictureUrl"`
	Description string `json:"description"`
}

type DeleteMediaRequest struct {
	MediaId uint `params:"mediaId"`
}
