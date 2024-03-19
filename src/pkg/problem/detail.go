package problem

type Detail struct {
	Field    string `json:"field"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
	Message  string `json:"error"`
}
