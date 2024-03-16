package validations

import "fmt"

var typeErr = map[string]string{
	"gte": "O atributo deve conter no minimo %s caracteres",
	"lte": "O atributo deve ter até no máximo %s",
}

func getErr(tag string, tagErr, param string) string {
	msg := fmt.Sprintf(typeErr[tag], param)

	if msg != "" {
		return msg
	}

	return tagErr
}
