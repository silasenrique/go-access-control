package problem

var problems = map[ProblemType]string{
	Validation: "validation",
	Internal:   "internal",
	Business:   "business",
}

type ProblemType int

const (
	Validation ProblemType = iota + 1
	Internal
	Business
)

func (d ProblemType) ToString() string {
	return problems[d]
}
