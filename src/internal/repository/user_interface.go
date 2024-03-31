package repository

import "go-access-control/src/internal/model"

type IUserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type IUserCreateRespository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
