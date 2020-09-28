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

func CreatePlayList(c *gin.Context){
	//var playListDTO usecase.CreatePlayListJson
	var playListJson CreatePlayListJson
	bindErr := c.Bind(&playListJson)
	if bindErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Error.",
			"messages": bindErr.Error(),
		})
		c.Abort()
		return
	}
	userId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	playListJson.UserID = userId

	validationErrors := make(map[string]string)
	var CreatePlayListDTO usecase.CreatePlayListDTO
	var userIDErr error
	CreatePlayListDTO.UserID,userIDErr = model.NewUserID(playListJson.UserID)
	if userIDErr != nil{
		validationErrors["user_id"] = userIDErr.Error()
	}

	var playListNameErr error
	CreatePlayListDTO.PlayListName,playListNameErr = model.NewPlayListName(playListJson.PlayListName)
	if playListNameErr != nil{
		validationErrors["play_list_name"] = playListNameErr.Error()
	}

	if len(validationErrors) != 0{
		validationErrors,_ := json.Marshal(validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error.",
			"messages": string(validationErrors),
		})
		c.Abort()
		return
	}

	userPersistence := infra.NewUserPersistence()
	playListPersistence := infra.NewPlayListPersistence()
	createPlayListUsecase := usecase.NewCreatePlayList(userPersistence,playListPersistence)
	createPlayListUsecaseErr := createPlayListUsecase.CreatePlayList(CreatePlayListDTO)
	if createPlayListUsecaseErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Error.",
			"messages": createPlayListUsecaseErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success.",
		"messages": "OK",
	})
}

type CreatePlayListJson struct{
	UserID uint64
	PlayListName string `json:"play_list_name"`
}