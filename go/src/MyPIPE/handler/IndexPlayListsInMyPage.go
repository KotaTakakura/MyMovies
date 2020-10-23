package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexPlayListsInMyPage struct {
	IndexPlayListsInMyPageQueryService queryService.IndexPlayListsInMyPageQueryService
	IndexPlayListsInMyPageUsecase      usecase.IIndexPlayListsInMyPage
}

func NewIndexPlayListsInMyPage(indexPlayListsInMyPageQueryService queryService.IndexPlayListsInMyPageQueryService, indexPlayListsInMyPageUsecase usecase.IIndexPlayListsInMyPage) *IndexPlayListsInMyPage {
	return &IndexPlayListsInMyPage{
		IndexPlayListsInMyPageQueryService: indexPlayListsInMyPageQueryService,
		IndexPlayListsInMyPageUsecase:      indexPlayListsInMyPageUsecase,
	}
}

func (indexPlayListsInMyPage IndexPlayListsInMyPage) IndexPlayListsInMyPage(c *gin.Context) {
	userIdInt := uint64(jwt.ExtractClaims(c)["id"].(float64))

	validationErrors := make(map[string]string)

	userId, userIdErr := model.NewUserID(userIdInt)
	if userIdErr != nil {
		validationErrors["user_id"] = userIdErr.Error()
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

	indexPlayListsInMyPageDTO := usecase.NewIndexPlayListsInMyPageDTO(userId)
	playLists := indexPlayListsInMyPage.IndexPlayListsInMyPageUsecase.All(indexPlayListsInMyPageDTO)

	jsonResult, jsonMarshalErr := json.Marshal(playLists)
	if jsonMarshalErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Internal Server Error.",
			"messages": jsonMarshalErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, string(jsonResult))
}
