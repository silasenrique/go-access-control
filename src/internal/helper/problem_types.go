package helper

type ProblemType int

const (
	Internal ProblemType = iota
	Business
	Validation
)

var problems = map[ProblemType]string{
	Validation: "validation",
	Internal:   "internal",
	Business:   "business",
}

func (d ProblemType) String() string {
	return problems[d]
}
