package handler

import (
	"MyPIPE/infra"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPlayList(c *gin.Context){
	var playListItemAddDTO usecase.PlayListItemAddJson
	c.Bind(&playListItemAddDTO)
	userId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	playListItemAddDTO.UserID = userId

	playListRepository := infra.NewPlayListPersistence()
	movieRepository := infra.NewMoviePersistence()
	playListItemAddUsecase := usecase.NewAddPlayListItem(playListRepository,movieRepository)
	err := playListItemAddUsecase.AddPlayListItem(playListItemAddDTO)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Error.",
			"messages": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success.",
		"messages": "OK",
	})
}
