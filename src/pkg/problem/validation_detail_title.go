package problem

var validationErrorTitle = map[string]string{
	"required": "Atributo obrigatório não informado",
	"gte":      "Atributo não atingiu a quantidade minima requerida",
}

func getValidationErrorTitle(errType string) string {
	return validationErrorTitle[errType]
}
