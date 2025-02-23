package packages

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func (h *Handler) HandleGetAllPackages(ctx context.Context, req *dto.HumaBody[dto.PaginationRequest]) (*dto.HumaHttpResponse[dto.PaginationResponse[[]dto.PackageResponse]], error) {

	if err := h.validate.Struct(req); err != nil {
		return nil, huma.Error400BadRequest("invalid request", err)
	}

	page, pageSize, offset := checkPaginationRequest(&req.Body)

	var packages []model.Package
	var totalCount int64
	if err := h.store.DB.Model(&model.Package{}).Count(&totalCount).Error; err != nil {
		return nil, errors.Wrap(err, "Error counting packages")
	}

	totalPage := (int(totalCount) + pageSize - 1) / pageSize

	if err := h.store.DB.
		Preload("Photographer.User").
		Preload("Tags").
		Preload("Media").
		Preload("Reviews.Customer").
		Preload("Categories").
		Preload("Quotations.Customer").
		Limit(pageSize).
		Offset(offset).
		Find(&packages).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, huma.Error404NotFound("Packages not found", err)
		}
		return nil, errors.Wrap(err, "Error retrieving packages")
	}

	var PackageResponses []dto.PackageResponse
	for _, Package := range packages {
		PackageResponses = append(PackageResponses, dto.ToPackageResponse(Package))
	}

	result := dto.PaginationResponse[[]dto.PackageResponse]{
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
		Data:      PackageResponses,
	}

	return &dto.HumaHttpResponse[dto.PaginationResponse[[]dto.PackageResponse]]{
		Body: dto.HttpResponse[dto.PaginationResponse[[]dto.PackageResponse]]{
			Result: result,
		},
	}, nil
}

func checkPaginationRequest(req *dto.PaginationRequest) (int, int, int) {
	page, pageSize := req.Page, req.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return page, pageSize, offset
}
