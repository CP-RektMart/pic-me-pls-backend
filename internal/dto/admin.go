package dto

type ListPhotographersRequest struct {
	PaginationRequest
	Name *string `query:"name"`
}
