package dto

type PaginationRequest struct {
	Page     int `query:"page"`
	PageSize int `query:"pageSize" min:"1"`
}

type PaginationResponse[T any] struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalPage int `json:"totalPage"`
	Data      []T `json:"data"`
}
