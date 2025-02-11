package dto

type HttpResponse[T any] struct {
	Result T `json:"result" doc:"Result"`
}

type HttpError struct {
	Error string `json:"error" doc:"Error message"`
}
