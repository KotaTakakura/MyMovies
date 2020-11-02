package demo

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
	emailStruct := model.NewTemporaryRegisterMail(model.UserEmail("complaint@simulator.amazonses.com"), model.UserToken("123-456-789"))
	email := infra.NewTemporaryRegisterMailRepository()
	err := email.Send(emailStruct)
	if err != nil {
		fmt.Println(err.Error())
	}
}
