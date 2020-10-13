package handler

import (
	queryService_infra "MyPIPE/infra/queryService"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexPlaylistMovies(c *gin.Context){
	userId :=  uint64(jwt.ExtractClaims(c)["id"].(float64))
	playListId,_ := strconv.ParseUint(c.Param("play_list_id"), 10, 64)

	indexPlayListItemUsecaseDTO := usecase.NewIndexPlayListItemInMyPageDTO(userId,playListId)
	NewIndexPlayListItemQueryService := queryService_infra.NewIndexPlayListMovieInMyPage()
	indexPlayListItemUsecaseUsecase := usecase.NewIndexPlayListItemInMyPage(NewIndexPlayListItemQueryService)
	result := indexPlayListItemUsecaseUsecase.Find(*indexPlayListItemUsecaseDTO)

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

type IndexPlayListMoviesJson struct{
	PlayListID uint64
}