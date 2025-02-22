package dto

type PaginationResponse[T any] struct {
	Page        int   `json:"page"`
	Total       int64 `json:"total"`
	Limit       int   `json:"limit"`
	TotalPages  int   `json:"total_pages"`
	HasNextPage bool  `json:"has_next_page"`
	HasPrevPage bool  `json:"has_prev_page"`
	Response    []T   `json:"response"`
}
