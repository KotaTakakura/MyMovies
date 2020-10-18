package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	"encoding/json"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeOrderOfPlayListMovies(c *gin.Context){
	var changeOrderOfPlayListMoviesJson ChangeOrderOfPlayListMoviesJson
	c.Bind(&changeOrderOfPlayListMoviesJson)
	userIdUint :=  uint64(jwt.ExtractClaims(c)["id"].(float64))

	fmt.Println(changeOrderOfPlayListMoviesJson)
	validationErrors := make(map[string]string)
	var movieIdAndOrderForChangeOrderOfPlayListMoviesDTO usecase.MovieIdAndOrderForChangeOrderOfPlayListMoviesDTO
	var changeOrderOfPlayListMoviesDTO usecase.ChangeOrderOfPlayListMoviesDTO
	for _,value := range changeOrderOfPlayListMoviesJson.PlayListMovieIdAndOrder{
		var movieIdErr error
		movieIdAndOrderForChangeOrderOfPlayListMoviesDTO.MovieID,movieIdErr = model.NewMovieID(value.MovieID)
		if movieIdErr != nil{
			validationErrors["movie_id"] = movieIdErr.Error()
		}

		var OrderErr error
		movieIdAndOrderForChangeOrderOfPlayListMoviesDTO.Order,OrderErr = model.NewPlayListMovieOrder(value.Order)
		if OrderErr != nil{
			validationErrors["order"] = OrderErr.Error()
		}

		changeOrderOfPlayListMoviesDTO.MovieIDAndOrder = append(changeOrderOfPlayListMoviesDTO.MovieIDAndOrder,movieIdAndOrderForChangeOrderOfPlayListMoviesDTO)
	}

	userId,userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil{
		validationErrors["user_id"] = userIdErr.Error()
	}

	playListID,playListIDErr := model.NewPlayListID(changeOrderOfPlayListMoviesJson.PlayListID)
	if playListIDErr != nil{
		validationErrors["play_list_id"] = playListIDErr.Error()
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

	changeOrderOfPlayListMoviesDTO.UserID = userId
	changeOrderOfPlayListMoviesDTO.PlayListID = playListID

	playListMovieRepository := infra.NewPlayListMoviePersistence()
	changeOrderOfPlayListMovies := usecase.NewChangeOrderOfPlayListMovies(playListMovieRepository)
	result := changeOrderOfPlayListMovies.ChangeOrderOfPlayListMovies(&changeOrderOfPlayListMoviesDTO)
	if result != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Validation Error.",
			"messages": result.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"OK",
	})
}

type ChangeOrderOfPlayListMoviesJson struct{
	PlayListID uint64	`json:"play_list_id"`
	PlayListMovieIdAndOrder	[]PlayListMovieIdAndOrderForChangeOrderOfPlayListMoviesJson	`json:"play_list_movie_id_and_order"`
}

type PlayListMovieIdAndOrderForChangeOrderOfPlayListMoviesJson struct{
	MovieID uint64	`json:"play_list_movie_id"`
	Order	int	`json:"play_lise_movie_order"`
}