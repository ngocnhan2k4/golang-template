package entity

type APIError struct {
	Index   int    `json:"index,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type APIResponse struct {
	Message *string     `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []APIError  `json:"errors,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

func Success(title string, data interface{}, message *string) APIResponse {
	return APIResponse{
		Data:    data,
		Errors:  nil,
		Message: message,
		Error:   nil,
	}
}

func BadRequest(title string, data interface{}, error *APIError, message *string, errors []APIError) APIResponse {
	return APIResponse{
		Data:    data,
		Error:   error,
		Message: message,
		Errors:  errors,
	}
}

func NotFound(title string, data interface{}, error APIError, message *string) APIResponse {
	return APIResponse{
		Data:    data,
		Error:   &error,
		Message: message,
	}
}

func InternalServerError(title string, data interface{}, error APIError, message *string) APIResponse {
	return APIResponse{
		Data:    data,
		Error:   &error,
		Message: message,
	}
}

func MultiStatus(title string, data interface{}, error APIError, message *string, errors []APIError) APIResponse {
	return APIResponse{
		Data:    data,
		Error:   &error,
		Message: message,
		Errors:  errors,
	}
}
