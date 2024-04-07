package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-access-control/src/internal/helper"
	"go-access-control/src/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Create(user *model.User) error {
	sql := `insert into "user" (DEFAULT_COMPANY_CODE, EMAIL, COMPLETE_NAME, PASSWORD, RESET_PASSWORD, EXPIRATION_DATE)`
	sql = fmt.Sprintf(`%s values ($1, $2, $3, $4, $5, $6)`, sql)

	_, err := u.db.Exec(
		sql,
		user.CompanyCode,
		user.Email,
		user.CompleteName,
		user.Password.Password,
		user.Password.Reset,
		user.Password.Expiration,
	)

	if err != nil {
		return helper.NewHelper(helper.ErrSQLCreateFailure).AddIntenal(err)
	}

	return nil
}

func (u *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := `select * from "user" where email = $1`
	user := &model.User{
		Password: &model.Password{},
	}

	err := u.db.QueryRow(query, email).
		Scan(
			user.Id,
			user.CompanyCode,
			user.Email,
			user.CompleteName,
			user.Password.Password,
			user.Password.Reset,
			user.Password.LatestResetDate,
			user.Password.Expiration,
			user.CreationDate,
			user.LastChangeDate,
		)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, helper.NewHelper(helper.ErrSQLNotFound).AddIntenal(err)
	}

	return user, nil
}
