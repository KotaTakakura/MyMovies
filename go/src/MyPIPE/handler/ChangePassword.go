package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangePassword(c *gin.Context){
	userIdUint := uint64(jwt.ExtractClaims(c)["id"].(float64))
	validationErrors := make(map[string]string)

	userId,userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
	}

	var changePasswordJson ChangePasswordJson
	c.Bind(&changePasswordJson)
	userPassword,userPasswordErr := model.NewUserPassword(changePasswordJson.Password)
	if userPasswordErr != nil{
		validationErrors["user_name"] = userPasswordErr.Error()
	}

	if len(validationErrors) != 0{
		validationErrors,_ :=  json.Marshal(validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error.",
			"messages": string(validationErrors),
		})
		c.Abort()
		return
	}

	changePasswordDTO := usecase.NewChangePasswordDTO(userId,userPassword)
	userRepository := infra.NewUserPersistence()
	changePasswordUsecase := usecase.NewChangePassword(userRepository)
	err := changePasswordUsecase.ChangePassword(changePasswordDTO)

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

type ChangePasswordJson struct{
	Password string	`json:"password"`
}