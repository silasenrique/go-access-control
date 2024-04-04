package handler

import (
	"database/sql"
	"encoding/json"
	"go-access-control/src/internal/api/dto"
	"go-access-control/src/internal/helper"
	"go-access-control/src/internal/services"
	"net/http"
)

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := new(dto.UserCreateRequest)

		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			helper.ProblemHttpResponse(w, err)
			return
		}

		newUser, err := services.NewUserService(db).CreateUser(user)
		if err != nil {
			helper.ProblemHttpResponse(w, err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	}
}
