package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"github.com/gin-gonic/gin"
)

func TemporaryRegisterUser(c *gin.Context) {
	userPersistence := infra.NewUserPersistence()
	userRegistration := usecase.NewUserTemporaryRegistration(userPersistence)
	var user model.User
	c.Bind(&user)
	userRegistration.TemporaryRegister(&user)
}
