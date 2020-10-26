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
	iuserId := uint64(jwt.ExtractClaims(c)["id"].(float64))

	var comment PostCommentJson
	c.Bind(&comment)

	errorMessages := map[string]string{}
	validationErrorFlag := false

	body, bodyErr := model.NewCommentBody(comment.CommentBody)
	if bodyErr != nil {
		validationErrorFlag = true
		errorMessages["comment_body"] = bodyErr.Error()
	}

	userId, userIdErr := model.NewUserID(iuserId)
	if userIdErr != nil {
		validationErrorFlag = true
		errorMessages["user_id"] = userIdErr.Error()
	}

	movieId, movieIdErr := model.NewMovieID(comment.MovieID)
	if movieIdErr != nil {
		validationErrorFlag = true
		errorMessages["movie_id"] = movieIdErr.Error()
	}

	if validationErrorFlag {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":   "Validation Error.",
			"messages": errorMessages,
		})
		c.Abort()
		return
	}

	postCommentDTO := usecase.NewPostCommentDTO(userId, movieId, body)

	postCommentErr := postComment.PostCommentUsecase.PostComment(postCommentDTO)
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
