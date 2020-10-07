package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	queryService_infra "MyPIPE/infra/queryService"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUploadedMovies(c *gin.Context){
	userId := jwt.ExtractClaims(c)["id"]
	iuserId := uint64(userId.(float64))

	userIdModel,err := model.NewUserID(iuserId)
	if err != nil{
		return
	}

	uploadedMoviesQueryService := queryService_infra.NewUploadedMovies()
	uploadedMoviesUsecase := usecase.NewUploadedMovies(uploadedMoviesQueryService)
	result := uploadedMoviesUsecase.Get(userIdModel)
	jsonResult, jsonMarshalErr := json.Marshal(result)
	if jsonMarshalErr != nil {
		return
	}

	c.JSON(http.StatusOK, string(jsonResult))
}

func UpdateMovie(c *gin.Context){
	userId := jwt.ExtractClaims(c)["id"]
	iuserId := uint64(userId.(float64))

	userIdModel,err := model.NewUserID(iuserId)

	var updateMovieDTO UpdateMovieDTO
	c.Bind(&updateMovieDTO)

	validationErrors := make(map[string]string)

	if err != nil{
		validationErrors["user_id"] = err.Error()
	}

	movieId,movieIdErr := model.NewMovieID(updateMovieDTO.MovieID)
	if movieIdErr != nil{
		validationErrors["movie_id"] = movieIdErr.Error()
	}

	displayName,displayNameErr := model.NewMovieDisplayName(updateMovieDTO.DisplayName)
	if displayNameErr != nil{
		validationErrors["display_name"] = displayNameErr.Error()
	}

	description,descriptionErr := model.NewMovieDescription(updateMovieDTO.Description)
	if descriptionErr != nil{
		validationErrors["description"] = descriptionErr.Error()
	}

	public,publicErr := model.NewMoviePublic(updateMovieDTO.Public)
	if publicErr != nil{
		validationErrors["public"] = publicErr.Error()
	}

	status,statusErr := model.NewMovieStatus(updateMovieDTO.Status)
	if statusErr != nil{
		validationErrors["status"] = statusErr.Error()
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

	updateDTO := usecase.UpdateDTO{
		UserID:      userIdModel,
		MovieID:     movieId,
		DisplayName: displayName,
		Description: description,
		Public:      public,
		Status:      status,
	}

	movieRepository := infra.NewMoviePersistence()
	updateMovieUsecase := usecase.NewUpdateMovie(movieRepository)
	result,updateMovieUsecaseErr := updateMovieUsecase.Update(updateDTO)
	if updateMovieUsecaseErr != nil{
		jsonUpdateMovieUsecaseErr,_ := json.Marshal(updateMovieUsecaseErr.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Server Error.",
			"messages": string(jsonUpdateMovieUsecaseErr),
		})
		c.Abort()
		return
	}

	updatedData,_ := json.Marshal(result)
	c.JSON(http.StatusOK, gin.H{
		"result": "Success.",
		"messages": "OK",
		"updatedData": string(updatedData),
	})
}

type UpdateMovieDTO struct{
	UserID uint64
	MovieID uint64	`json:"movie_id"`
	DisplayName string	`json:"display_name"`
	Description string	`json:"description"`
	Public uint	`json:"public"`
	Status uint	`json:"status"`
}