package services

import (
	"database/sql"
	"errors"
	"go-access-control/src/internal/api/dto"
	"go-access-control/src/internal/helper"
	"go-access-control/src/internal/model"
	"go-access-control/src/internal/repository"
	"strconv"
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
		return nil, helper.NewHelper(helper.ErrCodeAlreadyExists).AddIntenal(err)
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

func (c *CompanyService) GetById(id string) (*dto.CompanyResponse, error) {
	if id == "" {
		return nil, helper.NewHelper(helper.ErrTheIdCannotBeEmpty)
	}

	parseId, err := strconv.Atoi(id)
	if err != nil {
		return nil, helper.
			NewHelper(helper.ErrStringToIntegerConversion).
			AddIntenal(err)
	}

	company, err := c.rep.FindById(parseId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &dto.CompanyResponse{
		Id:             company.Id,
		Code:           company.Code,
		Name:           company.Name,
		SiteUrl:        company.SiteUrl,
		CreationDate:   time.Unix(company.CreationDate, 0).String(),
		LastChangeDate: time.Unix(company.LastChangeDate, 0).String(),
	}, nil
}

func (c *CompanyService) GetByCode(code string) (*dto.CompanyResponse, error) {
	if code == "" {
		return nil, helper.NewHelper(helper.ErrTheCodeCannotBeEmpty)
	}

	company, err := c.rep.FindByCode(code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &dto.CompanyResponse{
		Id:             company.Id,
		Code:           company.Code,
		Name:           company.Name,
		SiteUrl:        company.SiteUrl,
		CreationDate:   time.Unix(company.CreationDate, 0).String(),
		LastChangeDate: time.Unix(company.LastChangeDate, 0).String(),
	}, nil
}
