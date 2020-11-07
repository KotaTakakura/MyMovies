package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type IDeleteComment interface {
	DeleteComment(deleteCommentDTO *DeleteCommentDTO) error
}

type DeleteComment struct {
	CommentRepository repository.CommentRepository
}

func NewDeleteComment(commentRepository repository.CommentRepository) *DeleteComment {
	return &DeleteComment{
		CommentRepository: commentRepository,
	}
}

func (d DeleteComment) DeleteComment(deleteCommentDTO *DeleteCommentDTO) error {
	comment, findCommentErr := d.CommentRepository.FindByIdAndUserID(deleteCommentDTO.CommentID, deleteCommentDTO.UserID)
	if findCommentErr != nil {
		return findCommentErr
	}
	if comment == nil {
		return errors.New("No Such Comment.")
	}

	removeCommentErr := d.CommentRepository.Remove(comment)
	if removeCommentErr != nil {
		return removeCommentErr
	}

	return nil
}

type DeleteCommentDTO struct {
	CommentID model.CommentID
	UserID    model.UserID
}

func NewDeleteCommentDTO(commentId model.CommentID, userId model.UserID) *DeleteCommentDTO {
	return &DeleteCommentDTO{
		CommentID: commentId,
		UserID:    userId,
	}
}
