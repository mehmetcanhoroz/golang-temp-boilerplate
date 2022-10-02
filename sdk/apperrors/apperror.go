package apperrors

type AppError struct {
	Error       error
	RestMessage string
}

func NewAppError(restMessage string, err error) *AppError {

	if restMessage == "" && err == nil {
		return nil
	}
	return &AppError{
		Error:       err,
		RestMessage: restMessage,
	}
}
