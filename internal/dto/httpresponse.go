package dto

type HttpResponse[T any] struct {
	Result T `json:"result"`
}

type HumaHttpResponse[T any] struct {
	Body HttpResponse[T]
}

type HttpError struct {
	Error string `json:"error"`
}

func Response[T any](data T) *HumaHttpResponse[T] {
	return &HumaHttpResponse[T]{
		Body: HttpResponse[T]{Result: data},
	}
}
