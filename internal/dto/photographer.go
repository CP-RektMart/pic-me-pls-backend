package dto

import "time"

type CitizenCardRequest struct {
<<<<<<< HEAD
	CitizenID  string    `json:"citizenId" validate:"required"`
	LaserID    string    `json:"laserId" validate:"required"`
	Picture    string    `json:"picture" validate:"required"`
	ExpireDate time.Time `json:"expireDate" validate:"required"`
=======
	CitizenID  string                `from:"citizenId"`
	LaserID    string                `from:"laserId"`
	Picture    *multipart.FileHeader `from:"picture"`
	ExpireDate time.Time             `from:"expireDate"`
>>>>>>> parent of a3bbe97 (fix: dto required field)
}

type CitizenCardResponse struct {
	CitizenID  string    `json:"citizenId"`
	LaserID    string    `json:"laserId"`
	Picture    string    `json:"picture"`
	ExpireDate time.Time `json:"expireDate"`
}
