package helper

import (
	"errors"
	"go-access-control/src/pkg/problem"
	"net/http"
)

func getError(err error) *problem.Problem {
	helper, ok := err.(*helper)
	if !ok {
		return problem.NewProblem(
			Internal,
			ErrInternal.Error(),
			ValidationInternalProblemDetail,
			urlHelpers[ErrInternal],
			http.StatusInternalServerError).AddError(err)
	}

	switch {
	case errors.Is(helper.sentinel, ErrValidationStruct):
		return newProblemValidation(helper.internal)
	case errors.Is(helper.sentinel, ErrInternalValidationStruct):
		return problem.NewProblem(
			Internal,
			helper.sentinel.Error(),
			ValidationInternalProblemDetail,
			urlHelpers[ErrInternalValidationStruct],
			http.StatusInternalServerError).
			AddError(helper.internal).
			SetStatusCode(http.StatusInternalServerError)
	case errors.Is(helper.sentinel, ErrSQLNotFound):
		return problem.NewProblem(
			Internal,
			helper.sentinel.Error(),
			ValidationInternalProblemDetail,
			urlHelpers[ErrSQLNotFound],
			http.StatusInternalServerError).
			AddError(helper.internal).
			SetStatusCode(http.StatusBadRequest)
	case errors.Is(helper.sentinel, ErrSQLCreateFailure):
		return problem.NewProblem(
			Internal,
			helper.sentinel.Error(),
			ValidationInternalProblemDetail,
			urlHelpers[ErrSQLCreateFailure],
			http.StatusInternalServerError).AddError(helper.internal).SetStatusCode(http.StatusBadRequest)
	case errors.Is(helper.sentinel, ErrCodeAlreadyExists):
		return problem.NewProblem(
			Business,
			helper.sentinel.Error(),
			ValidationInternalProblemDetail,
			urlHelpers[ErrCodeAlreadyExists],
			http.StatusBadRequest).SetStatusCode(http.StatusBadRequest)
	default:
		return problem.NewProblem(
			Internal,
			ErrInternal.Error(),
			ValidationInternalProblemDetail,
			urlHelpers[ErrInternal],
			http.StatusInternalServerError).
			AddError(helper.internal).
			SetStatusCode(http.StatusInternalServerError)
	}
}
