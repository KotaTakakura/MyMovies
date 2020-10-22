package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeletePlayList struct{
	PlayListRepository repository.PlayListRepository
	DeletePlayListUsecase usecase.IDeletePlayList
}

func NewDeletePlayList(playListRepo repository.PlayListRepository,deletePlayListUsecase usecase.IDeletePlayList)*DeletePlayList{
	return &DeletePlayList{
		PlayListRepository: playListRepo,
		DeletePlayListUsecase: deletePlayListUsecase,
	}
}

func (deletePlayList DeletePlayList)DeletePlayList(c *gin.Context){
	validationErrors := make(map[string]string)
	userIdUint := uint64((jwt.ExtractClaims(c)["id"]).(float64))
	userId,userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
	}

	var playListId model.PlayListID
	var playListIdErr error
	playListIdString := c.Query("play_list_id")
	playListIdUint,playListIdUintErr := strconv.ParseUint(playListIdString, 10, 64)
	if playListIdUintErr != nil{
		validationErrors["play_list_id"] = playListIdUintErr.Error()
	}else{
		playListId,playListIdErr = model.NewPlayListID(playListIdUint)
		if playListIdErr != nil{
			validationErrors["play_list_id"] = playListIdErr.Error()
		}
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

	deletePlayListDTO := usecase.NewDeletePlayListDTO(userId,playListId)
	result := deletePlayList.DeletePlayListUsecase.Delete(deletePlayListDTO)

	if result != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Error.",
			"messages": result.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success.",
		"messages": "OK",
	})
}