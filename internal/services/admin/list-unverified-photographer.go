package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/utils/pagination"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			list unverified citizencard
// @Tags			admin
// @Router			/api/v1/admin/citizenCards/unverify [GET]
// @Security		ApiKeyAuth
// @Param			page		query		int	false	"Page number for pagination (default: 1)"
// @Param			pageSize	query		int	false	"Number of records per page (default: 5, max: 20)"
// @Param			name	query		string	false	"Filter by photographer's name (case-insensitive)"
// @Success			200	{object}	dto.HttpResponse[dto.PaginationResponse[dto.ListUnverifiedPhotographerResponse]]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleListUnverifiedPhotographer(c *fiber.Ctx) error {
	var req dto.ListUnverifiedPhotographerRequest
	if err := c.QueryParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	page, pageSize, offset := req.PaginationRequest.GetPaginationData(1, 10)
	photographers, err := h.listUnverifiedPhotographer(req.Name, offset, pageSize)
	if err != nil {
		return errors.Wrap(err, "failed query photographers")
	}

	count, err := h.countUnverifiedPhotographers(req.Name)
	if err != nil {
		return errors.Wrap(err, "failed count table photographers")
	}

	totalPage := pagination.TotalPageFromCount(count, pageSize)
	photosResp := dto.ToListUnverifiedPhotographersResponse(photographers)
	paginationResp := dto.NewPaginationResponse(photosResp, page, pageSize, totalPage)

	return c.Status(fiber.StatusOK).JSON(dto.Success(paginationResp))
}

func (h *Handler) listUnverifiedPhotographer(name *string, offset, limit int) ([]model.Photographer, error) {
	var ps []model.Photographer

	db := h.store.DB

	if name != nil {
		db = db.Where("\"User\".name ILIKE ?", "%"+*name+"%")
	}

	if err := db.Joins("User").Joins("CitizenCard").Where("is_verified = ?", false).Offset(offset).Limit(limit).Find(&ps).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NotFound("photographers not found", err)
		}
		return nil, err
	}

	return ps, nil
}

func (h *Handler) countUnverifiedPhotographers(name *string) (int, error) {
	var c int64

	db := h.store.DB

	if name != nil {
		db = db.Where("\"User\".name ILIKE ?", "%"+*name+"%")
	}

	if err := db.Joins("User").Where("is_verified = ?", false).Model(&model.Photographer{}).Count(&c).Error; err != nil {
		return 0, err
	}

	return int(c), nil
}
