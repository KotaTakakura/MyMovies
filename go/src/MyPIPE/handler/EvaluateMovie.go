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

type EvaluateMovie struct{
	MovieRepository repository.MovieRepository
	MovieEvaluationRepository	repository.MovieEvaluationRepository
	EvaluateMovieUsecase usecase.IEvaluateMovie
}

func NewEvaluateMovie(movieRepo repository.MovieRepository,movieEvaluateRepo repository.MovieEvaluationRepository,evaluateMovieUsecase usecase.IEvaluateMovie)*EvaluateMovie{
	return &EvaluateMovie{
		MovieRepository: movieRepo,
		MovieEvaluationRepository: movieEvaluateRepo,
		EvaluateMovieUsecase: evaluateMovieUsecase,
	}
}

func (evaluateMovie EvaluateMovie)EvaluateMovie(c *gin.Context){
	var evaluateMovieJson EvaluateMovieJson
	evaluateMovieJsonErr := c.Bind(&evaluateMovieJson)
	if evaluateMovieJsonErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Error.",
			"messages": evaluateMovieJsonErr.Error(),
		})
		c.Abort()
		return
	}
	userId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	var evaluateMovieDTO usecase.EvaluateMovieDTO
	validationErrors := make(map[string]string)
	var userIdErr error
	evaluateMovieDTO.UserID,userIdErr = model.NewUserID(userId)
	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
	}

	var movieIdErr error
	evaluateMovieDTO.MovieID,movieIdErr = model.NewMovieID(evaluateMovieJson.MovieID)
	if movieIdErr != nil{
		validationErrors["movie_id"] = movieIdErr.Error()
	}

	var evaluationErr error
	evaluateMovieDTO.Evaluation,evaluationErr = model.NewEvaluation(evaluateMovieJson.Evaluation)
	if evaluationErr != nil{
		validationErrors["evaluation"] = evaluationErr.Error()
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
	
	evaluateMovieUsecaseErr := evaluateMovie.EvaluateMovieUsecase.EvaluateMovie(&evaluateMovieDTO)
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

type EvaluateMovieJson struct{
	MovieID uint64	`json:"movie_id"`
	Evaluation string `json:"evaluate"`
}