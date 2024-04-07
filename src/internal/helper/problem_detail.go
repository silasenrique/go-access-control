package helper

const (
	InternalErrorDetail             string = "Um problema não previsto e não mapeado foi retornado"
	ValidationProblemDetail         string = "A entidade que está sendo cadastrada não possui todos os dados corretos"
	ValidationInternalProblemDetail string = "Não foi possível realizar a validação dos dados a serem cadastrados"
	CodeExistsDetails               string = "O código informado já foi utilizado anteriormente"
	SQLNotFoundProblemDetail        string = "Houve um erro interno e não foi possível recuperar um recurso"
	SQLCreateProblemDetail          string = "Houve um erro interno e não foi possível cadastrar um recurso"
	TheIdCannotBeEmpty              string = "Não foi informado o id para realizar a busca do recurso"
	TheCodeCannotBeEmpty            string = "Não foi informado o `code` para realizar a busca do recurso"
	StringToIntegerConversion       string = "Não foi possível realizar a conversão de `texto` para `numero`"
	PasswordNotStrong               string = "A senha não atende os requisitos minimos"
)
