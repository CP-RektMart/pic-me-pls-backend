package dto

type ListPhotographersRequest struct {
	PaginationRequest
	Name *string `query:"name"`
}

type AdminGetPhotographerByID struct {
	PhotographerID uint `params:"photographerID" validate:"required"`
}

type ListUnverifiedCitizenCardRequest struct {
	PaginationRequest
	PhotographerName *string `query:"photographerName"`
}
