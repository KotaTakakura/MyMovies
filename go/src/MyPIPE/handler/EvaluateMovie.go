package handler

import (
	"MyPIPE/infra"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EvaluateMovie(c *gin.Context){
	var evaluateMovieDTO usecase.EvaluateMovieDTO
	evaluateMovieDTOErr := c.Bind(&evaluateMovieDTO)
	userId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	evaluateMovieDTO.UserID = userId
	if evaluateMovieDTOErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Error.",
			"messages": evaluateMovieDTOErr.Error(),
		})
		c.Abort()
		return
	}

	userRepository := infra.NewUserPersistence()
	movieRepository := infra.NewMoviePersistence()
	evaluateMovieUsecase := usecase.NewEvaluateUsecase(userRepository,movieRepository)
	evaluateMovieUsecaseErr := evaluateMovieUsecase.EvaluateMovie(evaluateMovieDTO)
	if evaluateMovieUsecaseErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Error.",
			"messages": evaluateMovieUsecaseErr.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success.",
		"messages": "OK",
	})
}