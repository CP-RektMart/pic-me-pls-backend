package dto

type CreateGalleryRequest struct {
	Name        string  `validate:"required"`
	Description string  `validate:"required"`
	Price       float64 `validate:"required"`
}

type CreateGalleryResponse struct {
	ID               uint     `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Price            float64  `json:"price"`
	PhotographerID   uint     `json:"photographerId"`
	PhotographerName string   `json:"photographerName"`
	GalleryPhotos    []string `json:"galleryPhotos"`
}
