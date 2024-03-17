package problem

import "github.com/go-playground/validator/v10"

type ProblemDetail struct {
	Field    string `json:"field"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func NewProblemDetailForValidationError(err validator.FieldError) *ProblemDetail {
	return &ProblemDetail{
		Field:    err.Namespace(),
		Title:    getValidationErrorTitle(err.Tag()),
		Detail:   getErr(err.Tag()),
		Instance: getInstanceUrl(err.Tag()),
	}
}
