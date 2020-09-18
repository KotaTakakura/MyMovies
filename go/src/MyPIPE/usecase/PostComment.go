package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type PostComment struct {
	CommentRepository	repository.CommentRepository
	MovieRepository	repository.MovieRepository
}

func NewPostComment(c repository.CommentRepository,m repository.MovieRepository) *PostComment{
	return &PostComment{
		CommentRepository: c,
		MovieRepository: m,
	}
}

func (p PostComment) PostComment(comment model.Comment)error{

	_,movieFindErr := p.MovieRepository.FindById(comment.MovieID)
	if movieFindErr != nil{
		return movieFindErr
	}

	err := p.CommentRepository.Save(&comment)
	if err != nil {
		return err
	}
	return nil
}