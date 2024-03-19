package services

import (
	"database/sql"
	"go-access-control/src/internal/api/dto"
	"go-access-control/src/internal/helper"
	"go-access-control/src/internal/model"
	"go-access-control/src/internal/repository"
	"go-access-control/src/pkg/problem"
	"time"
)

type CompanyService struct {
	rep *repository.CompanyRepository
}

func NewCompanyService(db *sql.DB) *CompanyService {
	return &CompanyService{
		rep: repository.NewCompanyRepository(db),
	}
}

func (c *CompanyService) Create(req *dto.CompanyCreateRequest) (*dto.CompanyResponse, error) {
	company := model.NewCompany(req.Code, req.Name, req.SiteUrl)

	errValidator := company.Validate()
	if errValidator != nil {
		return nil, errValidator
	}

	findByCode, err := c.rep.FindByCode(company.Code)
	if err != nil {
		return nil, err
	}

	if findByCode != nil {
		return nil, problem.NewProblem(
			helper.Business,
			helper.AlreadyExistsProblemTitle,
			helper.CodeExistsDetails,
			"https://github.com/silasenrique/api-heper",
		)
	}

	createErr := c.rep.Create(company)
	if createErr != nil {
		return nil, createErr
	}

	findByCode, err = c.rep.FindByCode(company.Code)
	if err != nil {
		return nil, err
	}

	return &dto.CompanyResponse{
		Id:             findByCode.Id,
		Code:           findByCode.Code,
		Name:           findByCode.Name,
		SiteUrl:        findByCode.SiteUrl,
		CreationDate:   time.Unix(findByCode.CreationDate, 0).String(),
		LastChangeDate: time.Unix(findByCode.LastChangeDate, 0).String(),
	}, nil

}
