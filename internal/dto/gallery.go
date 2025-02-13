package dto

type getGalleriesResponse struct {
	ID             uint    `json:"id"`
	PhotographerID uint    `json:"photographerId"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
}
