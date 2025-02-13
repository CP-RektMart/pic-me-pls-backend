package dto

type GalleryDTO struct {
	ID           uint                 `json:"id"`
	Name         string               `json:"name"`
	Description  string               `json:"description,omitempty"`
	Price        float64              `json:"price"`
	Photographer PhotographerResponse `json:"photographer"`
	Tags         []TagResponse        `json:"tags,omitempty"`
	Media        []MediaResponse      `json:"media,omitempty"`
	Reviews      []ReviewResponse     `json:"reviews,omitempty"`
	Categories   []CategoryResponse   `json:"categories,omitempty"`
	Quotations   []QuotationResponse  `json:"quotations,omitempty"`
}
