package handler

import (
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetRandomUser(c *gin.Context){
	userPersistence := infra.NewUserPersistence()
	randomUser := usecase.NewRandomUser(userPersistence)
	user := randomUser.GetRandomIdUser()
	fmt.Print(user.Birthday)
}