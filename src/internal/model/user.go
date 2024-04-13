package model

import (
	"go-access-control/src/internal/helper"
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id             int
	CompanyCode    string
	Email          string
	Password       *Password
	CompleteName   string
	CreationDate   int64
	LastChangeDate int64
}

func NewUser(companyCode, completeName, email string) *User {
	return &User{
		CompanyCode:    companyCode,
		CompleteName:   completeName,
		Email:          email,
		CreationDate:   time.Now().Unix(),
		LastChangeDate: time.Now().Unix(),
	}
}

func (u *User) AddPassword(pass *Password) {
	u.Password = pass
}

func (u *User) Validate() error {
	err := validator.ValidateStruct(u,
		validator.Field(&u.CompanyCode, validator.Required, validator.Length(3, 10)),
		validator.Field(&u.Email, validator.Required, is.Email),
		validator.Field(&u.CompleteName, validator.Required, validator.Length(2, 100)),
	)

	if err != nil {
		return helper.OzzoValidationError(err)
	}

	return nil
}
