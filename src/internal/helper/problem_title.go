package helper

type ProblemTitle string

const (
	ValidationProblemTitle         ProblemTitle  = "Estrutura com dados inválidos"
	ValidationInternalProblemTitle ProblemTitle  = "Não foi possível realizar a validação da estrutura"
	AlreadyExistsProblemTitle      ProblemTitle  = "Houve uma tentativa de reutilizar um dado único"
	SQLNotFoundProblemTitle        ProblemDetail = "Não foi possível recuperar um recurso"
	SQLCreateProblemTitle          ProblemDetail = "Não foi possível cadastrar um recurso"
)

func (t ProblemTitle) String() string {
	return string(t)
}
