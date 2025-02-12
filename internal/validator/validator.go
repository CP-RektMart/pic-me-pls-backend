package validator

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	validate := validator.New()
	if err := validate.RegisterValidation("role", validateString(model.ValidateRole)); err != nil {
		logger.Panic("failed to register role validation", err)
	}
	if err := (validate.RegisterValidation("provider", validateString(model.ValidateProvider))); err != nil {
		logger.Panic("failed to register provider validation", err)
	}

	return validate
}

type EnumValidator func(field string) bool

func validateString(fn EnumValidator) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return fn(fl.Field().String())
	}
}
