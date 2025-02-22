package dto

type MediaGalleryRequest struct {
	PictureURL  string `json:"pictureUrl" validate:"required"`
	Description string `json:"description"`
}

type CreateGalleryRequest struct {
	Name        string                `json:"name" validate:"required"`
	Description string                `json:"description" validate:"required"`
	Price       float64               `json:"price" validate:"required"`
	Media       []MediaGalleryRequest `json:"media" validate:"required"`
}

type UpdateGalleryRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       *float64 `json:"price"`
	GalleryId   uint     `params:"galleryId"`
}

type UpdateGalleryResponse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
