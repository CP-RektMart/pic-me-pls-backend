package dto

type CreateReportRequest struct {
	QuotationID uint   `json:"quotationId" validate:"required"`
	Message     string `json:"message" validate:"required"`
	Title       string `json:"title" validate:"required"`
}

type UpdateReportRequest struct {
	ReportID uint   `params:"id" validate:"required"`
	Message  string `json:"message"`
	Title    string `json:"title"`
}
