package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type CommentPersistence struct{
	DatabaseAccessor *gorm.DB
}

func (c *CommentPersistence) GetAll() []model.Comment {
	var comments []model.Comment
	c.DatabaseAccessor.Find(comments)
	return comments
}

func (c *CommentPersistence) FindById(commentID model.CommentID) *model.Comment {
	var comment *model.Comment
	c.DatabaseAccessor.Find(comment, commentID)
	return comment
}

func (c *CommentPersistence) FindByUserId(userId model.UserID) []model.Comment {
	var comments []model.Comment
	c.DatabaseAccessor.Where("user_id = ?", userId).Find(comments)
	return comments
}

func (c *CommentPersistence) FindByMovieId(movieId model.MovieID) []model.Comment {
	var comments []model.Comment
	c.DatabaseAccessor.Where("movie_id = ?", movieId).Find(comments)
	return comments
}
