package api

import (
	"database/sql"
	"go-access-control/src/internal/api/handler"
	"go-access-control/src/internal/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type api struct {
	api  *chi.Mux
	port string
	db   *sql.DB
}

func NewServer(port string, dbConnection string) *api {
	return &api{api: chi.NewMux(), port: port, db: database.NewDB(dbConnection)}
}

func (a *api) LoadRoutes() *api {
	a.api.Post("/company/create", handler.CompanyCreateHandler(a.db))
	a.api.Get("/company/get-by-code/{code}", handler.CompatyGetByCodeHandler(a.db))
	a.api.Get("/company/get-by-id/{id}", handler.CompatyGetIdHandler(a.db))

	return a
}

func (a *api) Run() error {
	return http.ListenAndServe(a.port, a.api)
}
