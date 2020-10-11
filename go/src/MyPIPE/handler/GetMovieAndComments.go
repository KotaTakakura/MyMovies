package handler

import (
	"MyPIPE/domain/model"
	queryService_infra "MyPIPE/infra/queryService"
	"MyPIPE/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMovieAndComments(c *gin.Context){
	var getCommentsJson GetCommentsJson
	c.Bind(&getCommentsJson)

	validationErrors := make(map[string]string)
	movieIdInt,_ := strconv.ParseUint(c.Query("movie_id"), 10, 64)
	movieId,movieIdErr := model.NewMovieID(movieIdInt)
	if movieIdErr != nil{
		validationErrors["movie_id"] = movieIdErr.Error()
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

	getCommentsDTO := usecase.NewGetMovieAndCommentsDTO(movieId)
	commentQueryService := queryService_infra.NewCommentQueryService()
	getCommentsUsecase := usecase.NewGetMovieAndComments(commentQueryService)
	comments := getCommentsUsecase.Get(*getCommentsDTO)

	jsonResult, jsonMarshalErr := json.Marshal(comments)
	if jsonMarshalErr != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Validation Error.",
			"messages": jsonMarshalErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, string(jsonResult))
}

type GetCommentsJson struct{
	MovieID uint64
}