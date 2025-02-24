package packages

import (
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
// @Param        page      query    int    	    false  "Page number"
// @Param        pageSize  query    int    		false  "Page size"
// @Param        minPrice  query    float64    	false  "Minimum price"
// @Param        maxPrice  query    float64    	false  "Maximum price"
// @Param        photographerId  query    uint    false  "Photographer ID"
// @Success      200    {object}  dto.PaginationResponse[dto.PackageResponse]
// @Failure      400    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandleGetAllPackages(c *fiber.Ctx) error {

	req := new(dto.GetAllPackagesRequest)
	req.Pagination = new(dto.PaginationRequest)

	if err := c.QueryParser(req); err != nil {
		return apperror.BadRequest("Invalid query parameters", err)
	}

	if err := c.QueryParser(req.Pagination); err != nil {
		return apperror.BadRequest("Invalid query parameters", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	page, pageSize, offset := checkPaginationRequest(req)
	query, totalCount, err := filterPrice(h, req)
	if err != nil {
		return errors.Wrap(err, "Error filtering packages")
	}

	totalPage := (int(totalCount) + pageSize - 1) / pageSize
	var packages []model.Package
	if err := query.
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
			return apperror.NotFound("No packages found", err)
		}
		return errors.Wrap(err, "Error retrieving packages")
	}

	var PackageResponses []dto.PackageResponse
	for _, Package := range packages {
		PackageResponses = append(PackageResponses, dto.ToPackageResponse(Package))
	}

	return c.Status(fiber.StatusOK).JSON(dto.PaginationResponse[[]dto.PackageResponse]{
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
		Data:      PackageResponses,
	})
}

func checkPaginationRequest(req *dto.GetAllPackagesRequest) (int, int, int) {
	page, pageSize := req.Pagination.Page, req.Pagination.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return page, pageSize, offset
}

func filterPrice(h *Handler, req *dto.GetAllPackagesRequest) (*gorm.DB, int64, error) {
	query := h.store.DB.Model(&model.Package{})
	var totalCount int64
	if req.MinPrice > 0 {
		query = query.Where("price >= ?", req.MinPrice)
	}

	if req.MaxPrice > 0 {
		query = query.Where("price <= ?", req.MaxPrice)
	}

	if req.PhotographerID != nil {
		query = query.Where("photographer_id = ?", *req.PhotographerID)
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Wrap(err, "Error counting packages")
	}

	if err := h.store.DB.Model(&model.Package{}).Count(&totalCount).Error; err != nil {
		return nil, 0, errors.Wrap(err, "Error counting packages")
	}

	return query, totalCount, nil
}
