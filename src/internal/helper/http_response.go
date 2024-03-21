package helper

import (
	"encoding/json"
	"errors"
	"go-access-control/src/pkg/problem"
	"net/http"
)

func ProblemHttpResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusBadRequest)

	problem := getError(err)
	json.NewEncoder(w).Encode(&problem)
}

func getError(err error) *problem.Problem {
	switch errors.Unwrap(err) {
	case ErrValidation:
		return NewProblemValidation(err)
	default:
		return problem.NewProblemError(Internal, ValidationInternalProblemTitle, ValidationInternalProblemDetail, "deu ruim", err.Error())
	}
}
