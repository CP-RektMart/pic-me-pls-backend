package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type QuotationResponse struct {
	ID       uint    `json:"id"`
	Status   string  `json:"status"`
	Customer string  `json:"customer"`
	Price    float64 `json:"price"`
}

func ToQuotationResponses(quotations []model.Quotation) []QuotationResponse {
	var responses []QuotationResponse
	for _, quotation := range quotations {
		responses = append(responses, QuotationResponse{
			ID:       quotation.ID,
			Status:   quotation.Status,
			Price:    quotation.Price,
			Customer: quotation.Customer.Name,
		})
	}
	return responses
}
