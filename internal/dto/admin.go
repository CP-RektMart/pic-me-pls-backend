package dto

type ListPhotographersRequest struct {
	PaginationRequest
	Name *string `query:"name"`
}

type AdminGetPhotographerByIDRequest struct {
	PhotographerID uint `params:"photographerID" validate:"required"`
}

type AssignAdminRequest struct {
	UserID uint  `params:"userID" validate:"required"`
	Admin  *bool `json:"admin" validate:"required"`
}
