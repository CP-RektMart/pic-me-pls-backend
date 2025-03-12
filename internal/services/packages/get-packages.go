package packages

import (
	"strings"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary      Get all packages
// @Description  Show all available packages with pagination
// @Tags         packages
// @Router       /api/v1/packages [GET]
// @Param 	 	 name	  query    string    	false  "Filter by package name"
// @Param        page      query    int    	    false  "Page number"
// @Param        pageSize  query    int    		false  "Page size"
// @Param        minPrice  query    float64    	false  "Minimum price"
// @Param        maxPrice  query    float64    	false  "Maximum price"
// @Param        photographerId  query    uint    false  "Photographer ID"
// @Param        categoryIds  query    string    false  "list of categoryIDs separate by comma ex: 1,2"
// @Success      200    {object}  dto.PaginationResponse[dto.PackageResponse]
// @Failure      400    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandleGetAllPackages(c *fiber.Ctx) error {
	var req dto.GetAllPackagesRequest

	if err := c.QueryParser(&req); err != nil {
		return apperror.BadRequest("Invalid query parameters", err)
	}

	if err := c.QueryParser(&req.Pagination); err != nil {
		return apperror.BadRequest("Invalid pagination query parameters", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("Invalid query parameters", err)
	}

	page, pageSize, offset := req.Pagination.CheckPaginationRequest()

	query := h.createfilterPackageQuery(&req)

	packages, err := h.executefilterPackageQuery(query, pageSize, offset)
	if err != nil {
		return errors.Wrap(err, "failed fetch packages from DB")
	}

	totalCount, err := h.countPackages(query)
	if err != nil {
		return errors.Wrap(err, "failed count packages")
	}
	totalPages := (totalCount + pageSize - 1) / pageSize

	packageResponses := dto.ToPackageResponses(packages)

	return c.Status(fiber.StatusOK).JSON(dto.PaginationResponse[dto.PackageResponse]{
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPages,
		Data:      packageResponses,
	})
}

func (h *Handler) createfilterPackageQuery(req *dto.GetAllPackagesRequest) *gorm.DB {
	query := h.store.DB.Model(&model.Package{})

	if req.PackageName != "" {
		query = query.Where("name ILIKE ?", "%"+req.PackageName+"%")
	}

	if req.MinPrice > 0 {
		query = query.Where("price >= ?", req.MinPrice)
	}

	if req.MaxPrice > 0 {
		query = query.Where("price <= ?", req.MaxPrice)
	}

	if req.PhotographerID != nil {
		query = query.Where("photographer_id = ?", *req.PhotographerID)
	}

	if req.CategoryIDs != "" {
		categoryIDs := strings.Split(req.CategoryIDs, ",")
		query = query.Where("category_id IN ?", categoryIDs)
	}

	return query
}

func (h *Handler) executefilterPackageQuery(query *gorm.DB, limit, offset int) ([]model.Package, error) {
	var packages []model.Package
	if err := query.
		Preload("Category").
		Preload("Photographer.User").
		Preload("Tags").
		Preload("Media").
		Preload("Reviews.Customer").
		Limit(limit).
		Offset(offset).
		Find(&packages).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NotFound("No packages found", err)
		}
		return nil, errors.Wrap(err, "Error retrieving packages")
	}
	return packages, nil
}

func (h *Handler) countPackages(query *gorm.DB) (int, error) {
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return 0, errors.Wrap(err, "Failed count all rows")
	}
	return int(totalCount), nil
}
