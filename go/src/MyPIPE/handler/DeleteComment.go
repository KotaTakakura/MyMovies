package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteComment struct {
	DeleteCommentUsecase usecase.IDeleteComment
}

func NewDeleteComment(deleteCommentUsecase usecase.IDeleteComment) *DeleteComment {
	return &DeleteComment{
		DeleteCommentUsecase: deleteCommentUsecase,
	}
}

func (d DeleteComment) DeleteComment(c *gin.Context) {
	validationErrors := make(map[string]string)
	userIdUint := uint64((jwt.ExtractClaims(c)["id"]).(float64))
	userId, userIdErr := model.NewUserID(userIdUint)
	if userIdErr != nil {
		validationErrors["user_id"] = userIdErr.Error()
	}

	var commentId model.CommentID
	commentIdString := c.Query("comment_id")
	commentIdUint, commentIdUintErr := strconv.ParseUint(commentIdString, 10, 64)
	if commentIdUintErr != nil {
		validationErrors["comment_id"] = commentIdUintErr.Error()
	}
	commentId, commentIdErr := model.NewCommentID(commentIdUint)
	if commentIdErr != nil {
		validationErrors["comment_id"] = commentIdErr.Error()
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

	deleteCommentDTO := usecase.NewDeleteCommentDTO(commentId, userId)
	deleteCommentErr := d.DeleteCommentUsecase.DeleteComment(deleteCommentDTO)
	if deleteCommentErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Delete Comment Failed.",
			"messages": deleteCommentErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted!",
	})
}
