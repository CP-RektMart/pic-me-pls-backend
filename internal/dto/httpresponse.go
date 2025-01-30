package dto

type HttpResponse struct {
	Result any    `json:"result" doc:"Result"`
	Error  string `json:"error,omitempty" doc:"Error message"`
}
