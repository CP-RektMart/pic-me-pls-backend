package dto

type CreateReportRequest struct {
	QuotationID  uint   `json:"quotationId" validate:"required"`
	ReporterRole string `json:"reporterRole" validate:"required,oneof=CUSTOMER PHOTOGRAPHER"`
	Message      string `json:"message" validate:"required"`
	Title        string `json:"title" validate:"required"`
}
