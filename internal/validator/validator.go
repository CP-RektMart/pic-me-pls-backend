package validator

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/constant"
	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("role", validateString(constant.ValidateRole))
	validate.RegisterValidation("provider", validateString(constant.ValidateAuthProvider))

	return validate
}

type EnumValidator func(field string) bool

func validateString(fn EnumValidator) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return fn(fl.Field().String())
	}
}
