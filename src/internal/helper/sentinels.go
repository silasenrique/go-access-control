package helper

import "errors"

var (
	ErrInternal = errors.New("erro interno não previsto")

	// estrutura com dados inválidos
	ErrValidationStruct = errors.New("estrutura com dados inválidos")

	// ocorreu um erro interno ao validar o conteudo a ser cadastrado
	ErrInternalValidationStruct = errors.New("ocorreu um erro interno ao validar o conteudo a ser cadastrado")

	// o valor utilizado no atributo 'code' já foi cadastrado anteriormente
	ErrCodeAlreadyExists = errors.New("o valor utilizado no atributo 'code' já foi cadastrado anteriormente")

	// não foi possível recuperar um recurso
	ErrSQLNotFound = errors.New("não foi possível recuperar um recurso")

	// não foi possível criar o recurso
	ErrSQLCreateFailure = errors.New("não foi possível criar o recurso")

	// o id não pode ser nulo
	ErrTheIdCannotBeEmpty = errors.New("o id não pode ser nulo")

	// o id não pode ser nulo
	ErrTheCodeCannotBeEmpty = errors.New("o code não pode ser nulo")

	// erro de conversão de string para inteiro
	ErrStringToIntegerConversion = errors.New("erro de conversão de string para inteiro")
)
