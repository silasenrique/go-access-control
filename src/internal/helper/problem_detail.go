package helper

type ProblemDetail string

const (
	ValidationProblemDetail  ProblemDetail = "A entidade Pessoa que está sendo cadastrada não possui todos os dados corretos"
	InternalProblemDetail    ProblemDetail = "Não foi possível realizar a validação dos dados a serem cadastrados"
	CodeExistsDetails        ProblemDetail = "O código informado já foi utilizado anteriormente"
	SQLNotFoundProblemDetail ProblemDetail = "Houve um erro interno e não foi possível recuperar um recurso"
	SQLCreateProblemDetail   ProblemDetail = "Houve um erro interno e não foi possível cadastrar um recurso"
)

func (p ProblemDetail) String() string {
	return string(p)
}
