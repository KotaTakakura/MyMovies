package handler

import (
	"MyPIPE/domain/model"
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