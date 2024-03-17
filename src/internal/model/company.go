package model

import (
	"go-access-control/src/pkg/problem"

	"github.com/go-playground/validator/v10"
)

type Company struct {
	Id             int
	Code           string `validate:"required,alphanum,gte=3,lte=10"`
	Name           string `validate:"required,alphanum,gte=10,lte=100"`
	SiteUrl        string `validate:"alphanum,lte=200"`
	CreationDate   string
	LastChangeDate string
}

func NewCompany(code, name, url string) *Company {
	return &Company{Code: code, Name: name, SiteUrl: url}
}

func (c *Company) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		if problem.IsValidationProblem(err) {
			return problem.NewProblem(
				problem.Validation,
				problem.ValidationTitle,
				problem.ValidationDetail,
				"https://github.com/silasenrique/api-heper",
			).AddValidationDetails(err)
		}

		return problem.NewProblem(
			problem.Internal,
			problem.InternalValidationErrorTitle,
			problem.InternalValidationErrorDetail,
			"https://github.com/silasenrique/api-heper",
		)
	}

	return nil
}
