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

func GetPaginationData(req PaginationRequest, defaultPage, defaultPageSize int) (int, int, int) {
	page, pageSize := req.Page, req.PageSize

	if req.Page == 0 {
		page = defaultPage
	}
	if req.PageSize == 0 {
		pageSize = defaultPageSize
	}

	offset := (page - 1) * pageSize

	return page, pageSize, offset
}
