package validator

import (
	"github.com/54m/api_gen-example/backend/domain/werror"
	"github.com/go-playground/validator/v10"
)

// Validator - validator
type Validator struct {
	validator *validator.Validate
}

// NewValidator - constructor
func NewValidator() *Validator {
	return &Validator{validator: validator.New()}
}

// Validate - validate
func (v *Validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	return werror.NewValidationError(err)
}
