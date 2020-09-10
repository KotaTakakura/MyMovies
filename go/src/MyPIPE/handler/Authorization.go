package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TemporaryRegisterUser(c *gin.Context) {
	userPersistence := infra.NewUserPersistence()
	userRegistration := usecase.NewUserTemporaryRegistration(userPersistence)
	var user model.User
	c.Bind(&user)
	err := userRegistration.TemporaryRegister(&user)
	if err != nil{
		c.JSON(500, gin.H{"message": c.Error(err)})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Temporary Registered!",
	})
}

func RegisterUser(c *gin.Context) {

}
