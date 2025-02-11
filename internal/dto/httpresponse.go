package dto

type HttpResponse struct {
	Result interface{} `json:"result" doc:"Result"`
	Error  string      `json:"error,omitempty" doc:"Error message"`
}
