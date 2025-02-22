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
// @Tags         Packages
// @Router       /api/v1/packages [GET]
// @Param        page      query    int    false  "Page number"
// @Param        page_size query    int    false  "Page size"
// @Success      200    {object}  dto.HttpResponse[dto.PackageListResponse]
// @Failure      400    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandleGetAllPackages(c *fiber.Ctx) error {

	req := new(dto.GetAllPackagesRequest)
	if err := c.QueryParser(req); err != nil {
		return apperror.BadRequest("Invalid query parameters", err)
	}

	page, pageSize, offset := checkGetAllPackagesRequest(req)

	var packages []model.Package
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

	pagination := dto.PaginationResponse[dto.PackageResponse]{
		Page:       page,
		Total:      total,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
		Response:   PackageResponses,
	}

	result := dto.PackageListResponse{
		Pagination: pagination,
		Packages:   PackageResponses,
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.PackageListResponse]{
		Result: result,
	})

}

func checkGetAllPackagesRequest(req *dto.GetAllPackagesRequest) (int, int, int) {
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
