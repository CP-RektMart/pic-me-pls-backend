package constant

type AuthProvider string

const GoogleProvider AuthProvider = "google"
const FacebookProvider AuthProvider = "facebook"
const AppleProvider AuthProvider = "apple"

func (a AuthProvider) String() string {
	return string(a)
}

func ValidateAuthProvider(provider string) bool {
	switch AuthProvider(provider) {
	case GoogleProvider, FacebookProvider, AppleProvider:
		return true
	}
	return false
}
