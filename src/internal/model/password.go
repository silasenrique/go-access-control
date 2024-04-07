package model

import (
	"go-access-control/src/internal/helper"

	regexp "github.com/dlclark/regexp2"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Password        string
	Reset           bool
	Expiration      int64
	LatestResetDate int64
}

func NewPassword(password string, reset bool, expiration int64) *Password {
	return &Password{
		Password:   password,
		Reset:      reset,
		Expiration: expiration,
	}
}

func (p *Password) IsStronger() error {
	validate, err := regexp.Compile(`^(?=.*[0-9]{3})(?=.*[a-z]{3})(?=.*[A-Z]{3})(?=.*[@#$%^&+=!]{3}).{12,}$`, 0)
	if err != nil {
		return helper.NewHelper(helper.ErrPasswordRegexpCheck).AddIntenal(err)
	}

	result, err := validate.MatchString(p.Password)
	if err != nil {
		return helper.NewHelper(helper.ErrPasswordRegexpCheck).AddIntenal(err)
	}

	if !result {
		return helper.NewHelper(helper.ErrPasswordNotStrong)
	}

	return nil
}

func (p *Password) Hash() error {
	pass, err := bcrypt.GenerateFromPassword([]byte(p.Password), 14)

	if err != nil {
		return err
	}

	p.Password = string(pass)

	return nil
}
