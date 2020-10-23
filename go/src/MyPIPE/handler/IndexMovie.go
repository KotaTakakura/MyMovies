package handler

import (
	"MyPIPE/domain/queryService"
	"MyPIPE/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IndexMovie struct {
	IndexMovieQueryService queryService.IndexMovieQueryService
	IndexMovieUsecase      usecase.IIndexMovie
}

func NewIndexMovie(indexMovieQueryService queryService.IndexMovieQueryService, indexMovieUsecase usecase.IIndexMovie) *IndexMovie {
	return &IndexMovie{
		IndexMovieQueryService: indexMovieQueryService,
		IndexMovieUsecase:      indexMovieUsecase,
	}
}

func (indexMovie IndexMovie) IndexMovie(c *gin.Context) {
	keyWord := c.Query("keyWord")

	validationErrors := make(map[string]string)

	var page queryService.IndexMovieQueryServicePage
	pageInt, err := strconv.ParseUint(c.Query("page"), 10, 64)
	if err != nil {
		page,_ = queryService.NewIndexMovieQueryServicePage(1)
	}else{
		var pageErr error
		page,pageErr = queryService.NewIndexMovieQueryServicePage(uint(pageInt))
		if pageErr != nil{
			validationErrors["page"] = pageErr.Error()
		}
	}

	order, orderErr := queryService.NewIndexMovieQueryServiceOrder(c.Query("order"))
	if orderErr != nil {
		validationErrors["order"] = orderErr.Error()
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

	indexMovieSearchDTO := usecase.NewIndexMovieSearchDTO(page, keyWord, order)
	movies := indexMovie.IndexMovieUsecase.Search(indexMovieSearchDTO)

	jsonResult, jsonMarshalErr := json.Marshal(movies)
	if jsonMarshalErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Server Error.",
			"messages": jsonMarshalErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, string(jsonResult))
}

type IndexMovieJson struct {
	KeyWord string
}
