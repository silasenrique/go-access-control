package problem

type Problem struct {
	StatusCode int      `json:"-"`
	Type       string   `json:"type"`
	Title      string   `json:"title"`
	Detail     string   `json:"detail"`
	Instance   string   `json:"instance"`
	Err        string   `json:"error,omitempty"`
	Errors     []Detail `json:"errors,omitempty"`
}

func (p Problem) Error() string {
	return ""
}

func (p *Problem) AddDetails(detail Detail) {
	p.Errors = append(p.Errors, detail)
}

func (p *Problem) AddError(err error) *Problem {
	p.Err = err.Error()

	return p
}

func (p *Problem) SetStatusCode(status int) *Problem {
	p.StatusCode = status

	return p
}
