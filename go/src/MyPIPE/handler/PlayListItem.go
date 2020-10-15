package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"MyPIPE/infra/factory"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddPlayListMovie(c *gin.Context){
	var playListItemAddJson AddPlayListItemJson
	c.Bind(&playListItemAddJson)
	userId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	playListItemAddJson.UserID = userId

	validationErrors := make(map[string]string)
	var playListItemAddDTO usecase.AddPlayListItemAddJson
	var playListIDErr error
	playListItemAddDTO.PlayListID,playListIDErr = model.NewPlayListID(playListItemAddJson.PlayListID)
	if playListIDErr != nil{
		validationErrors["play_list_id"] = playListIDErr.Error()
	}

	var UserIDErr error
	playListItemAddDTO.UserID,UserIDErr = model.NewUserID(playListItemAddJson.UserID)
	if UserIDErr != nil{
		validationErrors["user_id"] = UserIDErr.Error()
	}

	var MovieIDErr error
	playListItemAddDTO.MovieID,MovieIDErr = model.NewMovieID(playListItemAddJson.MovieID)
	if MovieIDErr != nil{
		validationErrors["movie_id"] = MovieIDErr.Error()
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

	playListRepository := infra.NewPlayListPersistence()
	playListMovieRepository := infra.NewPlayListMoviePersistence()
	playListMovieFactory := factory.NewPlayListMovieFactory()
	playListItemAddUsecase := usecase.NewAddPlayListItem(playListRepository,playListMovieRepository,playListMovieFactory)
	err := playListItemAddUsecase.AddPlayListItem(playListItemAddDTO)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
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

type AddPlayListItemJson struct{
	PlayListID uint64 `json:"play_list_id"`
	UserID uint64 `json:"user_id"`
	MovieID uint64 `json:"movie_id"`
}

func DeletePlayListMovie(c *gin.Context){
	var playListItemDeleteJson DeletePlayListItemJson
	c.Bind(&playListItemDeleteJson)

	validationErrors := make(map[string]string)
	playListIdFromQuery,_ :=strconv.ParseUint(c.Query("play_list_id"),10,64)
	playListID,playListIDErr := model.NewPlayListID(playListIdFromQuery)
	if playListIDErr != nil{
		validationErrors["play_list_id"] = playListIDErr.Error()
	}

	userId,UserIDErr := model.NewUserID(uint64(jwt.ExtractClaims(c)["id"].(float64)))
	if UserIDErr != nil{
		validationErrors["user_id"] = UserIDErr.Error()
	}

	movieIdFromQuery,_ :=strconv.ParseUint(c.Query("movie_id"),10,64)
	movieId,MovieIDErr := model.NewMovieID(movieIdFromQuery)
	if MovieIDErr != nil{
		validationErrors["movie_id"] = MovieIDErr.Error()
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

	playListMovieDeleteDTO := usecase.NewDeletePlayListItemAddJson(playListID,userId,movieId)

	playListRepository := infra.NewPlayListPersistence()
	playListMovieRepository := infra.NewPlayListMoviePersistence()

	playListMovieDeleteUsecase := usecase.NewDeletePlayListMovie(playListRepository,playListMovieRepository)
	err := playListMovieDeleteUsecase.DeletePlayListItem(playListMovieDeleteDTO)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
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

type DeletePlayListItemJson struct{
	PlayListID uint64 `json:"play_list_id"`
	UserID uint64 `json:"user_id"`
	MovieID uint64 `json:"movie_id"`
}