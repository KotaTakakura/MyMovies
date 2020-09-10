package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type CommentPersistence struct {
	DatabaseAccessor *gorm.DB
}

func NewCommentPersistence() *CommentPersistence {
	return &CommentPersistence{
		DatabaseAccessor: ConnectGorm(),
	}
}

func (c *CommentPersistence) GetAll() ([]model.Comment, error) {
	var comments []model.Comment
	result := c.DatabaseAccessor.Find(comments)
	if result != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (c *CommentPersistence) FindById(commentID model.CommentID) (*model.Comment, error) {
	var comment *model.Comment
	result := c.DatabaseAccessor.Find(comment, commentID)
	if result != nil {
		return nil, result.Error
	}
	return comment, nil
}

func (c *CommentPersistence) FindByUserId(userId model.UserID) ([]model.Comment, error) {
	var comments []model.Comment
	result := c.DatabaseAccessor.Where("user_id = ?", userId).Find(comments)
	if result != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (c *CommentPersistence) FindByMovieId(movieId model.MovieID) ([]model.Comment, error) {
	var comments []model.Comment
	result := c.DatabaseAccessor.Where("movie_id = ?", movieId).Find(comments)
	if result != nil {
		return nil, result.Error
	}
	return comments, nil
}
