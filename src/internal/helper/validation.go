package helper

import (
	"errors"
	"go-access-control/src/pkg/problem"

	"github.com/go-playground/validator/v10"
)

type ProblemValidation struct {
	problem         *problem.Problem
	validationError error
}

func NewProblemValidation(err error) *problem.Problem {
	problem := &ProblemValidation{
		problem: problem.NewProblem(
			Validation,
			ValidationProblemTitle,
			ValidationProblemDetail,
			"https://github.com/silasenrique/api-heper",
		),
		validationError: err,
	}

	problem.AddValidationDetails()

	return problem.problem
}

func (p *ProblemValidation) AddValidationDetails() {
	for _, err := range p.validationError.(validator.ValidationErrors) {
		detail := problem.NewProblemDetail(
			err.Namespace(),
			getValidationErrorTitle(err.Tag()),
			getErr(err.Tag()),
			getInstanceUrl(err.Tag()),
			err.Error(),
		)

		p.problem.AddDetails(detail)
	}
}

func IsValidationProblem(err error) bool {
	return errors.Is(err, validator.ValidationErrors{})
}

var validationErrorTitle = map[string]string{
	"required": "Atributo obrigatório não informado",
	"gte":      "Atributo não atingiu a quantidade minima requerida",
}

func getValidationErrorTitle(errType string) string {
	mess := validationErrorTitle[errType]
	if mess != "" {
		return mess
	}

	return errType
}

var instanceUrl = map[string]string{
	"required": "https://github.com/silasenrique/api-heper/required",
	"gte":      "https://github.com/silasenrique/api-heper/gte",
}

func getInstanceUrl(instance string) string {
	if instanceUrl[instance] == "" {
		return instance
	}

	return instanceUrl[instance]
}

var typeErr = map[string]string{
	"gte":      "O atributo não atingiu a quantidade minima de caracteres esperado",
	"lte":      "O atributo ultrapassou a quantidade de caracteres esperado",
	"required": "Um atributo obrigatório não foi informado",
}

func getErr(tag string) string {
	if typeErr[tag] == "" {
		return tag
	}

	return typeErr[tag]
}
