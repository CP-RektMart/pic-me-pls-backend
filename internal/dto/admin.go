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

type VerifyPhotographerRequest struct {
	PhotographerID uint `params:"photographerId" validate:"required"`
}

type BanPhotographerRequest struct {
	ID uint `json:"id" validate:"required"`
}

type AdminDeletePackageByID struct {
	PackageID uint `params:"packageID" validate:"required"`
}

type AssignAdminRequest struct {
	UserID uint  `params:"userID" validate:"required"`
	Admin  *bool `json:"admin" validate:"required"`
}
