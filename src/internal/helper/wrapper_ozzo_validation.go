package helper

import (
	"go-access-control/src/pkg/problem/v2"

	validator "github.com/go-ozzo/ozzo-validation"
)

func OzzoValidationError(err error) error {
	if e, ok := err.(validator.InternalError); ok {
		return problem.NewProblem(problem.Internal, "a validação falhou").AddError(e)
	}

	return problem.NewProblem(problem.Validation, "a validação falhou").
		AddError(err).
		AddDetail("as informações inseridas na estrutura não estão corretas")
}
