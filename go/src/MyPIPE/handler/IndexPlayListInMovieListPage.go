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

func IndexPlayListInMovieListPage(c *gin.Context){
	userId,userIdErr :=  model.NewUserID(uint64(jwt.ExtractClaims(c)["id"].(float64)))

	validationErrors := make(map[string]string)

	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
	}

	movieIdUint64,_ := strconv.ParseUint(c.Param("movie_id"), 10, 64)
	movieId,movieIdErr := model.NewMovieID(movieIdUint64)

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
	indexPlayListInMovieListPageUsecase := usecase.NewIndexPlayListInMovieListPage()
	findDTO := usecase.NewFindDTO(userId,movieId)
	result := indexPlayListInMovieListPageUsecase.Find(*findDTO)

	jsonResult, jsonMarshalErr := json.Marshal(result)
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