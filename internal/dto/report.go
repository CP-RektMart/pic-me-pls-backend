package dto

type CreateReportRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	QuotationID uint   `json:"quotationID"`
}

type ReportResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	QuotationID uint   `json:"quotationId"`
}

type UpdateReportRequest struct {
	ID          uint   `params:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
