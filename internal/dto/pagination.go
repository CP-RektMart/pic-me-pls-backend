package dto

type PaginationResponse[T any] struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalPage int `json:"totalPage"`
	Data      []T `json:"data"`
}
