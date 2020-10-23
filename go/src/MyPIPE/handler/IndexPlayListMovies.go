package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IndexPlaylistMovies struct {
	IndexPlaylistMoviesQueryService queryService.IndexPlayListMovieQueryService
	IndexPlaylistMoviesUsecase      usecase.IIndexPlaylistItemInMyPage
}

func NewIndexPlaylistMovies(indexPlaylistMoviesQueryService queryService.IndexPlayListMovieQueryService, indexPlaylistMoviesUsecase usecase.IIndexPlaylistItemInMyPage) *IndexPlaylistMovies {
	return &IndexPlaylistMovies{
		IndexPlaylistMoviesQueryService: indexPlaylistMoviesQueryService,
		IndexPlaylistMoviesUsecase:      indexPlaylistMoviesUsecase,
	}
}

func (indexPlaylistMovies IndexPlaylistMovies) IndexPlaylistMovies(c *gin.Context) {
	userIdInt := uint64(jwt.ExtractClaims(c)["id"].(float64))
	playListIdInt, _ := strconv.ParseUint(c.Param("play_list_id"), 10, 64)

	validationErrors := make(map[string]string)

	userId, userIdErr := model.NewUserID(userIdInt)
	if userIdErr != nil {
		validationErrors["user_id"] = userIdErr.Error()
	}

	playListId, playListIdErr := model.NewPlayListID(playListIdInt)
	if playListIdErr != nil {
		validationErrors["play_list_id"] = playListIdErr.Error()
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

	indexPlayListItemUsecaseDTO := usecase.NewIndexPlayListItemInMyPageDTO(userId, playListId)
	result := indexPlaylistMovies.IndexPlaylistMoviesUsecase.Find(indexPlayListItemUsecaseDTO)

	jsonResult, jsonMarshalErr := json.Marshal(result)
	if jsonMarshalErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Server Error.",
			"messages": jsonMarshalErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, string(jsonResult))
}

type IndexPlayListMoviesJson struct {
	PlayListID uint64
}
