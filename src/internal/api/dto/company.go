package dto

type CompanyCreateRequest struct {
	Code    string `json:"code,omitempty" validate:"required,alphanum,gte=3,lte=10"`
	Name    string `json:"name,omitempty" validate:"required,alphanum,gte=10,lte=100"`
	SiteUrl string `json:"siteUrl,omitempty" validate:"alphanum,lte=200"`
}

type CompanyResponse struct {
	Id             int    `json:"id"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	SiteUrl        string `json:"siteUrl"`
	CreationDate   string `json:"creationDate"`
	LastChangeDate string `json:"lastChangeDate"`
}
