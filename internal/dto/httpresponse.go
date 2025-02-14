package dto

type HttpResponse struct {
	Result interface{} `json:"result"`
}

type HttpError struct {
	Error string `json:"error"`
}
