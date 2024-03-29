package helper

import (
	"go-access-control/src/pkg/problem"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type problemValidation struct {
	problem         *problem.Problem
	validationError error
}

func newProblemValidation(err error) *problem.Problem {
	problem := &problemValidation{
		problem: problem.NewProblem(
			Validation,
			ErrValidationStruct.Error(),
			ValidationProblemDetail,
			urlHelpers[ErrValidationStruct],
			http.StatusInternalServerError,
		),
		validationError: err,
	}
	problem.problem.SetStatusCode(http.StatusBadRequest)
	problem.addValidationDetails()

	return problem.problem
}

func (p *problemValidation) addValidationDetails() {
	for _, err := range p.validationError.(validator.ValidationErrors) {
		detail := problem.NewProblemDetail(
			err.Namespace(),
			err.Tag(),
			err.Tag(),
			err.Tag(),
			err.Error(),
		)

		p.problem.AddDetails(detail)
	}
}

func IsValidationProblem(err error) bool {
	_, isType := err.(validator.ValidationErrors)
	return isType
}
