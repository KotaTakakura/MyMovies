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

type GetLoggedInUserData struct {
	GetLoggedInUserDataQueryService queryService.GetLoggedInUserDataQueryService
	GetLoggedInUserDataUsecase      usecase.IGetLoggedInUserData
}

func NewGetLoggedInUserData(GetLoggedInUserDataQueryService queryService.GetLoggedInUserDataQueryService, getLoggedInUserDataUsecase usecase.IGetLoggedInUserData) *GetLoggedInUserData {
	return &GetLoggedInUserData{
		GetLoggedInUserDataQueryService: GetLoggedInUserDataQueryService,
		GetLoggedInUserDataUsecase:      getLoggedInUserDataUsecase,
	}
}

func (getLoggedInUserData GetLoggedInUserData) GetLoggedInUserData(c *gin.Context) {
	validationErrors := make(map[string]string)
	userIdUint := uint64((jwt.ExtractClaims(c)["id"]).(float64))
	userId, userIdErr := model.NewUserID(userIdUint)
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

	getLoggedInUserDataDTO := usecase.NewGetLoggedInUserDataDTO(userId)
	loggedInUser := getLoggedInUserData.GetLoggedInUserDataUsecase.Find(getLoggedInUserDataDTO)

	c.JSON(http.StatusOK, loggedInUser)
}
