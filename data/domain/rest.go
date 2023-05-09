package domain

type Response struct {
	Status       string `json:"status"`
	InnerMessage string `json:"inner_message,omitempty"`
	Message      string `json:"message"`
	Body         any    `json:"body,omitempty"`
}
