package utils


type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}