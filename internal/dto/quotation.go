package dto

import "time"

type CreateQuotationRequest struct {
	GalleryID   uint    `json:"gallery_id" validate:"required"`
	CustomerID  uint    `json:"customer_id" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description,omitempty"`

	FromDate time.Time `json:"from_date" validate:"required"`
	ToDate   time.Time     `json:"to_date" validate:"required"`
}

type UpdateQuotationRequest struct {
	CreateQuotationRequest
	Status string `json:"status" validate:"required"`
}

type QuotationResponse struct {
	QuotationID uint `json:"quotation_id"`
}