package common

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// Validator instance
var Validator = validator.New()

// ValidationErrorJSON struct
type ValidationErrorJSON struct {
	error
	Key     string `json:"key"`
	Message string `json:"message"`
}

// Validate struct
func Validate(s interface{}) error {
	if err := Validator.Struct(s); err != nil {
		e := err.(validator.ValidationErrors)[0].(validator.FieldError)

		return ValidationErrorJSON{
			Key:     e.Field(),
			Message: e.Tag(),
		}
	}

	return nil
}
