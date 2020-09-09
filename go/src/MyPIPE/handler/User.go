package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"github.com/gin-gonic/gin"
	"time"
)

func RegisterUser(c *gin.Context) {
	userPersistence := infra.NewUserPersistence()
	New := usecase.NewUser(userPersistence)

	var newUser model.User
	newUser.Name = "tatata"
	newUser.Email = "tatata@tatata.jp"
	newUser.Birthday = time.Now()
	newUser.SetHashedPassword("takakura")

	New.RegisterUser(&newUser)
}
