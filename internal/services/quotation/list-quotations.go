package quotation

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			list quotations
// @Description			list quotations
// @Tags			quotations
// @Router			/api/v1/quotations [GET]
// @Security			ApiKeyAuth
// @Param        		page      query    int    false  "Page number"
// @Param        		page_size query    int    false  "Page size"
// @Success			200 	{object} 	dto.PaginationResponse[QuotationResponse]
// @Success			401 	{object} 	dto.PaginationResponse[QuotationResponse]
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleListQuotations(c *fiber.Ctx) error {
	jwtEntity, err := h.authMiddleware.GetJWTEntityFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed getting jwtEntity from context")
	}

	var req dto.PaginationRequest
	if err := c.QueryParser(&req); err != nil {
		return apperror.BadRequest("invalid query", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid parameter", err)
	}

	var quotations *dto.PaginationResponse[dto.QuotationResponse]

	switch jwtEntity.Role {
	case model.UserRoleCustomer:
		quotations, err = h.listUserQuotations(req, jwtEntity.ID)
	case model.UserRolePhotographer:
		quotations, err = h.listPhotographerQuotations(req, jwtEntity.ID)
	default:
		return apperror.Forbidden("invalid role to get quotations", err)
	}

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.PaginationResponse[dto.QuotationResponse]{
		Page:      quotations.Page,
		PageSize:  quotations.PageSize,
		TotalPage: quotations.TotalPage,
		Data:      quotations.Data,
	})
}

func (h *Handler) listUserQuotations(req dto.PaginationRequest, userID uint) (*dto.PaginationResponse[dto.QuotationResponse], error) {
	page, pageSize, offset := dto.GetPaginationData(req, 1, 20)
	var quotations []model.Quotation
	if err := h.store.DB.
		Offset(offset).
		Limit(pageSize).
		Preload("Package").
		Preload("Package.Tags").
		Preload("Package.Media").
		Preload("Customer").
		Preload("Photographer.User").
		Where("customer_id = ?", userID).Find(&quotations).Error; err != nil {
		return nil, errors.Wrap(err, "failed getting quotations")
	}

	var count int64
	if err := h.store.DB.Model(&model.Quotation{}).Where("id = ?", userID).Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "failed counting quotaions")
	}
	totalPage := (int(count) + pageSize - 1) / pageSize

	paginationResponse := dto.PaginationResponse[dto.QuotationResponse]{
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
		Data:      dto.ToQuotationResponses(quotations),
	}

	return &paginationResponse, nil
}

func (h *Handler) listPhotographerQuotations(req dto.PaginationRequest, userID uint) (*dto.PaginationResponse[dto.QuotationResponse], error) {
	var photographer model.Photographer
	if err := h.store.DB.Where("user_id = ?", userID).First(&photographer).Error; err != nil {
		return nil, errors.Wrap(err, "failed getting photographer data")
	}

	page, pageSize, offset := dto.GetPaginationData(req, 1, 20)
	var quotations []model.Quotation
	if err := h.store.DB.
		Offset(offset).
		Limit(pageSize).
		Preload("Package").
		Preload("Package.Tags").
		Preload("Package.Media").
		Preload("Customer").
		Preload("Photographer.User").
		Where("photographer_id = ?", photographer.UserID).Find(&quotations).Error; err != nil {
		return nil, errors.Wrap(err, "failed getting quotations")
	}

	var count int64
	if err := h.store.DB.Model(&model.Quotation{}).Where("photographer_id = ?", photographer.UserID).Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "failed counting quotaions")
	}
	totalPage := (int(count) + pageSize - 1) / pageSize

	paginationResponse := dto.PaginationResponse[dto.QuotationResponse]{
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
		Data:      dto.ToQuotationResponses(quotations),
	}

	return &paginationResponse, nil
}
