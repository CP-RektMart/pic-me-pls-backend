package dto

type QuotationResponse struct {
	ID       uint   `json:"id"`
	Status   string `json:"status"`
	Customer string `json:"customer"`
}
