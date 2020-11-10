package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SetEmailChangeToken struct {
	SetEmailChangeTokenUsecase usecase.ISetEmailChangeToken
}

func NewSetEmailChangeToken(setEmailChangeTokenUsecase usecase.ISetEmailChangeToken) *SetEmailChangeToken {
	return &SetEmailChangeToken{
		SetEmailChangeTokenUsecase: setEmailChangeTokenUsecase,
	}
}

func (s SetEmailChangeToken) SetEmailChangeToken(c *gin.Context) {
	userIdUint := uint64((jwt.ExtractClaims(c)["id"]).(float64))
	var setEmailChangeTokenJson SetEmailChangeTokenJson
	c.Bind(&setEmailChangeTokenJson)

	validationErrors := make(map[string]string)
	userId, userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil {
		validationErrors["user_id"] = userIdErr.Error()
	}

	email, emailErr := model.NewUserEmail(setEmailChangeTokenJson.Email)
	if emailErr != nil {
		validationErrors["email"] = emailErr.Error()
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

	if userIdUint == 9 || userIdUint == 10 || userIdUint == 11 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Validation Error.",
			"messages": "Cant Change Test User.",
		})
		c.Abort()
		return
	}

	setEmailChangeTokenDTO := usecase.NewSetEmailChangeTokenDTO(userId, email)
	usecaseErr := s.SetEmailChangeTokenUsecase.SetEmailChangeToken(setEmailChangeTokenDTO)
	if usecaseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Change Email Failed.",
			"messages": usecaseErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success!",
	})

}

type SetEmailChangeTokenJson struct {
	Email string `json:"email"`
}
