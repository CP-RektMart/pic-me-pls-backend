package dto

type HttpResponse[T any] struct {
	Result T `json:"result"`
}

type HttpError struct {
	Error string `json:"error"`
}
