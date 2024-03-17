package repository

import (
	"database/sql"
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
	sql := "select * from company a where a.code = ?"

	var company model.Company

	err := c.QueryRow(sql, code).Scan(&company)
	if err != nil {
		return nil, problem.NewProblemError(
			problem.Internal,
			"Não foi possível recuperar um recurso",
			"Houve um erro interno e não foi possível recuperar um recurso",
			"https://github.com/silasenrique/api-heper",
			err.Error(),
		)
	}

	return &company, nil
}

func (c *CompanyRepository) Create(company *model.Company) error {
	sql := "insert into company values (?, ?, ?, sysdate, sysdate)"

	_, err := c.Exec(sql, company.Code, company.Name, company.SiteUrl)
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
