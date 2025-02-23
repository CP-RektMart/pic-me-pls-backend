package dto

import "github.com/danielgtaylor/huma/v2"

type HttpResponse[T any] struct {
	Result T `json:"result"`
}

type HumaHttpResponse[T any] struct {
	Body HttpResponse[T]
}

type HttpError struct {
	Error string `json:"error"`
}

type HumaBody[T any] struct {
	Body T
}

type HumaFormData[T any] struct {
	RawBody huma.MultipartFormFiles[T]
}
