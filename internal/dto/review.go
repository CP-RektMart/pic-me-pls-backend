package dto

type ReviewResponse struct {
	ID       uint    `json:"id"`
	Rating   float64 `json:"rating"`
	Comment  string  `json:"comment,omitempty"`
	Customer string  `json:"customer"`
}
