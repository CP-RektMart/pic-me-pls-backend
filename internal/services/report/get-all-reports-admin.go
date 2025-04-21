package report

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary     Admin get all reports
// @Description Get all reports in the system
// @Tags		admin
// @Router      /api/v1/admin/reports [GET]
// @Security			ApiKeyAuth
// @Param        		page      query    int    false  "Page number"
// @Param        		pageSize query    int    false  "Page size"
// @Success     200 	{object}  dto.PaginationResponse[dto.ReportResponse]
// @Failure     400   	{object}  dto.HttpError
// @Failure     401   	{object}  dto.HttpError
// @Failure     403   	{object}  dto.HttpError
// @Failure     404   	{object}  dto.HttpError
// @Failure     500   	{object}  dto.HttpError
func (h *Handler) HandleAdminGetAllReports(c *fiber.Ctx) error {

	var req dto.AdminGetAllReportsRequest

	if err := c.QueryParser(&req); err != nil {
		return apperror.BadRequest("Invalid query", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("Invalid parameter", err)
	}

	var reports *dto.PaginationResponse[dto.ReportResponse]

	reports, err := h.adminGetAllReports(req)
	if err != nil {
		return errors.Wrap(err, "Failed getting reports")
	}

	return c.Status(fiber.StatusOK).JSON(reports)
}

func (h *Handler) adminGetAllReports(req dto.AdminGetAllReportsRequest) (*dto.PaginationResponse[dto.ReportResponse], error) {
	page, pageSize, offset := dto.GetPaginationData(req.PaginationRequest, 1, 20)

	var reports []model.Report
	db := h.store.DB.Model(&model.Report{}).Where("title ILIKE ?", "%"+req.Title+"%")

	if err := db.Offset(offset).Limit(pageSize).Find(&reports).Error; err != nil {
		return nil, errors.Wrap(err, "Failed getting reports")
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "Failed counting reports")
	}

	totalPage := (int(count) + pageSize - 1) / pageSize

	paginationResponse := dto.PaginationResponse[dto.ReportResponse]{
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
		Data:      dto.ToReportResponses(reports),
	}

	return &paginationResponse, nil
}
