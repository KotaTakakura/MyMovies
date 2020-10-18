package handler

import (
	"MyPIPE/domain/model"
	queryService_infra "MyPIPE/infra/queryService"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLoggedInUserData(c *gin.Context){
	validationErrors := make(map[string]string)
	userIdUint := uint64((jwt.ExtractClaims(c)["id"]).(float64))
	userId,userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
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

	getLoggedInUserDataQueryService := queryService_infra.NewGetLoggedInUserData()
	getLoggedInUserDataUsecase := usecase.NewGetLoggedInUserData(getLoggedInUserDataQueryService)
	getLoggedInUserDataDTO := usecase.NewGetLoggedInUserDataDTO(userId)
	loggedInUser := getLoggedInUserDataUsecase.Find(getLoggedInUserDataDTO)

	c.JSON(http.StatusOK, loggedInUser)
}
