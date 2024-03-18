package problem

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type Problem struct {
	Type     string          `json:"type"`
	Title    string          `json:"title"`
	Detail   string          `json:"detail"`
	Instance string          `json:"instance"`
	Err      string          `json:"error,omitempty"`
	Errors   []ProblemDetail `json:"errors,omitempty"`
}

func (p Problem) ToHttpError() ([]byte, error) {
	return json.MarshalIndent(&p, "", "  ")
}

func (p Problem) Error() string {
	return ""
}

func (p *Problem) AddValidationDetails(err error) *Problem {
	if IsValidationProblem(err) {
		for _, err := range err.(validator.ValidationErrors) {
			p.Errors = append(p.Errors, NewProblemDetailForValidationError(err))
		}
	}

	return p
}

func NewProblem(problemType ProblemType, title, detail, instance string) *Problem {
	return &Problem{
		Type:     problemType.ToString(),
		Title:    title,
		Detail:   detail,
		Instance: instance,
	}
}

func NewProblemError(problemType ProblemType, title, detail, instance, err string) *Problem {
	return &Problem{
		Type:     problemType.ToString(),
		Title:    title,
		Detail:   detail,
		Instance: instance,
		Err:      err,
	}
}
