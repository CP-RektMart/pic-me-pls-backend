package dto

type GalleryRequest struct {
	Name        string  `validate:"required"`
	Description string  `validate:"required"`
	Price       float64 `validate:"required"`
}
