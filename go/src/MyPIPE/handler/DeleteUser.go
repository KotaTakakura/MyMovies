package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteUser struct {
	DeleteUserUsecase usecase.IDeleteUser
}

func NewDeleteUser(deleteUserUsecase usecase.IDeleteUser) *DeleteUser {
	return &DeleteUser{
		DeleteUserUsecase: deleteUserUsecase,
	}
}

func (d DeleteUser) DeleteUser(c *gin.Context) {
	userIdUint := uint64((jwt.ExtractClaims(c)["id"]).(float64))
	userId, userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Validation Error.",
			"messages": userIdErr.Error(),
		})
		c.Abort()
		return
	}

	deleteUserDTO := usecase.NewDeleteUserDTO(userId)
	deleteUserErr := d.DeleteUserUsecase.DeleteUser(deleteUserDTO)

	if deleteUserErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Error.",
			"messages": deleteUserErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":   "Success.",
		"messages": "OK",
	})
}
