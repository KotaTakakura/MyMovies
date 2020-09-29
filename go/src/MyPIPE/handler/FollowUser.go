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

func FollowUser(c *gin.Context){
	var followUserJson FollowUserJson
	c.Bind(&followUserJson)
	userIdJWT :=  uint64(jwt.ExtractClaims(c)["id"].(float64))
	followUserJson.UserID = userIdJWT

	validationErrors := make(map[string]string)
	userId,userIdErr := model.NewUserID(followUserJson.UserID)
	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
	}

	followId,followIdErr := model.NewUserID(followUserJson.FollowID)
	if followIdErr != nil{
		validationErrors["follow_id"] = followIdErr.Error()
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

	followUserDTO := usecase.NewFollowDTO(userId,followId)
	userRepository := infra.NewUserPersistence()
	followUserRepository := infra.NewFollowUserPersistence()
	followUserUsecase := usecase.NewFollowUser(userRepository,followUserRepository)
	followUserUsecaseErr :=followUserUsecase.Follow(followUserDTO)
	if followUserUsecaseErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Error.",
			"messages": followUserUsecaseErr.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success.",
		"messages": "OK",
	})
}

type FollowUserJson struct{
	UserID uint64
	FollowID uint64 `json:"follow_id"`
}