package model

import (
	"go-access-control/src/internal/helper"

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
		return helper.NewValidationError(nil, false, 40004, err)
	}

	return nil
}
