package dto

type ChatResponse struct {
	User     PublicUserResponse        `json:"user"`
	Messages []RealTimeMessageResponse `json:"messages"`
}
