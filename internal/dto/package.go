package dto

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/samber/lo"
)

type MediaPackageRequest struct {
	PictureURL  string `json:"pictureUrl" validate:"required"`
	Description string `json:"description"`
}

type GetAllPackagesRequest struct {
	Pagination     PaginationRequest
	MinPrice       float64 `query:"minPrice" validate:"omitempty,min=0"`
	MaxPrice       float64 `query:"maxPrice" validate:"omitempty,min=0"`
	PhotographerID *uint   `query:"photographerId" validate:"omitempty"`
	CategoryIDs    string  `query:"categoryIds"`
}

type CreatePackageRequest struct {
	Name        string                `json:"name" validate:"required"`
	Description string                `json:"description"`
	Price       float64               `json:"price" validate:"required,min=0"`
	CategoryID  uint                  `json:"categoryId"`
	Media       []MediaPackageRequest `json:"media" validate:"required"`
}

type PackageResponse struct {
	ID           uint                 `json:"id"`
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	Price        float64              `json:"price"`
	Photographer PhotographerResponse `json:"photographer"`
	Tags         []TagResponse        `json:"tags"`
	Media        []MediaResponse      `json:"media"`
	Reviews      []ReviewResponse     `json:"reviews"`
	Category     *CategoryResponse    `json:"category"`
}

type SmallPackageResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreatePackageResponse struct {
	ID               uint    `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	PhotographerID   uint    `json:"photographerId"`
	PhotographerName string  `json:"photographerName"`
}

type UpdatePackageRequest struct {
	ID          uint    `params:"id" validate:"required"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"omitempty,min=0"`
	CategoryID  uint    `json:"categoryId"`
}

type UpdatePackageResponse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
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
		Category:     lo.EmptyableToPtr(ToCategoryResponse(Package.Category)),
	}
}

func ToPackageResponses(packages []model.Package) []PackageResponse {
	return lo.Map(packages, func(pkg model.Package, _ int) PackageResponse {
		return ToPackageResponse(pkg)
	})
}

func ToPackageMediaModel(media MediaPackageRequest) model.Media {
	return model.Media{
		PictureURL:  media.PictureURL,
		Description: media.Description,
	}
}

func ToPackageMediaModels(media []MediaPackageRequest) []model.Media {
	return lo.Map(media, func(m MediaPackageRequest, _ int) model.Media {
		return ToPackageMediaModel(m)
	})
}
