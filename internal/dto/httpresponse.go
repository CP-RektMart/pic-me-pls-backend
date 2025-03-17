package dto

type HttpResponse[T any] struct {
	Result T `json:"result"`
}

type HttpError struct {
	Error string `json:"error"`
}

// this struct is created to fix swaggo error with slice generic ex. HttpResponse[[]dto.something] => panic
type HttpListResponse[T any] struct {
	Result []T `json:"result"`
}
