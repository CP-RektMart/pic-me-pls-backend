package apperror

import (
	"fmt"
	"net/http"

	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
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
	logger.Error(msg, err)
	return &AppError{
		Code:    code,
		Message: msg,
		Err:     fmt.Errorf("%s: %v", msg, err),
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
