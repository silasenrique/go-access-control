package model

import (
	"go-access-control/src/internal/helper"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Password        string
	Reset           bool
	Expiration      int64
	LatestResetDate int64
}

func NewPassword(password string, reset bool, expiration int64) *Password {
	return &Password{Password: password, Reset: reset, Expiration: expiration}
}

func (p *Password) IsStronger() error {
	if len(p.Password) < 12 {
		return helper.ErrPasswordWrongSize
	}

	validate := regexp.MustCompilePOSIX("^(?=.*[0-9]{3})(?=.*[a-z]{3})(?=.*[A-Z]{3})(?=.*[@#$%^&+=]){3}.{12,}$")

	strong := validate.MatchString(p.Password)
	if !strong {
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
