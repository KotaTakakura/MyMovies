package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteMovie struct {
	DeleteMovieUsecase usecase.IDeleteMovie
}

func NewDeleteMovie(deleteMovieUsecase usecase.IDeleteMovie) *DeleteMovie {
	return &DeleteMovie{
		DeleteMovieUsecase: deleteMovieUsecase,
	}
}

func (d DeleteMovie) DeleteMovie(c *gin.Context) {

	validationErrors := make(map[string]string)

	userIdUint := uint64((jwt.ExtractClaims(c)["id"]).(float64))
	userId, userIdErr := model.NewUserID(userIdUint)

	if userIdErr != nil {
		validationErrors["user_id"] = userIdErr.Error()
	}

	movieIdString := c.Query("movie_id")
	movieIdUint, movieIdUintErr := strconv.ParseUint(movieIdString, 10, 64)
	if movieIdUintErr != nil {
		validationErrors["movie_id"] = movieIdUintErr.Error()
	}

	movieId, movieIdErr := model.NewMovieID(movieIdUint)
	if movieIdErr != nil {
		validationErrors["movie_id"] = movieIdErr.Error()
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

	deleteMovieDTO := usecase.NewDeleteMovieDTO(userId, movieId)
	deleteMovieErr := d.DeleteMovieUsecase.DeleteMovie(deleteMovieDTO)
	if deleteMovieErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Error.",
			"messages": deleteMovieErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":   "Success.",
		"messages": "OK",
	})
}