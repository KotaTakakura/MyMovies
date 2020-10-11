package handler

import (
	queryService_infra "MyPIPE/infra/queryService"
	"MyPIPE/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexMovie(c *gin.Context){
	keyWord := c.Query("keyWord")

	page,err := strconv.Atoi(c.Query("page"))
	if err != nil{
		page = 1
	}

	indexMovieQueryService := queryService_infra.NewIndexMovie()
	indexMovieUsecase := usecase.NewIndexMovie(indexMovieQueryService)
	movies := indexMovieUsecase.Search(page,keyWord)

	jsonResult, jsonMarshalErr := json.Marshal(movies)
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

type IndexMovieJson struct{
	KeyWord string
}