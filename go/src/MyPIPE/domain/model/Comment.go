package model

import "time"

type CommentID uint64

func NewCommentID(commentId uint64) CommentID {
	return CommentID(commentId)
}

type CommentBody string

func NewCommentBody(commentBody string) CommentBody {
	return CommentBody(commentBody)
}

type Comment struct {
	ID        CommentID
	Body      CommentBody
	MovieID   MovieID
	Movie     Movie
	UserID    UserID
	User      User
	Replies   []Comment `gorm:"many2many:replies;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
