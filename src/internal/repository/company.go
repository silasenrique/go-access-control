package repository

import (
	"database/sql"
	"errors"
	"go-access-control/src/internal/model"
	"go-access-control/src/pkg/problem"
)

type CompanyRepository struct {
	*sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{db}
}

func (c *CompanyRepository) FindByCode(code string) (*model.Company, error) {
	query := "select * from company a where a.code = $1"

	var company model.Company

	err := c.QueryRow(query, code).Scan(
		&company.Id,
		&company.Code,
		&company.Name,
		&company.SiteUrl,
		&company.CreationDate,
		&company.LastChangeDate,
	)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, problem.NewProblemError(
				problem.Internal,
				"Não foi possível recuperar um recurso",
				"Houve um erro interno e não foi possível recuperar um recurso",
				"https://github.com/silasenrique/api-heper",
				err.Error(),
			)
		}
	}

	return &company, nil
}

func (c *CompanyRepository) Create(company *model.Company) error {
	sql := "insert into company (code, name, site_url, creation_date, last_change_date) values ($1, $2, $3, $4, $5)"

	_, err := c.Exec(
		sql,
		company.Code,
		company.Name,
		company.SiteUrl,
		company.CreationDate,
		company.LastChangeDate,
	)
	if err != nil {
		return problem.NewProblemError(
			problem.Internal,
			"Erro no momento de cadastrar um recurso",
			"Houve um erro interno e não foi possível cadastrar um recurso",
			"https://github.com/silasenrique/api-heper",
			err.Error(),
		)
	}

	return nil
}
