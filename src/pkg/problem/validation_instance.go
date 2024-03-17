package problem

var instanceUrl = map[string]string{
	"required": "https://github.com/silasenrique/api-heper/required",
	"gte":      "https://github.com/silasenrique/api-heper/gte",
}

func getInstanceUrl(instance string) string {
	return instanceUrl[instance]
}
