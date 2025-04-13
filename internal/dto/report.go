package dto

type CreateReportRequest struct {
	QuotationID  uint   `json:"quotation_id" validate:"required"`
	ReporterRole string `json:"reporter_role" validate:"required,oneof=CUSTOMER PHOTOGRAPHER"`
	Message      string `json:"message" validate:"required"`
}
