package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SetPasswordRememberToken struct{
	SetPasswordRememberTokenUsecase usecase.ISetPasswordRememberToken
}


func NewSetPasswordRememberToken(setPasswordRememberTokenUsecase usecase.ISetPasswordRememberToken)*SetPasswordRememberToken{
	return &SetPasswordRememberToken{
		SetPasswordRememberTokenUsecase: setPasswordRememberTokenUsecase,
	}
}

func (r SetPasswordRememberToken)SetPasswordRememberToken(c *gin.Context){
	var resetPasswordJson ResetPasswordJson
	c.Bind(&resetPasswordJson)

	userEmail,userEmailErr := model.NewUserEmail(resetPasswordJson.Email)
	if userEmailErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Validation Error.",
			"messages": userEmailErr.Error(),
		})
		c.Abort()
		return
	}

	setPasswordRememberTokenDTO := usecase.NewSetPasswordRememberTokenDTO(userEmail)
	setPasswordRememberTokenErr := r.SetPasswordRememberTokenUsecase.SetPasswordRememberToken(setPasswordRememberTokenDTO)
	if setPasswordRememberTokenErr != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Reset Password Error.",
			"messages": setPasswordRememberTokenErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success!",
	})
}

type ResetPasswordJson struct{
	Email string	`json:"email"`
}