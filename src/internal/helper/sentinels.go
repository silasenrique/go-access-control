package helper

import "errors"

var (
	ErrValidation = errors.New("estrutura com dados inválidos")
	ErrInternal   = errors.New("não foi possível realizar a validação da estrutura")
)
