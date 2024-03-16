package repository

import (
	"database/sql"
	"go-access-control/src/internal/helper"
	"go-access-control/src/internal/model"
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
		return nil, helper.NewDatabaseError(false, 20001, err)
	}

	return &company, nil
}

func (c *CompanyRepository) Create(company *model.Company) error {
	sql := "insert into company values (?, ?, ?, sysdate, sysdate)"

	_, err := c.Exec(sql, company.Code, company.Name, company.SiteUrl)
	if err != nil {
		return helper.NewDatabaseError(false, 20001, err)
	}

	return nil
}
