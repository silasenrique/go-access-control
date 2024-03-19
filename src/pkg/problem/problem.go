package problem

type Problem struct {
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Detail   string   `json:"detail"`
	Instance string   `json:"instance"`
	Err      string   `json:"error,omitempty"`
	Errors   []Detail `json:"errors,omitempty"`
}

func (p Problem) Error() string {
	return ""
}

func (p *Problem) AddDetails(detail Detail) {
	p.Errors = append(p.Errors, detail)
}
