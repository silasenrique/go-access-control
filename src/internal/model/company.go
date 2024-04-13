package model

import (
	"go-access-control/src/internal/helper"

	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Company struct {
	Id             int
	Code           string
	Name           string
	SiteUrl        string
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
	err := validator.ValidateStruct(c,
		validator.Field(&c.Code, validator.Required, validator.Length(3, 10)),
		validator.Field(&c.Name, validator.Required, validator.Length(10, 100)),
		validator.Field(&c.SiteUrl, is.URL, validator.Length(0, 200)),
	)

	if err != nil {
		return helper.OzzoValidationError(err)
	}

	return nil
}
