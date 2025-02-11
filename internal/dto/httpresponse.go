package dto

type HttpResponse struct {
	Result interface{} `json:"result" doc:"Result"`
	Error  string      `json:"error" doc:"Error message"`
}
