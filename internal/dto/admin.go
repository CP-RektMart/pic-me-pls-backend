package dto

type ListPhotographersRequest struct {
	PaginationRequest
	Name *string `query:"name"`
}

type AdminGetPhotographerByID struct {
	PhotographerID uint `params:"photographerID" validate:"required"`
}
