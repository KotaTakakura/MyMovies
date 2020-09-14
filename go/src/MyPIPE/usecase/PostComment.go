package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
)

type PostComment struct {
	UserRepository repository.UserRepository
}

func NewPostComment(u repository.UserRepository) *PostComment{
	return &PostComment{
		UserRepository: u,
	}
}

func (p PostComment) PostComment(comment model.Comment)error{
	poster,userFindErr := p.UserRepository.FindById(comment.UserID)
	if userFindErr != nil {
		return userFindErr
	}
	errPostComment := poster.PostComment(comment)
	if errPostComment != nil{
		return errPostComment
	}
	err := p.UserRepository.UpdateUser(poster)
	if err!= nil{
		return err
	}
	return nil
}