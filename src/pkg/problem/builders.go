package problem

func NewProblem(problemType, title, detail, instance string, statusCode int) *Problem {
	return &Problem{
		Type:       problemType,
		Title:      title,
		Detail:     detail,
		Instance:   instance,
		StatusCode: statusCode,
	}
}

func NewProblemDetail(field, title, detail, instance, message string) Detail {
	return Detail{
		Field:    field,
		Title:    title,
		Detail:   detail,
		Instance: instance,
		Message:  message,
	}
}
