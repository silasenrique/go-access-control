package helper

const (
	ValidationProblemDetail         string = "A entidade que está sendo cadastrada não possui todos os dados corretos"
	ValidationInternalProblemDetail string = "Não foi possível realizar a validação dos dados a serem cadastrados"
	CodeExistsDetails               string = "O código informado já foi utilizado anteriormente"
	SQLNotFoundProblemDetail        string = "Houve um erro interno e não foi possível recuperar um recurso"
	SQLCreateProblemDetail          string = "Houve um erro interno e não foi possível cadastrar um recurso"
)
