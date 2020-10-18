package infra

import (
	"MyPIPE/domain/model"
)

type CommentPersistence struct {}

func NewCommentPersistence() *CommentPersistence {
	return &CommentPersistence{}
}

func (c *CommentPersistence) GetAll() ([]model.Comment, error) {
	db := ConnectGorm()
	defer db.Close()
	var comments []model.Comment
	result := db.Find(comments)
	if result != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (c *CommentPersistence) FindById(commentID model.CommentID) (*model.Comment, error) {
	db := ConnectGorm()
	defer db.Close()
	var comment *model.Comment
	result := db.Find(comment, commentID)
	if result != nil {
		return nil, result.Error
	}
	return comment, nil
}

func (c *CommentPersistence) FindByUserId(userId model.UserID) ([]model.Comment, error) {
	db := ConnectGorm()
	defer db.Close()
	var comments []model.Comment
	result := db.Where("user_id = ?", userId).Find(comments)
	if result != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (c *CommentPersistence) FindByMovieId(movieId model.MovieID) ([]model.Comment, error) {
	db := ConnectGorm()
	defer db.Close()
	var comments []model.Comment
	result := db.Where("movie_id = ?", movieId).Find(comments)
	if result != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (c *CommentPersistence)Save(comment *model.Comment)error{
	db := ConnectGorm()
	defer db.Close()
	if comment.ID == 0{
		createResult := db.Create(&comment)
		if createResult != nil{
			return createResult.Error
		}
		return nil
	}
	updateResult := db.Update(&comment)
	if updateResult.Error != nil{
		return updateResult.Error
	}
	return nil
}
