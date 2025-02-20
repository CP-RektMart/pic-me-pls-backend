package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type GalleryResponse struct {
	ID           uint                 `json:"id"`
	Name         string               `json:"name"`
	Description  string               `json:"description,omitempty"`
	Price        float64              `json:"price"`
	Photographer PhotographerResponse `json:"photographer"`
	Tags         []TagResponse        `json:"tags,omitempty"`
	Media        []MediaResponse      `json:"media,omitempty"`
	Reviews      []ReviewResponse     `json:"reviews,omitempty"`
	Categories   []CategoryResponse   `json:"categories,omitempty"`
	Quotations   []QuotationResponse  `json:"quotations,omitempty"`
}

type PaginationResponse struct {
	Page        int   `json:"page"`
	Total       int64 `json:"total"`
	Limit       int   `json:"limit"`
	TotalPages  int   `json:"total_pages"`
	HasNextPage bool  `json:"has_next_page"`
	HasPrevPage bool  `json:"has_prev_page"`
}

func ToGalleryResponse(gallery model.Gallery) GalleryResponse {
	return GalleryResponse{
		ID:           gallery.ID,
		Name:         gallery.Name,
		Description:  gallery.Description,
		Price:        gallery.Price,
		Photographer: ToPhotographerResponse(gallery.Photographer),
		Tags:         ToTagResponses(gallery.Tags),
		Media:        ToMediaResponses(gallery.Media),
		Reviews:      ToReviewResponses(gallery.Reviews),
		Categories:   ToCategoryResponses(gallery.Categories),
		Quotations:   ToQuotationResponses(gallery.Quotations),
	}
}
