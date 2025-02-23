package photographer

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func (h *Handler) HandleGetAllPhotographers(ctx context.Context, req *dto.HumaBody[dto.PhotographerRequest]) (*dto.HumaHttpResponse[dto.PaginationResponse[[]dto.PhotographerResponse]], error) {
	var photographers []model.Photographer

	// Validate query parameters
	if err := h.validate.Struct(req); err != nil {
		return nil, huma.Error400BadRequest("invalid request", err)
	}

	params := req.Body

	// Set default values
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 5
	}

	offset := (params.Page - 1) * params.PageSize

	// Query photographers
	query := h.store.DB.
		Joins("User").
		Where("\"User\".name ILIKE ?", "%"+params.Name+"%").
		Where("is_verified = ? AND active_status = ?", true, true)

	// Calculate total page
	var totalCount int64
	if err := query.Model(&model.Photographer{}).Count(&totalCount).Error; err != nil {
		return nil, errors.Wrap(err, "Error counting photographers")
	}
	totalPage := (int(totalCount) + params.PageSize - 1) / params.PageSize

	// Retrieve photographers
	if err := query.Limit(params.PageSize).Offset(offset).Find(&photographers).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, huma.Error404NotFound("Photographers not found", err)
		}
		return nil, errors.Wrap(err, "Error retrieving photographers")
	}

	// Convert to response
	photographerResponses := make([]dto.PhotographerResponse, 0, len(photographers))
	for _, photographer := range photographers {
		photographerResponses = append(photographerResponses, dto.ToPhotographerResponse(photographer))
	}

	result := dto.PaginationResponse[[]dto.PhotographerResponse]{
		Page:      params.Page,
		PageSize:  params.PageSize,
		TotalPage: totalPage,
		Data:      photographerResponses,
	}

	return &dto.HumaHttpResponse[dto.PaginationResponse[[]dto.PhotographerResponse]]{
		Body: dto.HttpResponse[dto.PaginationResponse[[]dto.PhotographerResponse]]{
			Result: result,
		},
	}, nil
}
