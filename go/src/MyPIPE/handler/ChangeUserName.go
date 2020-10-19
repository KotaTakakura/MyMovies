package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func ChangeUserName(c *gin.Context){
	userIdUint := uint64(jwt.ExtractClaims(c)["id"].(float64))
	validationErrors := make(map[string]string)

	userId,userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
	}

	var changeUserNameJson ChangeUserNameJson
	c.Bind(&changeUserNameJson)
	userName,userNameErr := model.NewUserName(changeUserNameJson.UserName)
	if userNameErr != nil{
		validationErrors["user_name"] = userNameErr.Error()
	}

	fmt.Println(userId)
	fmt.Println(userName)
	if len(validationErrors) != 0{
		validationErrors,_ :=  json.Marshal(validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error.",
			"messages": string(validationErrors),
		})
		c.Abort()
		return
	}

	changeUserNameDTO := usecase.NewChangeUserNameDTO(userId,userName)
	userRepository := infra.NewUserPersistence()
	changeUserNameUsecase := usecase.NewChangeUserName(userRepository)
	err := changeUserNameUsecase.ChangeUserName(changeUserNameDTO)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Error.",
			"messages": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success.",
		"messages": "OK",
	})

}

type ChangeUserNameJson struct{
	UserName string	`json:"user_name"`
}