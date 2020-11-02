package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateMovieStatus struct {
	UpdateMovieStatusUsecase usecase.IUpdateMovie
}

func NewUpdateMovieStatus(updateMovieStatusUsecase usecase.IUpdateMovie) *UpdateMovieStatus {
	return &UpdateMovieStatus{
		UpdateMovieStatusUsecase: updateMovieStatusUsecase,
	}
}

func (u UpdateMovieStatus) UpdateMovieStatus(c *gin.Context) {
	var updateMovieStatusJson UpdateMovieStatusJson
	c.Bind(&updateMovieStatusJson)
	fmt.Println("TEST")

	movieId, movieIdErr := model.NewMovieID(updateMovieStatusJson.MovieID)
	if movieIdErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Change Status Failed.",
			"messages": movieIdErr.Error(),
		})
		c.Abort()
		return
	}

	updateStatusDTO := usecase.NewUpdateStatusDTO(movieId)
	result := u.UpdateMovieStatusUsecase.UpdateStatus(updateStatusDTO)
	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Change Status Failed.",
			"messages": result.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK!",
	})
}

type UpdateMovieStatusJson struct {
	MovieID uint64 `json:"movie_id"`
}
