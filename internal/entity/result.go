package entity

type Result struct {
	Success      bool
	Data         interface{}
	Message      string
	ErrorCode    string
	ErrorMessage string
	Errors       interface{}
}

func Ok(data interface{}, message string) Result {
	return Result{
		Success: true,
		Data:    data,
		Message: message,
	}
}

func Fail(errorCode string, errorMessage string, errors interface{}) Result {
	return Result{
		Success:      false,
		Data:         nil,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		Errors:       errors,
	}
}