package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type PackageResponse struct {
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

type MediaPackageRequest struct {
	PictureURL  string `json:"pictureUrl" validate:"required"`
	Description string `json:"description"`
}

type CreatePackageRequest struct {
	Name        string                `json:"name" validate:"required"`
	Description string                `json:"description" validate:"required"`
	Price       float64               `json:"price" validate:"required"`
	Media       []MediaPackageRequest `json:"media" validate:"required"`
}

type PaginationResponse struct {
	Page        int   `json:"page"`
	Total       int64 `json:"total"`
	Limit       int   `json:"limit"`
	TotalPages  int   `json:"total_pages"`
	HasNextPage bool  `json:"has_next_page"`
	HasPrevPage bool  `json:"has_prev_page"`
}

type CreatePackageResponse struct {
	ID               uint    `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	PhotographerID   uint    `json:"photographerId"`
	PhotographerName string  `json:"photographerName"`
}

type PackageListResponse struct {
	Pagination PaginationResponse `json:"pagination"`
	Response   []PackageResponse  `json:"response"`
}

func ToPackageResponse(Package model.Package) PackageResponse {
	return PackageResponse{
		ID:           Package.ID,
		Name:         Package.Name,
		Description:  Package.Description,
		Price:        Package.Price,
		Photographer: ToPhotographerResponse(Package.Photographer),
		Tags:         ToTagResponses(Package.Tags),
		Media:        ToMediaResponses(Package.Media),
		Reviews:      ToReviewResponses(Package.Reviews),
		Categories:   ToCategoryResponses(Package.Categories),
		Quotations:   ToQuotationResponses(Package.Quotations),
	}
}
