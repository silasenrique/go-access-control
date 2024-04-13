package helper

import (
	"go-access-control/src/pkg/problem/v2"

	validator "github.com/go-ozzo/ozzo-validation"
)

func OzzoValidationError(err error) error {
	if e, ok := err.(validator.InternalError); ok {
		return problem.NewProblem(problem.Internal, "não foi possível validar os dados enviados").AddError(e)
	}

	return problem.NewProblem(problem.Validation, "os dados enviados estão incorretos").
		AddDetail("Os dados informados não seguem as regras especificadas para os atributos").
		AddValidationError(err)
}
