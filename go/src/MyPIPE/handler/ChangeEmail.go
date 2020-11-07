package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChangeEmail struct {
	ChangeEmailUsecase usecase.IChangeEmail
}

func NewChangeEmail(changeEmailUsecase usecase.IChangeEmail) *ChangeEmail {
	return &ChangeEmail{
		ChangeEmailUsecase: changeEmailUsecase,
	}
}

func (c ChangeEmail) ChangeEmail(ctx *gin.Context) {
	var changeEmailJson ChangeEmailJson
	ctx.Bind(&changeEmailJson)

	emailChangeToken, emailChangeTokenErr := model.NewUserEmailChangeToken(changeEmailJson.Token)
	if emailChangeTokenErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result":   "Validation Error.",
			"messages": "Invalid Token.",
		})
		ctx.Abort()
		return
	}

	emailChangeDTO := usecase.NewChangeEmailDTO(emailChangeToken)
	usecaseErr := c.ChangeEmailUsecase.ChangeEmail(emailChangeDTO)
	if usecaseErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Error.",
			"messages": usecaseErr.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result":   "Success.",
		"messages": "OK",
	})
}

type ChangeEmailJson struct {
	Token string `json:"email_change_token"`
}
