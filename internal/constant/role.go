package constant

type Role string

const CustomerRole Role = "customer"
const PhotographerRole Role = "photographer"

func (r Role) String() string {
	return string(r)
}

func ValidateRole(role string) bool {
	switch Role(role) {
	case CustomerRole, PhotographerRole:
		return true
	}
	return false
}
