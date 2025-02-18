package apperror

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func New(code int, msg string, err error) *AppError {
	var Err error
	if err == nil {
		Err = fmt.Errorf("%s", msg)
	} else {
		Err = fmt.Errorf("%s: %v", msg, err)
	}

	return &AppError{
		Code:    code,
		Message: msg,
		Err:     Err,
	}
}

func Internal(msg string, err error) *AppError {
	return New(http.StatusInternalServerError, msg, err)
}

func BadRequest(msg string, err error) *AppError {
	return New(http.StatusBadRequest, msg, err)
}

func NotFound(msg string, err error) *AppError {
	return New(http.StatusNotFound, msg, err)
}

func UnAuthorized(msg string, err error) *AppError {
	return New(http.StatusUnauthorized, msg, err)
}

func Forbidden(msg string, err error) *AppError {
	return New(http.StatusForbidden, msg, err)
}
