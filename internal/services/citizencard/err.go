package citizencard

import "github.com/cockroachdb/errors"

var (
	ErrNoExistingCitizenCard = errors.New("NO_EXISTING_CITIZEN_CARD")
	ErrAlreadyVerified       = errors.New("ALREADY_VERIFIED")
)
