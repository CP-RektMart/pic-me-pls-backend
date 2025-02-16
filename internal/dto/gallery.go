package dto

type GalleryRequest struct {
	Name        string  `validate:"required"`
	Description string  `validate:"required"`
	Price       float64 `validate:"required"`
}

type GalleryResponse struct {
	ID               uint     `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Price            float64  `json:"price"`
	PhotographerID   uint     `json:"photographerId"`
	PhotographerName string   `json:"photographerName"`
	GalleryPhotos    []string `json:"galleryPhotos"`
}
