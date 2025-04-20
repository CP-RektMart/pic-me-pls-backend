package dto

type ListPhotographersRequest struct {
	PaginationRequest
	Name *string `query:"name"`
}

type AdminGetPhotographerByIDRequest struct {
	PhotographerID uint `params:"photographerID" validate:"required"`
}

type ListUnverifiedPhotographerRequest struct {
	PaginationRequest
	Name *string `query:"name"`
}

type BanPhotographerRequest struct {
	ID uint `json:"id" validate:"required"`
}
