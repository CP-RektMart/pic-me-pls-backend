package dto

import (
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
)

type QuotationResponse struct {
	ID           uint                  `json:"id"`
	Package      PackageResponse       `json:"package,omitempty"`
	Customer     UserResponse          `json:"customer,omitempty"`
	Photographer PhotographerResponse  `json:"photographer,omitempty"`
	Status       model.QuotationStatus `json:"status"`
	Price        float64               `json:"price"`
	Description  string                `json:"description"`
	FromDate     time.Time             `json:"fromDate"`
	ToDate       time.Time             `json:"toDate"`
}

type AcceptQuotationRequest struct {
	QuotationID string `params:"id"`
}

func ToQuotationResponse(quotation model.Quotation) QuotationResponse {
	return QuotationResponse{
		ID:           quotation.ID,
		Package:      ToPackageResponse(quotation.Package),
		Customer:     ToUserResponse(quotation.Customer),
		Photographer: ToPhotographerResponse(quotation.Photographer),
		Status:       quotation.Status,
		Price:        quotation.Price,
		Description:  quotation.Description,
		FromDate:     quotation.FromDate,
		ToDate:       quotation.ToDate,
	}
}

func ToQuotationResponses(quotations []model.Quotation) []QuotationResponse {
	var responses []QuotationResponse
	for _, quotation := range quotations {
		responses = append(responses, ToQuotationResponse(quotation))
	}
	return responses
}
