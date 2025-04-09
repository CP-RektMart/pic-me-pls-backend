package mock

import "time"
import _ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"

type ReportDetail struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	CustomerId     uint      `json:"customer_id"`
	PhotographerId uint      `json:"photographer_id"`
	CreatedAt      time.Time `json:"created_at"`
	QuotationId    uint      `json:"quotation_id"`
	Status         string    `json:"status"`
}

type CreateReportRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	QuotationId uint   `json:"quotation_id"`
}

// @Summary			Get a specific report by ID
// @Description			Get a specific report by ID
// @Tags			report
// @Router			/api/reports/{reportId} [GET]
// @Success			200	{object}	dto.HttpResponse[ReportDetail]
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func GetReportById() {}
// @Summary			Get All own reports
// @Description			Get All own reports
// @Tags			report
// @Router			/api/reports [GET]
// @Security			ApiKeyAuth
// @Success			200	{object}	dto.HttpListResponse[ReportDetail]
// @Failure			400	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func GetAllOwnReports() {}

// @Summary			Create a new report under a quotation
// @Description			Create a new report under a quotation
// @Tags			report
// @Router			/api/quotations/{quotationId}/reports [POST]
// @Security			ApiKeyAuth
// @Param        		RequestBody 	body  CreateReportRequest  true  "Report details"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func CreateReport() {}

// @Summary			Update an existing own report
// @Description			Update an existing own report
// @Tags			report
// @Router			/api/reports/{reportId} [PATCH]
// @Security			ApiKeyAuth
// @Param        		RequestBody 	body  CreateReportRequest  true  "Report details"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func UpdateReport() {}

// @Summary			Delete own report
// @Description			Delete own report
// @Tags			report
// @Router			/api/reports/{reportId} [DELETE]
// @Security			ApiKeyAuth
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func DeleteReport() {}

// @Summary			Get all reports for a specific quotation
// @Description			Get all reports for a specific quotation
// @Tags			report
// @Router			/api/admin/quotations/{quotationId}/reports [GET]
// @Security			ApiKeyAuth
// @Success			200 	{object}	dto.HttpListResponse[ReportDetail]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func GetAllReportsByQuotationId() {}

// @Summary			Update the status of a report
// @Description			(e.g., pending → resolved)
// @Tags			report
// @Router			/api/admin/reports/{reportId}/status [PATCH]
// @Security			ApiKeyAuth
// @Success			204	{object}	dto.HttpResponse[string]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func UpdateReportStatus() {}

// @Summary			Get all users
// @Description			Get all users, ordered by the number of their pending reports with their quotationId that is reported
// @Tags			report
// @Router			/api/admin/users/reports/pending [GET]
// @Security			ApiKeyAuth
// @Success			200	{object}	dto.HttpListResponse[dto.CustomerPublicResponse]	
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func GetAllUsersWithPendingReports() {}

// @Summary			Ban a user by ID
// @Description			Ban a user by ID
// @Tags			report
// @Router			/api/admin/users/{userId}/ban [PATCH]
// @Security			ApiKeyAuth
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func BanUser() {}