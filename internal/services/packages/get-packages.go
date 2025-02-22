package packages

import (
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary      Get all packages
// @Description  Show all avaliable packages with pagination
// @Tags         Package
// @Router       /api/v1/package [GET]
// @Param        page   query   int  false  "Page number (default is 1)"
// @Param        limit  query   int  false  "Number of items per page (default is 20)"
// @Success      200    {object}  dto.PackageListHttpResponse
// @Failure      400    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandleGetAllPackages(c *fiber.Ctx) error {
	var packages []model.Package

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return apperror.BadRequest("Invalid page number", err)
	}

	limit, err := strconv.Atoi(c.Query("limit", "20"))
	if err != nil {
		return apperror.BadRequest("Invalid page limit number", err)
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var total int64
	if err := h.store.DB.Model(&model.Package{}).Count(&total).Error; err != nil {
		return errors.Wrap(err, "Error counting packages")
	}

	if err := h.store.DB.
		Preload("Photographer.User").
		Preload("Tags").
		Preload("Media").
		Preload("Reviews.Customer").
		Preload("Categories").
		Preload("Quotations.Customer").
		Limit(limit).
		Offset(offset).
		Find(&packages).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("No packages found", err)
		}
		return errors.Wrap(err, "Error retrieving packages")
	}

	var PackageResponses []dto.PackageResponse
	for _, Package := range packages {
		PackageResponses = append(PackageResponses, dto.ToPackageResponse(Package))
	}

	// Pagination response
	pagination := dto.PaginationResponse[dto.PackageResponse]{
		Page:        page,
		Total:       total,
		Limit:       limit,
		TotalPages:  int((total + int64(limit) - 1) / int64(limit)),
		HasNextPage: int64(offset+limit) < total,
		HasPrevPage: page > 1,
		Response:    PackageResponses,
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.PaginationResponse[dto.PackageResponse]]{
		Result: pagination,
	})

}
