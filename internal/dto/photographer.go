package dto

import (
	"mime/multipart"
	"time"
)

type CitizenCardRequest struct {
	CitizenID  string                `from:"citizenId"`
	LaserID    string                `from:"laserId"`
	Picture    *multipart.FileHeader `from:"picture"`
	ExpireDate time.Time             `from:"expireDate"`
}

type CitizenCardResponse struct {
	CitizenID  string    `json:"citizenId"`
	LaserID    string    `json:"laserId"`
	PictureURL string    `json:"picture_url"`
	ExpireDate time.Time `json:"expireDate"`
}
