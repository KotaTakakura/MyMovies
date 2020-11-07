package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResetPassword struct {
	ResetPasswordUsecase usecase.IResetPassword
}

func NewResetPassword(resetPasswordUsecase usecase.IResetPassword) *ResetPassword {
	return &ResetPassword{
		ResetPasswordUsecase: resetPasswordUsecase,
	}
}

func (r ResetPassword) ResetPassword(c *gin.Context) {
	var resetPasswordJson ResetPasswordJson
	c.Bind(&resetPasswordJson)
	validationErrors := make(map[string]string)

	passwordRememberToken, tokenErr := model.NewUserPasswordRememberToken(resetPasswordJson.PasswordRememberToken)
	if tokenErr != nil {
		validationErrors["password_remember_token"] = tokenErr.Error()
	}

	password, passwordErr := model.NewUserPassword(resetPasswordJson.Password)
	if passwordErr != nil {
		validationErrors["password"] = passwordErr.Error()
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

	resetPasswordDTO := usecase.NewResetPasswordDTO(passwordRememberToken, password)
	resetPasswordErr := r.ResetPasswordUsecase.ResetPassword(resetPasswordDTO)
	if resetPasswordErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Reset Password Error.",
			"messages": resetPasswordErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":   "Success.",
		"messages": "OK",
	})
}

type ResetPasswordJson struct {
	PasswordRememberToken string `json:"password_remember_token"`
	Password              string `json:"password"`
}
