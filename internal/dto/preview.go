package dto

type CreatePreviewPhotoRequest struct {
	Link        string `json:"link" validate:"required"`
	QuotationID uint   `json:"quotationId" validate:"required"`
}
