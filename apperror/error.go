package apperror

type AppError struct {
	ErrCode
	Message string
	Errors  []ErrDetail
	Err     error
}

type ErrDetail struct {
	Source  string `json:"source"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {

	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}
