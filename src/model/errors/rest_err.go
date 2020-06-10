package errors

type AppError struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

func NewGenericError(msg string, statusCode int) *AppError {
	return &AppError{
		Message: msg,
		Status:  statusCode,
	}
}