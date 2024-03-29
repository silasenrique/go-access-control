package helper

import (
	"encoding/json"
	"net/http"
)

func ProblemHttpResponse(w http.ResponseWriter, err error) {
	problem := getError(err)

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(problem.StatusCode)
	json.NewEncoder(w).Encode(&problem)
}
