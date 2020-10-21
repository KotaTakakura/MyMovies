package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreatePlayList struct{
	UserRepository	repository.UserRepository
	PlayListRepository	repository.PlayListRepository
	CreatePlayListUsecase usecase.ICreatePlayList
}

func NewCreatePlayList(
	userRepository	repository.UserRepository,
	playListRepository	repository.PlayListRepository,
	createPlayListUsecase usecase.ICreatePlayList,
	)*CreatePlayList{
	return &CreatePlayList{
		UserRepository:        userRepository,
		PlayListRepository:    playListRepository,
		CreatePlayListUsecase: createPlayListUsecase,
	}
}

func (createPlayList CreatePlayList)CreatePlayList(c *gin.Context){
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

	var playListDescriptionErr error
	CreatePlayListDTO.PlayListDescription,playListDescriptionErr = model.NewPlayListDescription(playListJson.PlayListDescription)
	if playListDescriptionErr != nil{
		validationErrors["play_list_description"] = playListDescriptionErr.Error()
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
	
	createPlayListUsecaseErr := createPlayList.CreatePlayListUsecase.CreatePlayList(CreatePlayListDTO)
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
	PlayListDescription string `json:"play_list_description"`
}