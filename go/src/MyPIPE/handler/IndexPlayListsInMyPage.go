package handler

import (
	queryService_infra "MyPIPE/infra/queryService"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPlayListsInMyPage(c *gin.Context){
	userId :=  uint64(jwt.ExtractClaims(c)["id"].(float64))
	indexPlayListsInMyPageQueryService := queryService_infra.NewIndexPlayListsInMyPage()
	indexPlayListsInMyPageUsecase := usecase.NewIndexPlayListsInMyPage(indexPlayListsInMyPageQueryService)
	playLists := indexPlayListsInMyPageUsecase.All(userId)

	jsonResult, jsonMarshalErr := json.Marshal(playLists)
	if jsonMarshalErr != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Internal Server Error.",
			"messages": jsonMarshalErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, string(jsonResult))
}
