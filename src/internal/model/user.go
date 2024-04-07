package model

import (
	"go-access-control/src/internal/helper"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id             int
	CompanyCode    string `validate:"required,gte=3,lte=10"`
	Email          string `validate:"required,email"`
	Password       *Password
	CompleteName   string `validate:"required,gte=3,lte=200"`
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
	err := validator.New().Struct(u)
	if err != nil {
		if helper.IsValidationProblem(err) {
			return helper.NewHelper(helper.ErrValidationStruct).AddIntenal(err)
		}

		return helper.NewHelper(helper.ErrInternalValidationStruct).AddIntenal(err)
	}

	return nil
}
