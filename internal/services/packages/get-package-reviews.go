package packages

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary      get package reviews by package id
// @Description  Show reviews of a package
// @Tags         packages
// @Router       /api/v1/customer/packages/:packageID/reviews [GET]
// @Param        packageID  path    uint     true  "package id"
// @Success      200    {object}  dto.HttpListResponse[dto.ReviewResponse]
// @Failure      400    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandleGetPackageReviews(c *fiber.Ctx) error {
	var req dto.GetReviewsByPackageIDRequest

	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid params", err)
	}

	if err := c.QueryParser(&req.PaginationRequest); err != nil {
		return apperror.BadRequest("invalid pagination query parameters", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid query parameters", err)
	}

	page, pageSize, offset := dto.GetPaginationData(req.PaginationRequest, 1, 5)

	query := h.store.DB.Model(&model.Review{}).Where("package_id = ?", req.PackageID)

	reviews, err := h.executeReviewsQuery(query, pageSize, offset)
	if err != nil {
		return errors.Wrap(err, "failed to fetch reviews from DB")
	}

	totalCount, err := h.countReviews(query)
	if err != nil {
		return errors.Wrap(err, "failed to count reviews")
	}

	totalPages := (totalCount + pageSize - 1) / pageSize

	reviewResponses := dto.ToReviewResponses(reviews)

	return c.Status(fiber.StatusOK).JSON(dto.PaginationResponse[dto.ReviewResponse]{
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPages,
		Data:      reviewResponses,
	})
}

func (h *Handler) executeReviewsQuery(query *gorm.DB, limit, offset int) ([]model.Review, error) {
	var reviews []model.Review
	if err := query.
		Preload("Customer").
		Limit(limit).
		Offset(offset).
		Find(&reviews).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NotFound("No reviews found", err)
		}
		return nil, errors.Wrap(err, "error retrieving reviews")
	}
	return reviews, nil
}

func (h *Handler) countReviews(query *gorm.DB) (int, error) {
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return 0, errors.Wrap(err, "failed to count reviews")
	}
	return int(totalCount), nil
}
