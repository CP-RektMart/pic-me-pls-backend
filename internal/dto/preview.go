package dto

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/samber/lo"
)

type CreatePreviewPhotoRequest struct {
	Link        string `json:"link" validate:"required"`
	QuotationID uint   `json:"quotationId" validate:"required"`
}

type ListPreviewResponse struct {
	ID          uint   `json:"id"`
	Link        string `json:"link"`
	QuotationID uint   `json:"quotationId"`
}

func ToListPreviewResponse(preview model.Preview) ListPreviewResponse {
	return ListPreviewResponse{
		ID:          preview.ID,
		Link:        preview.Link,
		QuotationID: preview.QuotationID,
	}
}

func ToListPreviewResponses(previews []model.Preview) []ListPreviewResponse {
	return lo.Map(previews, func(preview model.Preview, _ int) ListPreviewResponse {
		return ToListPreviewResponse(preview)
	})
}
