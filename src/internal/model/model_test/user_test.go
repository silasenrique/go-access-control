package modeltest

import (
	"go-access-control/src/internal/model"
	"reflect"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	user := &model.User{
		CompanyCode:    "XPTO",
		Email:          "xpto@xpto.com",
		CompleteName:   "xpto ltda.",
		CreationDate:   time.Now().Unix(),
		LastChangeDate: time.Now().Unix(),
	}

	newUser := model.NewUser(user.CompanyCode, user.CompleteName, user.Email)

	if !reflect.DeepEqual(user, newUser) {
		t.Errorf("as estruturas não são iguais")
	}
}

func TestValidateUser(t *testing.T) {
	t.Run("validate email", func(t *testing.T) {
		cases := []struct {
			*model.User
		}{
			{model.NewUser("xpto", "xpto ltda.", "xpto@pto.com")},
			{model.NewUser("xpto", "xpto ltda.", "")},
			{model.NewUser("xpto", "xpto ltda.", "xpto")},
		}

		for _, user := range cases {
			err := user.Validate()
		}
	})

}
