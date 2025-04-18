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

func (p *PaginationRequest) CheckPaginationRequest() (page int, pageSize int, offset int) {
	page, pageSize = p.Page, p.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset = (page - 1) * pageSize
	return
}

func GetPaginationData(req PaginationRequest, defaultPage, defaultPageSize int) (page, pageSize, offset int) {
	page, pageSize = req.Page, req.PageSize

	if req.Page == 0 {
		page = defaultPage
	}
	if req.PageSize == 0 {
		pageSize = defaultPageSize
	}

	offset = (page - 1) * pageSize
	return
}

func (p *PaginationRequest) GetPaginationData(defaultPage, defaultPageSize int) (page, pageSize, offset int) {
	page, pageSize = p.Page, p.PageSize

	if p.Page == 0 {
		page = defaultPage
	}
	if p.PageSize == 0 {
		pageSize = defaultPageSize
	}

	offset = (page - 1) * pageSize
	return
}

func NewPaginationResponse[T any](data []T, page, pageSize, totalPage int) PaginationResponse[T] {
	return PaginationResponse[T]{
		Data:      data,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
	}
}
