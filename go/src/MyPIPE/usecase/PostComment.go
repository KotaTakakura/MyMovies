package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type IPostComment interface {
	PostComment(postCommentDTO *PostCommentDTO) error
}

type PostComment struct {
	CommentRepository repository.CommentRepository
	MovieRepository   repository.MovieRepository
}

func NewPostComment(c repository.CommentRepository, m repository.MovieRepository) *PostComment {
	return &PostComment{
		CommentRepository: c,
		MovieRepository:   m,
	}
}

func (p PostComment) PostComment(postCommentDTO *PostCommentDTO) error {

	movie, movieFindErr := p.MovieRepository.FindById(postCommentDTO.MovieID)
	if movieFindErr != nil {
		return movieFindErr
	}
	if movie == nil {
		return errors.New("No Such Movie.")
	}

	newComment := model.NewComment(postCommentDTO.UserID, postCommentDTO.MovieID, postCommentDTO.Body)

	err := p.CommentRepository.Save(newComment)
	if err != nil {
		return err
	}
	return nil
}

type PostCommentDTO struct {
	UserID  model.UserID
	MovieID model.MovieID
	Body    model.CommentBody
}

func NewPostCommentDTO(userId model.UserID, movieId model.MovieID, body model.CommentBody) *PostCommentDTO {
	return &PostCommentDTO{
		UserID:  userId,
		MovieID: movieId,
		Body:    body,
	}
}
