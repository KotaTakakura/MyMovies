package handler

import (
	"MyPIPE/domain/factory"
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PlayList struct {
	PlayListRepository         repository.PlayListRepository
	PlayListMovieRepository    repository.PlayListMovieRepository
	PlayListMovieFactory       factory.IPlayListMovie
	AddPlayListItemUsecase     usecase.IAddPlayListItem
	DeletePlayListMovieUsecase usecase.IDeletePlayListMovie
}

func NewPlayListItem(
	playListRepository repository.PlayListRepository,
	playListMovieRepository repository.PlayListMovieRepository,
	playListMovieFactory factory.IPlayListMovie,
	addPlayListItemUsecase usecase.IAddPlayListItem,
	deletePlayListMovieUsecase usecase.IDeletePlayListMovie,
) *PlayList {
	return &PlayList{
		PlayListRepository:         playListRepository,
		PlayListMovieRepository:    playListMovieRepository,
		PlayListMovieFactory:       playListMovieFactory,
		AddPlayListItemUsecase:     addPlayListItemUsecase,
		DeletePlayListMovieUsecase: deletePlayListMovieUsecase,
	}
}

func (playList PlayList) AddPlayListMovie(c *gin.Context) {
	var playListItemAddJson AddPlayListItemJson
	c.Bind(&playListItemAddJson)
	userId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	playListItemAddJson.UserID = userId

	validationErrors := make(map[string]string)
	var playListItemAddDTO usecase.AddPlayListItemAddJson
	var playListIDErr error
	playListItemAddDTO.PlayListID, playListIDErr = model.NewPlayListID(playListItemAddJson.PlayListID)
	if playListIDErr != nil {
		validationErrors["play_list_id"] = playListIDErr.Error()
	}

	var UserIDErr error
	playListItemAddDTO.UserID, UserIDErr = model.NewUserID(playListItemAddJson.UserID)
	if UserIDErr != nil {
		validationErrors["user_id"] = UserIDErr.Error()
	}

	var MovieIDErr error
	playListItemAddDTO.MovieID, MovieIDErr = model.NewMovieID(playListItemAddJson.MovieID)
	if MovieIDErr != nil {
		validationErrors["movie_id"] = MovieIDErr.Error()
	}

	if len(validationErrors) != 0 {
		validationErrors, _ := json.Marshal(validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Validation Error.",
			"messages": string(validationErrors),
		})
		c.Abort()
		return
	}

	err := playList.AddPlayListItemUsecase.AddPlayListItem(&playListItemAddDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Error.",
			"messages": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":   "Success.",
		"messages": "OK",
	})
}

type AddPlayListItemJson struct {
	PlayListID uint64 `json:"play_list_id"`
	UserID     uint64 `json:"user_id"`
	MovieID    uint64 `json:"movie_id"`
}

func (playList PlayList) DeletePlayListMovie(c *gin.Context) {

	validationErrors := make(map[string]string)
	playListIdFromQuery, _ := strconv.ParseUint(c.Query("play_list_id"), 10, 64)
	playListID, playListIDErr := model.NewPlayListID(playListIdFromQuery)
	if playListIDErr != nil {
		validationErrors["play_list_id"] = playListIDErr.Error()
	}

	userId, UserIDErr := model.NewUserID(uint64(jwt.ExtractClaims(c)["id"].(float64)))
	if UserIDErr != nil {
		validationErrors["user_id"] = UserIDErr.Error()
	}

	movieIdFromQuery, _ := strconv.ParseUint(c.Query("movie_id"), 10, 64)
	movieId, MovieIDErr := model.NewMovieID(movieIdFromQuery)
	if MovieIDErr != nil {
		validationErrors["movie_id"] = MovieIDErr.Error()
	}

	if len(validationErrors) != 0 {
		validationErrors, _ := json.Marshal(validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Validation Error.",
			"messages": string(validationErrors),
		})
		c.Abort()
		return
	}

	playListMovieDeleteDTO := usecase.NewDeletePlayListMovieJson(playListID, userId, movieId)

	err := playList.DeletePlayListMovieUsecase.DeletePlayListItem(playListMovieDeleteDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Error.",
			"messages": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"result":   "Success.",
		"messages": "OK",
	})
}

type DeletePlayListItemJson struct {
	PlayListID uint64 `json:"play_list_id"`
	UserID     uint64 `json:"user_id"`
	MovieID    uint64 `json:"movie_id"`
}
