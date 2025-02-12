package dto

import (
	"mime/multipart"
	"time"
)

type CitizenCardRequest struct {
	CitizenID  string                `from:"citizenId" validate:"required"`
	LaserID    string                `from:"laserId" validate:"required"`
	Picture    *multipart.FileHeader `from:"picture" validate:"required"`
	ExpireDate time.Time             `from:"expireDate" validate:"required"`
}

type CitizenCardResponse struct {
	CitizenID  string    `json:"citizenId"`
	LaserID    string    `json:"laserId"`
	PictureURL string    `json:"picture_url"`
	ExpireDate time.Time `json:"expireDate"`
}
