package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"MyPIPE/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CheckUserAlreadyLikedMovie struct{
	MovieEvaluationRepository repository.MovieEvaluationRepository
	CheckUserAlreadyLikedMovieUsecase usecase.ICheckUserAlreadyLikedMovie
}

func NewCheckUserAlreadyLikedMovie(movieEvaluationRepo repository.MovieEvaluationRepository, checkUserAlreadyLikedMovieUsecase usecase.ICheckUserAlreadyLikedMovie)*CheckUserAlreadyLikedMovie{
	return &CheckUserAlreadyLikedMovie{
		MovieEvaluationRepository: movieEvaluationRepo,
		CheckUserAlreadyLikedMovieUsecase: checkUserAlreadyLikedMovieUsecase,
	}
}

func (checkUserAlreadyLikedMovie CheckUserAlreadyLikedMovie)CheckUserAlreadyLikedMovie(c *gin.Context){
	var checkUserAlreadyLikedMovieJson CheckUserAlreadyLikedMovieJson
	c.Bind(&checkUserAlreadyLikedMovieJson)
	userIdString :=c.Query("user_id")
	movieIdString :=c.Query("movie_id")

	validationErrors := make(map[string]string)
	var userId model.UserID
	var userIdErr error
	var movieId model.MovieID
	var movieIdErr error

	userIdUint64,userIdUint64Err := strconv.ParseUint(userIdString, 10, 64)

	if userIdUint64Err != nil{
		validationErrors["user_id"] = userIdUint64Err.Error()
	}else{
		userId,userIdErr = model.NewUserID(userIdUint64)
		if userIdErr != nil{
			validationErrors["user_id"] = userIdErr.Error()
		}
	}

	movieIdUint64,movieIdUint64Err := strconv.ParseUint(movieIdString, 10, 64)
	if movieIdUint64Err != nil{
		validationErrors["movie_id"] = movieIdUint64Err.Error()
	}else{
		movieId,movieIdErr = model.NewMovieID(movieIdUint64)
		if movieIdErr != nil{
			validationErrors["movie_id"] = movieIdErr.Error()
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

	checkUserAlreadyLikedMovieFindDTO := usecase.NewCheckUserAlreadyLikedMovieFindDTO(userId,movieId)
	result := checkUserAlreadyLikedMovie.CheckUserAlreadyLikedMovieUsecase.Find(checkUserAlreadyLikedMovieFindDTO)

	if result {
		c.JSON(http.StatusOK, gin.H{
			"evaluated": "true",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"evaluated": "false",
		})
	}
}

type CheckUserAlreadyLikedMovieJson struct{
	UserID uint64
	MovieID uint64
}