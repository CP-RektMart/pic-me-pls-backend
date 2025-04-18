package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type CreateReportRequest struct {
	QuotationID uint   `json:"quotationId" validate:"required"`
	Message     string `json:"message" validate:"required"`
	Title       string `json:"title" validate:"required"`
}

type GetReportByIDRequest struct {
	ReportID uint `params:"id" validate:"required"`
}

type ReportResponse struct {
	ID          uint   `json:"reportId"`
	QuotationID uint   `json:"quotationId"`
	ReporterID  uint   `json:"reporterId"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	Title       string `json:"title"`
}

type ReportListResponse struct {
	Reports []ReportResponse `json:"reports"`
}

func ToReportResponses(reports []model.Report) []ReportResponse {
	var reportResponses []ReportResponse
	for _, report := range reports {
		reportResponses = append(reportResponses, ReportResponse{
			QuotationID: report.QuotationID,
			ReporterID:  report.ReporterID,
			Status:      string(report.Status),
			Message:     report.Message,
			Title:       report.Title,
		})
	}
	return reportResponses
}
