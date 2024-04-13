package problem

import "encoding/json"

type Problem struct {
	statusCode int
	reason     string
	title      string
	detail     string
	instance   string
	err        error
	errors     any
}

func NewProblem(reason, title string) *Problem {
	return &Problem{
		reason: reason,
		title:  title,
	}
}

func (p *Problem) Error() string {
	if p.err != nil {
		return p.err.Error()
	}
	return ""
}

func (p *Problem) AddDetail(detail string) *Problem {
	p.detail = detail
	return p
}

func (p *Problem) SetStatusCode(status int) *Problem {
	p.statusCode = status
	return p
}

func (p *Problem) SetInstance(instance string) *Problem {
	p.instance = instance
	return p
}

func (p *Problem) AddError(err error) *Problem {
	p.err = err
	return p
}

func (p *Problem) AddValidationError(validationErr any) *Problem {
	p.errors = validationErr
	return p
}

func (p *Problem) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		StatusCode int    `json:"-"`
		Reason     string `json:"type"`
		Title      string `json:"title"`
		Detail     string `json:"detail,omitempty"`
		Instance   string `json:"instance,omitempty"`
		Err        string `json:"error,omitempty"`
		Errors     any    `json:"errors,omitempty"`
	}{
		StatusCode: p.statusCode,
		Reason:     p.reason,
		Title:      p.title,
		Detail:     p.detail,
		Instance:   p.instance,
		Err:        p.Error(),
		Errors:     p.errors,
	})
}
