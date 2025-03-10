package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type MediaPackageRequest struct {
	PictureURL  string `json:"pictureUrl" validate:"required"`
	Description string `json:"description"`
}

type GetAllPackagesRequest struct {
	Pagination     *PaginationRequest
	MinPrice       float64 `query:"minPrice" validate:"omitempty,min=0"`
	MaxPrice       float64 `query:"maxPrice" validate:"omitempty,min=0"`
	PhotographerID *uint   `query:"photographerId" validate:"omitempty"`
}

type CreatePackageRequest struct {
	Name        string                `json:"name" validate:"required"`
	Description string                `json:"description" validate:"required"`
	Price       float64               `json:"price" validate:"required"`
	Media       []MediaPackageRequest `json:"media" validate:"required"`
}

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
}

type CreatePackageResponse struct {
	ID               uint    `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	PhotographerID   uint    `json:"photographerId"`
	PhotographerName string  `json:"photographerName"`
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
	}
}

type UpdatePackageRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description" validate:"min=0"`
	Price       float64 `json:"price" validate:"omitempty,min=0"`
	PackageID   uint    `params:"packageId"`
}

type UpdatePackageResponse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
