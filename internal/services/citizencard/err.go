package citizencard

import "errors"

var (
	ErrNoExistingCitizenCard = errors.New("NO_EXISTING_CITIZEN_CARD")
	ErrAlreadyVerified       = errors.New("ALREADY_VERIFIED")
)
