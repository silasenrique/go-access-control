package handler

import (
	"database/sql"
	"encoding/json"
	"go-access-control/src/internal/api/dto"
	"go-access-control/src/internal/services"
	"go-access-control/src/pkg/problem"
	"net/http"
)

func CompanyCreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company := new(dto.CompanyCreateRequest)

		err := json.NewDecoder(r.Body).Decode(company)
		if err != nil {
			problem.ProblemHttpResponse(w, err)
			return
		}

		response, err := services.NewCompanyService(db).Create(company)
		if err != nil {
			problem.ProblemHttpResponse(w, err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
