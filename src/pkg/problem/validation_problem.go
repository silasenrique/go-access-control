package problem

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var typeErr = map[string]string{
	"gte":      "O atributo não atingiu a quantidade minima de caracteres esperado",
	"lte":      "O atributo ultrapassou a quantidade de caracteres esperado",
	"required": "Um atributo obrigatório não foi informado",
}

func getErr(tag string) string {
	return typeErr[tag]
}

func IsValidationProblem(err error) bool {
	return errors.As(err, &validator.ValidationErrors{})
}
