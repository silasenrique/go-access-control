package helper

import "go-access-control/src/internal/validations"

type BaseHttpResponse struct {
	Result          any                            `json:"result"`
	Success         bool                           `json:"success"`
	Code            int                            `json:"code"`
	ValidationError *[]validations.ValidationError `json:"validationErros,omitempty"`
	Err             any                            `json:"error,omitempty"`
}

func NewBaseHttpResponse(result any, success bool, code int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:  result,
		Success: success,
		Code:    code,
	}
}

func NewValidationError(result any, success bool, code int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:          result,
		Success:         success,
		Code:            code,
		ValidationError: validations.GetValidationErrors(err),
	}
}

func NewBussinessError(result any, success bool, code int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:  result,
		Success: success,
		Code:    code,
		Err:     err,
	}
}

func NewHttpError(success bool, code int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success: success,
		Code:    code,
		Err:     err,
	}
}

func NewDatabaseError(success bool, code int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success: success,
		Code:    code,
		Err:     err,
	}
}

func (a BaseHttpResponse) Error() string {
	return "nil"
}
