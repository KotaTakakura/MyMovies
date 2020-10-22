package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChangePassword struct {
	UsrRepository         repository.UserRepository
	ChangePasswordUsecase usecase.IChangePassword
}

func NewChangePassword(u repository.UserRepository, c usecase.IChangePassword) *ChangePassword {
	return &ChangePassword{
		UsrRepository:         u,
		ChangePasswordUsecase: c,
	}
}

func (changePassword ChangePassword) ChangePassword(c *gin.Context) {
	userIdUint := uint64(jwt.ExtractClaims(c)["id"].(float64))
	validationErrors := make(map[string]string)
	userId, userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil {
		validationErrors["user_id"] = userIdErr.Error()
	}

	var changePasswordJson ChangePasswordJson
	c.Bind(&changePasswordJson)
	userPassword, userPasswordErr := model.NewUserPassword(changePasswordJson.Password)
	if userPasswordErr != nil {
		validationErrors["password"] = userPasswordErr.Error()
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

	changePasswordDTO := usecase.NewChangePasswordDTO(userId, userPassword)
	err := changePassword.ChangePasswordUsecase.ChangePassword(changePasswordDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Error.",
			"messages": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":   "Success.",
		"messages": "OK",
	})

}

type ChangePasswordJson struct {
	Password string `json:"password"`
}
