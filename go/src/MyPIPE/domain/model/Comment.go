package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type CommentID uint64

func NewCommentID(commentId uint64) (CommentID, error) {
	err := validation.Validate(commentId,
		validation.Required,
	)
	if err != nil {
		return CommentID(0), err
	}
	return CommentID(commentId), nil
}

type CommentBody string

func NewCommentBody(commentBody string) (CommentBody, error) {
	err := validation.Validate(commentBody,
		validation.Required,
		validation.RuneLength(1, 1000),
	)
	if err != nil {
		return CommentBody(""), err
	}
	return CommentBody(commentBody), nil
}

type Comment struct {
	ID        CommentID `gorm:"primaryKey"`
	Body      CommentBody
	MovieID   MovieID
	Movie     Movie
	UserID    UserID
	User      User
	Reply     CommentID `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewComment(userId UserID, movieId MovieID, body CommentBody)*Comment{
	return &Comment{
		UserID: userId,
		MovieID: movieId,
		Body: body,
	}
}