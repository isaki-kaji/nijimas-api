package apperror

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrCode string

const (
	UnKnown             ErrCode = "U000"
	InsertDataFailed    ErrCode = "S001"
	UpdateDataFailed    ErrCode = "S002"
	GetDataFailed       ErrCode = "S003"
	DeleteDataFailed    ErrCode = "S004"
	DataNotFound        ErrCode = "S005"
	DataConflict        ErrCode = "S006"
	ValidationFailed    ErrCode = "R001"
	ReqBodyDecodeFailed ErrCode = "R002"
	BadPathParam        ErrCode = "R003"
	Unauthorized        ErrCode = "A001"
	OtherInternalErr    ErrCode = "A002"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &AppError{
		ErrCode: code,
		Message: message,
		Err:     err,
	}
}

func WrapValidationErr(validationErrs validator.ValidationErrors) error {
	var errs []ErrDetail
	for _, vErr := range validationErrs {
		errs = append(errs, ErrDetail{
			Source:  vErr.Field(),
			Message: vErr.ActualTag(),
		})
	}

	return &AppError{
		ErrCode: ValidationFailed,
		Message: "validation failed",
		Errors:  errs,
		Err:     validationErrs,
	}
}

func (code ErrCode) Equal(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.ErrCode == code
	}
	return false
}
