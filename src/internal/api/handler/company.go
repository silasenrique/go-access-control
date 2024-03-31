package handler

import (
	"database/sql"
	"encoding/json"
	"go-access-control/src/internal/api/dto"
	"go-access-control/src/internal/helper"
	"go-access-control/src/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CompanyCreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company := new(dto.CompanyCreateRequest)

		err := json.NewDecoder(r.Body).Decode(company)
		if err != nil {
			helper.ProblemHttpResponse(w, err)
			return
		}

		response, err := services.NewCompanyService(db).Create(company)
		if err != nil {
			helper.ProblemHttpResponse(w, err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func CompatyGetIdHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		response, err := services.NewCompanyService(db).GetById(id)
		if err != nil {
			helper.ProblemHttpResponse(w, err)
			return
		}

		if response == nil {
			http.NotFound(w, r)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func CompatyGetByCodeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		response, err := services.NewCompanyService(db).GetByCode(code)
		if err != nil {
			helper.ProblemHttpResponse(w, err)
			return
		}

		if response == nil {
			http.NotFound(w, r)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
