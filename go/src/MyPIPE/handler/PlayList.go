package handler

import (
	"MyPIPE/infra"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePlayList(c *gin.Context){
	var playListDTO usecase.CreatePlayListJson
	bindErr := c.Bind(&playListDTO)
	if bindErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Error.",
			"messages": bindErr.Error(),
		})
		c.Abort()
		return
	}
	userId := uint64(jwt.ExtractClaims(c)["id"].(float64))
	playListDTO.UserID = userId

	userPersistence := infra.NewUserPersistence()
	playListPersistence := infra.NewPlayListPersistence()
	createPlayListUsecase := usecase.NewCreatePlayList(userPersistence,playListPersistence)
	createPlayListUsecaseErr := createPlayListUsecase.CreatePlayList(playListDTO)
	if createPlayListUsecaseErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Error.",
			"messages": createPlayListUsecaseErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success.",
		"messages": "OK",
	})
}