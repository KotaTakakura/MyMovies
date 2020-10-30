package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateMovieThumbnailStatus struct {
	UpdateMovieThumbnailStatusUsecase usecase.IUpdateMovie
}

func NewUpdateMovieThumbnailStatus(updateMovieThumbnailStatusUsecase usecase.IUpdateMovie) *UpdateMovieThumbnailStatus {
	return &UpdateMovieThumbnailStatus{
		UpdateMovieThumbnailStatusUsecase: updateMovieThumbnailStatusUsecase,
	}
}

func (u UpdateMovieThumbnailStatus) UpdateMovieThumbnailStatus(c *gin.Context) {
	var updateMovieThumbnailStatusJson UpdateMovieThumbnailStatusJson
	c.Bind(&updateMovieThumbnailStatusJson)

	movieId, movieIdErr := model.NewMovieID(updateMovieThumbnailStatusJson.MovieID)
	if movieIdErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Change Thumbnail Status Failed.",
			"messages": movieIdErr.Error(),
		})
		c.Abort()
		return
	}

	updateThumbnailStatusDTO := usecase.NewUpdateThumbnailStatusDTO(movieId)
	result := u.UpdateMovieThumbnailStatusUsecase.UpdateThumbnailStatus(updateThumbnailStatusDTO)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Change Thumbnail Status Failed.",
			"messages": result.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK!",
	})
}

type UpdateMovieThumbnailStatusJson struct {
	MovieID uint64 `json:"movie_id"`
}
