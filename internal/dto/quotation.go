package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type AcceptQuotationRequest struct {
	QuotationID string `params:"id"`
}

type QuotationResponse struct {
	ID       uint    `json:"id"`
	Status   string  `json:"status"`
	Customer string  `json:"customer"`
	Price    float64 `json:"price"`
}

func ToQuotationResponse(quotation model.Quotation) QuotationResponse {
	return QuotationResponse{
		ID:       quotation.ID,
		Status:   quotation.Status.String(),
		Customer: quotation.Customer.Name,
		Price:    quotation.Price,
	}
}

func ToQuotationResponses(quotations []model.Quotation) []QuotationResponse {
	var responses []QuotationResponse
	for _, quotation := range quotations {
		responses = append(responses, ToQuotationResponse(quotation))
	}
	return responses
}
