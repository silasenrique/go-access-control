package model

import (
	"errors"
	"go-access-control/src/internal/helper"

	"time"

	"github.com/go-playground/validator/v10"
)

type Company struct {
	Id             int
	Code           string `validate:"required,gte=3,lte=10"`
	Name           string `validate:"required,gte=10,lte=100"`
	SiteUrl        string `validate:"lte=200"`
	CreationDate   int64
	LastChangeDate int64
}

func NewCompany(code, name, url string) *Company {
	return &Company{
		Code:           code,
		Name:           name,
		SiteUrl:        url,
		CreationDate:   time.Now().Unix(),
		LastChangeDate: time.Now().Unix(),
	}
}

func (c *Company) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		if helper.IsValidationProblem(err) {
			return errors.Join(helper.ErrValidation, err)
		}

		return helper.ErrInternal
	}

	return nil
}
