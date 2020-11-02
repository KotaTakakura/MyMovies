package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdatePlayList struct {
	UpdatePlayListUsecase usecase.IUpdatePlayList
}

func NewUpdatePlayListHandler(updatePlayListUsecase usecase.IUpdatePlayList) *UpdatePlayList {
	return &UpdatePlayList{
		UpdatePlayListUsecase: updatePlayListUsecase,
	}
}

func (u UpdatePlayList) Update(c *gin.Context) {
	var updatePlayListJson UpdatePlayListJson
	c.Bind(&updatePlayListJson)
	iuserId := uint64(jwt.ExtractClaims(c)["id"].(float64))

	validationErrors := make(map[string]string)

	userId, userIdErr := model.NewUserID(iuserId)
	if userIdErr != nil {
		validationErrors["user_id"] = userIdErr.Error()
	}

	playListId, playListIdErr := model.NewPlayListID(updatePlayListJson.PlayListID)
	if playListIdErr != nil {
		validationErrors["play_list_id"] = playListIdErr.Error()
	}

	playListName, playListNameErr := model.NewPlayListName(updatePlayListJson.PlayListName)
	if playListNameErr != nil {
		validationErrors["play_list_name"] = playListNameErr.Error()
	}

	playListDescription, playListDescriptionErr := model.NewPlayListDescription(updatePlayListJson.PlayListDescription)
	if playListDescriptionErr != nil {
		validationErrors["play_list_description"] = playListDescriptionErr.Error()
	}

	thumbnailMovieId, thumbnailMovieIdErr := model.NewPlayListThumbnailMovieID(updatePlayListJson.PlayListThumbnailMovieID)
	if thumbnailMovieIdErr != nil {
		validationErrors["thumbnail_movie_id"] = thumbnailMovieIdErr.Error()
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

	updatePlayListDTO := usecase.NewUpdatePlayListDTO(userId, playListId, playListName, playListDescription, thumbnailMovieId)
	usecaseError := u.UpdatePlayListUsecase.Update(updatePlayListDTO)

	if usecaseError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Update PlayList Failed.",
			"messages": usecaseError.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated!",
	})
}

type UpdatePlayListJson struct {
	PlayListID               uint64 `json:"play_list_id"`
	PlayListName             string `json:"play_list_name"`
	PlayListDescription      string `json:"play_list_description"`
	PlayListThumbnailMovieID uint64 `json:"play_list_thumbnail_movie_id"`
}
