package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type MediaResponse struct {
	ID          uint   `json:"id"`
	PictureURL  string `json:"pictureUrl"`
	Description string `json:"description"`
}

type CreateMediaRequest struct {
	PictureURL  string `json:"pictureUrl" validate:"required"`
	Description string `json:"description"`
	PackageID   uint   `json:"packageId" validate:"required,min=1"`
}

type UpdateMediaRequest struct {
	PictureURL  string `json:"pictureUrl"`
	Description string `json:"description"`
	MediaID     uint   `params:"mediaId" validate:"required,min=1"`
}

type UpdateMediaResponse struct {
	PictureURL  string `json:"pictureUrl"`
	Description string `json:"description"`
}

type DeleteMediaRequest struct {
	MediaID uint `params:"mediaId"`
}

func ToMediaResponse(media model.Media) MediaResponse {
	return MediaResponse{
		ID:          media.ID,
		PictureURL:  media.PictureURL,
		Description: media.Description,
	}
}

func ToMediaResponses(media []model.Media) []MediaResponse {
	responses := make([]MediaResponse, 0)
	for _, m := range media {
		responses = append(responses, ToMediaResponse(m))
	}
	return responses
}
