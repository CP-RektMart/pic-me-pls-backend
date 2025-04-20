package report

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary     Get all reports
// @Description Get all reports of a user
// @Tags		customer
// @Router      /api/v1/customer/reports [GET]
// @Security			ApiKeyAuth
// @Param        		page      query    int    false  "Page number"
// @Param        		pageSize query    int    false  "Page size"
// @Success     200 	{object}  dto.PaginationResponse[dto.ReportResponse]
// @Failure     400   	{object}  dto.HttpError
// @Failure     401   	{object}  dto.HttpError
// @Failure     403   	{object}  dto.HttpError
// @Failure     404   	{object}  dto.HttpError
// @Failure     500   	{object}  dto.HttpError
func (h *Handler) HandleGetAllReports(c *fiber.Ctx) error {

	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "Failed to get user id from context")
	}

	var req dto.PaginationRequest
	if err := c.QueryParser(&req); err != nil {
		return apperror.BadRequest("Invalid query", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("Invalid parameter", err)
	}

	var reports *dto.PaginationResponse[dto.ReportResponse]

	reports, err = h.getAllReports(req, userID)
	if err != nil {
		return errors.Wrap(err, "Failed getting reports")
	}

	return c.Status(fiber.StatusOK).JSON(reports)
}

func (h *Handler) getAllReports(req dto.PaginationRequest, userID uint) (*dto.PaginationResponse[dto.ReportResponse], error) {

	page, pageSize, offset := dto.GetPaginationData(req, 1, 20)
	var reports []model.Report
	if err := h.store.DB.
		Offset(offset).
		Limit(pageSize).
		Where("reporter_id = ?", userID).Find(&reports).Error; err != nil {
		return nil, errors.Wrap(err, "Failed getting reports")
	}

	var count int64
	if err := h.store.DB.Model(&model.Report{}).Where("reporter_id = ?", userID).Count(&count).Error; err != nil {
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
