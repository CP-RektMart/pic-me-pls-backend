package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type AcceptQuotationRequest struct {
	QuotationID string `params:"id"`
}

type QuotationResponse struct {
	ID           uint                 `json:"id"`
	Status       string               `json:"status"`
	Price        float64              `json:"price"`
	Package      PackageResponse      `json:"package"`
	Customer     CustomerResponse     `json:"customer"`
	Photographer PhotographerResponse `json:"photographer"`
}

func ToQuotationResponse(quotation model.Quotation) QuotationResponse {
	return QuotationResponse{
		ID:           quotation.ID,
		Status:       quotation.Status.String(),
		Price:        quotation.Price,
		Package:      ToPackageResponse(quotation.Package),
		Customer:     ToCustomerResponse(quotation.Customer),
		Photographer: ToPhotographerResponse(quotation.Photographer),
	}
}

func ToQuotationResponses(quotations []model.Quotation) []QuotationResponse {
	var responses []QuotationResponse
	for _, quotation := range quotations {
		responses = append(responses, ToQuotationResponse(quotation))
	}
	return responses
}
