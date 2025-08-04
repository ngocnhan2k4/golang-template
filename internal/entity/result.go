package entity

type Result struct {
	Success      bool
	Data         interface{}
	Message      *string
	ErrorCode    string
	ErrorMessage string
	Errors       []APIError
}

func Ok(data interface{}, message *string) Result {
	return Result{
		Success: true,
		Data:    data,
		Message: message,
	}
}

func Fail(errorCode string, errorMessage string, errors []APIError) Result {
	return Result{
		Success:      false,
		Data:         nil,
		Message:      nil,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		Errors:       errors,
	}
}
