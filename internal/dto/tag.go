package dto

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/samber/lo"
)

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
	return lo.Map(tags, func(tag model.Tag, _ int) TagResponse {
		return ToTagResponse(tag)
	})
}
