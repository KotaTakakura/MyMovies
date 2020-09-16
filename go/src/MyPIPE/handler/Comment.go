package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostComment(c *gin.Context) {
	userRepository := infra.NewUserPersistence()
	postCommentUsecase := usecase.NewPostComment(userRepository)
	userId := jwt.ExtractClaims(c)["id"]
	iuserId := uint64(userId.(float64))

	var comment PostCommentJson
	bindErr := c.Bind(&comment)
	if bindErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Server Error.",
			"messages": bindErr.Error(),
		})
		c.Abort()
		return
	}

	var newComment model.Comment
	validationErrors := map[string]error{}
	errorMessages := map[string]string{}
	validationErrorFlag := false


	newComment.Body,validationErrors["comment_body"] = model.NewCommentBody(comment.CommentBody)
	if validationErrors["comment_body"] != nil{
		validationErrorFlag = true
		errorMessages["comment_body"] = validationErrors["comment_body"].Error()
	}

	newComment.UserID,validationErrors["user_id"] = model.NewUserID(iuserId)
	if validationErrors["user_id"] != nil{
		validationErrorFlag = true
		errorMessages["user_id"] = validationErrors["user_id"].Error()
	}

	newComment.MovieID,validationErrors["movie_id"] = model.NewMovieID(comment.MovieID)
	if validationErrors["movie_id"] != nil{
		validationErrorFlag = true
		errorMessages["movie_id"] = validationErrors["movie_id"].Error()
	}

	if validationErrorFlag {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Validation Error.",
			"messages": errorMessages,
		})
		c.Abort()
		return
	}

	postCommentErr := postCommentUsecase.PostComment(newComment)
	if postCommentErr != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Comment Post Failed.",
			"messages": postCommentErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posted!",
	})
}

type PostCommentJson struct{
	CommentBody	string	`json:"comment_body"`
	MovieID string	`json:"movie_id"`
}