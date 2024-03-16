package handler

import (
	"database/sql"
	"encoding/json"
	"go-access-control/src/internal/api/dto"
	"go-access-control/src/internal/helper"
	"go-access-control/src/internal/services"
	"net/http"
)

func CompanyCreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company := new(dto.CompanyCreateRequest)
		err := json.NewDecoder(r.Body).Decode(company)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(helper.NewHttpError(false, 60001, err))
			return
		}

		response, err := services.NewCompanyService(db).Create(company)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(helper.NewBaseHttpResponse(err, true, 4001))
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(helper.NewBaseHttpResponse(response, true, 4001))
	}
}
