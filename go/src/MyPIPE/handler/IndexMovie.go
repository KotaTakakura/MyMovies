package handler

import (
	"MyPIPE/domain/queryService"
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

	validationErrors := make(map[string]string)

	order,orderErr := queryService.NewIndexMovieQueryServiceOrder(c.Query("order"))
	if orderErr != nil{
		validationErrors["order"] = orderErr.Error()
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

	indexMovieQueryService := queryService_infra.NewIndexMovie()
	indexMovieUsecase := usecase.NewIndexMovie(indexMovieQueryService)
	indexMovieSearchDTO := usecase.NewIndexMovieSearchDTO(page,keyWord,order)
	movies := indexMovieUsecase.Search(indexMovieSearchDTO)

	jsonResult, jsonMarshalErr := json.Marshal(movies)
	if jsonMarshalErr != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Server Error.",
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