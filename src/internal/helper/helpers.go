package helper

var urlHelpers = map[error]string{
	ErrInternal:                 "deu ruim",
	ErrValidationStruct:         "https://github.com/silasenrique/go-access-control/wiki/Error-Index#estrutura-com-dados-inv%C3%A1lidos",
	ErrInternalValidationStruct: "https://github.com/silasenrique/go-access-control/wiki/Error-Index#n%C3%A3o-foi-poss%C3%ADvel-realizar-a-valida%C3%A7%C3%A3o-da-estrutura",
	ErrCodeAlreadyExists:        "https://github.com/silasenrique/go-access-control/wiki/Error-Index#houve-uma-tentativa-de-reutilizar-um-dado-%C3%BAnico",
	ErrSQLNotFound:              "https://github.com/silasenrique/go-access-control/wiki/Error-Index#n%C3%A3o-foi-poss%C3%ADvel-recuperar-um-recurso",
	ErrSQLCreateFailure:         "https://github.com/silasenrique/go-access-control/wiki/Error-Index#n%C3%A3o-foi-poss%C3%ADvel-cadastrar-um-recurso",
}
