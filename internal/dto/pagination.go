package dto

type PaginationResponse[T any] struct {
	Page       int   `json:"page"`
	Total      int64 `json:"total"`
	PageSize   int   `json:"page_size"`
	TotalPages int   `json:"total_pages"`
	Response   []T   `json:"response"`
}
