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
			InternalErrorDetail,
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
			SQLNotFoundProblemDetail,
			urlHelpers[ErrSQLNotFound],
			http.StatusInternalServerError).
			AddError(helper.internal).
			SetStatusCode(http.StatusBadRequest)
	case errors.Is(helper.sentinel, ErrSQLCreateFailure):
		return problem.NewProblem(
			Internal,
			helper.sentinel.Error(),
			SQLCreateProblemDetail,
			urlHelpers[ErrSQLCreateFailure],
			http.StatusInternalServerError).AddError(helper.internal).SetStatusCode(http.StatusBadRequest)
	case errors.Is(helper.sentinel, ErrCodeAlreadyExists):
		return problem.NewProblem(
			Business,
			helper.sentinel.Error(),
			CodeExistsDetails,
			urlHelpers[ErrCodeAlreadyExists],
			http.StatusBadRequest).SetStatusCode(http.StatusBadRequest)
	case errors.Is(helper.sentinel, ErrTheIdCannotBeEmpty):
		return problem.NewProblem(
			Internal,
			helper.sentinel.Error(),
			TheIdCannotBeEmpty,
			urlHelpers[ErrTheIdCannotBeEmpty],
			http.StatusInternalServerError).SetStatusCode(http.StatusBadRequest)
	case errors.Is(helper.sentinel, ErrTheCodeCannotBeEmpty):
		return problem.NewProblem(
			Internal,
			helper.sentinel.Error(),
			TheCodeCannotBeEmpty,
			urlHelpers[ErrTheCodeCannotBeEmpty],
			http.StatusInternalServerError).SetStatusCode(http.StatusBadRequest)
	case errors.Is(helper.sentinel, ErrStringToIntegerConversion):
		return problem.NewProblem(
			Internal,
			helper.sentinel.Error(),
			StringToIntegerConversion,
			urlHelpers[ErrStringToIntegerConversion],
			http.StatusInternalServerError).
			AddError(helper.internal).
			SetStatusCode(http.StatusBadRequest)
	default:
		return problem.NewProblem(
			Internal,
			ErrInternal.Error(),
			InternalErrorDetail,
			urlHelpers[ErrInternal],
			http.StatusInternalServerError).
			AddError(helper.internal).
			SetStatusCode(http.StatusInternalServerError)
	}
}
