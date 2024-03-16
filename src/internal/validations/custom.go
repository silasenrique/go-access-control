package validations

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property   string `json:"property"`
	InputValue any    `json:"inputValue"`
	Message    string `json:"message"`
	Type       string `json:"type"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var (
		validationErrors []ValidationError
		ve               validator.ValidationErrors
	)

	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			validationError := ValidationError{
				Property:   err.Field(),
				InputValue: err.Value(),
				Type:       err.Tag(),
				Message:    getErr(err.Tag(), err.Error(), err.Param()),
			}

			validationErrors = append(validationErrors, validationError)
		}

		return &validationErrors
	}

	return nil
}
