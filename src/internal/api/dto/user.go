package dto

import "go-access-control/src/internal/model"

type UserCreateRequest struct {
	CompanyCode      string `json:"companyCode,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	ResetFirstAccess bool   `json:"resetFirstAccess,omitempty"`
	Expiration       int64  `json:"expiration,omitempty"`
	CompleteName     string `json:"completeName,omitempty"`
}

type UserCreateResponse struct {
	Id              int    `json:"id"`
	CompanyCode     string `json:"companyCode"`
	Email           string `json:"email"`
	CompleteName    string `json:"completeName"`
	Password        string `json:"password"`
	Reset           bool   `json:"resetPassword"`
	Expiration      int64  `json:"expirationDate"`
	LatestResetDate int64  `json:"latestResetDate"`
	CreationDate    int64  `json:"creationDate"`
	LastChangeDate  int64  `json:"lastChangeDate"`
}

func NewUserCreateResponse(user *model.User) *UserCreateResponse {
	return &UserCreateResponse{
		Id:              user.Id,
		CompanyCode:     user.CompanyCode,
		Email:           user.Email,
		CompleteName:    user.CompleteName,
		Password:        user.Password.Password,
		Reset:           user.Password.Reset,
		Expiration:      user.Password.Expiration,
		LatestResetDate: user.Password.LatestResetDate,
		CreationDate:    user.CreationDate,
		LastChangeDate:  user.LastChangeDate,
	}
}
