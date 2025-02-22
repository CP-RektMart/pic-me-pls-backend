package dto

type PaginationRequest struct {
	Page     int `query:"page" validate:"omitempty,min=1"`
	PageSize int `query:"pageSize" validate:"omitempty,min=1"`
}

type PaginationResponse[T any] struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalPage int `json:"totalPage"`
	Data      []T `json:"data"`
}
