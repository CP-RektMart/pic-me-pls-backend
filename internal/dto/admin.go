package dto

type ListPhotographersRequest struct {
	PaginationRequest
	Name *string `query:"name"`
}

type AdminGetPhotographerByIDRequest struct {
	PhotographerID uint `params:"photographerID" validate:"required"`
}
