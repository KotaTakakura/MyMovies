package handler

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"MyPIPE/usecase"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostComment struct {
	CommentRepository  repository.CommentRepository
	MovieRepository    repository.MovieRepository
	PostCommentUsecase usecase.IPostComment
}

func NewPostComment(commentRepo repository.CommentRepository, movieRepo repository.MovieRepository, postCommentUsecase usecase.IPostComment) *PostComment {
	return &PostComment{
		CommentRepository:  commentRepo,
		MovieRepository:    movieRepo,
		PostCommentUsecase: postCommentUsecase,
	}
}

func (postComment PostComment) PostComment(c *gin.Context) {
	userId := jwt.ExtractClaims(c)["id"]
	iuserId := uint64(userId.(float64))

	var comment PostCommentJson
	bindErr := c.Bind(&comment)
	if bindErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Server Error.",
			"messages": bindErr.Error(),
		})
		c.Abort()
		return
	}

	var newComment model.Comment
	validationErrors := map[string]error{}
	errorMessages := map[string]string{}
	validationErrorFlag := false

	newComment.Body, validationErrors["comment_body"] = model.NewCommentBody(comment.CommentBody)
	if validationErrors["comment_body"] != nil {
		validationErrorFlag = true
		errorMessages["comment_body"] = validationErrors["comment_body"].Error()
	}

	newComment.UserID, validationErrors["user_id"] = model.NewUserID(iuserId)
	if validationErrors["user_id"] != nil {
		validationErrorFlag = true
		errorMessages["user_id"] = validationErrors["user_id"].Error()
	}

	newComment.MovieID, validationErrors["movie_id"] = model.NewMovieID(comment.MovieID)
	if validationErrors["movie_id"] != nil {
		validationErrorFlag = true
		errorMessages["movie_id"] = validationErrors["movie_id"].Error()
	}

	if validationErrorFlag {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Validation Error.",
			"messages": errorMessages,
		})
		c.Abort()
		return
	}

	postCommentErr := postComment.PostCommentUsecase.PostComment(newComment)
	if postCommentErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":   "Comment Post Failed.",
			"messages": postCommentErr.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posted!",
	})
}

type PostCommentJson struct {
	CommentBody string `json:"comment_body"`
	MovieID     uint64 `json:"movie_id"`
}
