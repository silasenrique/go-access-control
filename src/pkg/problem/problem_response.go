package problem

import (
	"encoding/json"
	"net/http"
)

func ProblemHttpResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusBadRequest)

	abc := err.(*Problem)
	json.NewEncoder(w).Encode(abc)
}
