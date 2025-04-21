package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type CreateReportRequest struct {
	QuotationID uint   `json:"quotationId" validate:"required"`
	Message     string `json:"message" validate:"required"`
	Title       string `json:"title" validate:"required"`
}

func ToReportResponse(report model.Report) ReportResponse {
	return ReportResponse{
		ID:          report.ID,
		QuotationID: report.QuotationID,
		ReporterID:  report.ReporterID,
		Status:      string(report.Status),
		Message:     report.Message,
		Title:       report.Title,
		CreatedAt:   report.CreatedAt.Format("02 Jan 2006, 15:04"),
		UpdatedAt:   report.UpdatedAt.Format("02 Jan 2006, 15:04"),
	}
}

func ToReportResponses(reports []model.Report) []ReportResponse {
	var reportResponses []ReportResponse
	for _, report := range reports {
		reportResponses = append(reportResponses, ToReportResponse(report))
	}
	return reportResponses
}

type GetReportByIDRequest struct {
	ReportID uint `params:"id" validate:"required"`
}

type ReportResponse struct {
	ID          uint   `json:"id"`
	QuotationID uint   `json:"quotationId"`
	ReporterID  uint   `json:"reporterId"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	Title       string `json:"title"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type ReportListResponse struct {
	Reports []ReportResponse `json:"reports"`
}

func ToGetReportByIDResponse(report model.Report) ReportResponse {
	return ReportResponse{
		ID:          report.ID,
		QuotationID: report.QuotationID,
		ReporterID:  report.ReporterID,
		Status:      string(report.Status),
		Message:     report.Message,
		Title:       report.Title,
		CreatedAt:   report.CreatedAt.Format("02 Jan 2006, 15:04"),
		UpdatedAt:   report.UpdatedAt.Format("02 Jan 2006, 15:04"),
	}
}

type UpdateReportRequest struct {
	ReportID uint   `params:"id" validate:"required"`
	Message  string `json:"message"`
	Title    string `json:"title"`
}

type AdminGetAllReportsRequest struct {
	PaginationRequest
	Title string `query:"title" default:""`
}
