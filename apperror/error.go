package apperror

type AppError struct {
	ErrCode
	Message string
	Errors  []ErrDetail
	Err     error
}

type ErrDetail struct {
	source  string
	message string
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}
