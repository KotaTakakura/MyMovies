package handler

import (
	"MyPIPE/domain/queryService"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexPlayListsInMyPage struct{
	IndexPlayListsInMyPageQueryService	queryService.IndexPlayListsInMyPageQueryService
	IndexPlayListsInMyPageUsecase	usecase.IIndexPlayListsInMyPage
}

func NewIndexPlayListsInMyPage(indexPlayListsInMyPageQueryService queryService.IndexPlayListsInMyPageQueryService,indexPlayListsInMyPageUsecase usecase.IIndexPlayListsInMyPage)*IndexPlayListsInMyPage{
	return &IndexPlayListsInMyPage{
		IndexPlayListsInMyPageQueryService: indexPlayListsInMyPageQueryService,
		IndexPlayListsInMyPageUsecase: indexPlayListsInMyPageUsecase,
	}
}

func (indexPlayListsInMyPage IndexPlayListsInMyPage)IndexPlayListsInMyPage(c *gin.Context){
	userId :=  uint64(jwt.ExtractClaims(c)["id"].(float64))
	playLists := indexPlayListsInMyPage.IndexPlayListsInMyPageUsecase.All(userId)

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
