package problem

func NewProblem(problemType Type, title Title, detail Details, instance string) *Problem {
	return &Problem{
		Type:     problemType.String(),
		Title:    title.String(),
		Detail:   detail.String(),
		Instance: instance,
	}
}

func NewProblemError(problemType Type, title Title, detail Details, instance string, err string) *Problem {
	return &Problem{
		Type:     problemType.String(),
		Title:    title.String(),
		Detail:   detail.String(),
		Instance: instance,
		Err:      err,
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
