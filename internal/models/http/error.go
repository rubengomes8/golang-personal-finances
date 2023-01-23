package http

// ErrorResponse is the error model for http responses
type ErrorResponse struct {
	ErrorMsg string `json:"error,omitempty"`
}
