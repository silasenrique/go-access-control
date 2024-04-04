package services

import (
	"database/sql"
	"errors"
	"go-access-control/src/internal/api/dto"
	"go-access-control/src/internal/helper"
	"go-access-control/src/internal/model"
	"go-access-control/src/internal/repository"
	"time"
)

type UserService struct {
	companyRep *repository.CompanyRepository
	userRep    repository.IUserCreateRespository
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		companyRep: repository.NewCompanyRepository(db),
		userRep:    repository.NewUserRepository(db),
	}
}

func (u *UserService) CreateUser(userRequest *dto.UserCreateRequest) (*dto.UserCreateResponse, error) {
	user := model.NewUser(userRequest.CompanyCode, userRequest.Email)
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	_, err = u.companyRep.FindByCode(user.CompanyCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, helper.NewHelper(helper.ErrCompanyNotExists)
		}

		return nil, err
	}

	userExist, err := u.userRep.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if userExist != nil {
		return nil, helper.NewHelper(helper.ErrUserEmailExists)
	}

	dtExpiration, err := time.Parse(userRequest.Expiration, time.DateOnly)
	if err != nil {
		return nil, err
	}

	pass := model.NewPassword(
		userRequest.Password,
		userRequest.ResetFirstAccess,
		dtExpiration.Unix(),
	)

	err = pass.IsStronger()
	if err != nil {
		return nil, err
	}

	err = pass.Hash()
	if err != nil {
		return nil, err
	}

	user.AddPassword(pass)

	err = u.userRep.Create(user)
	if err != nil {
		return nil, err
	}

	return dto.NewUserCreateResponse(user), nil
}
