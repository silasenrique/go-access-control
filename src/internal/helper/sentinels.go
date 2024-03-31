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

	// erro de conversão de string para inteiro
	ErrPasswordNotStrong = errors.New("a senha não é forte o suficiente")

	// erro de conversão de string para inteiro
	ErrPasswordWrongSize = errors.New("a senha não tem o tamanho minimo")

	// a empresa não existe
	ErrCompanyNotExists = errors.New("a empresa não existe")

	// o email já foi cadastrado anteriormente
	ErrUserEmailExists = errors.New("o email já foi cadastrado anteriormente")
)
