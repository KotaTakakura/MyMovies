package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type ISetEmailChangeToken interface {
	SetEmailChangeToken(setEmailChangeTokenDTO *SetEmailChangeTokenDTO) error
}

type SetEmailChangeToken struct {
	UserRepository            repository.UserRepository
	ChangeEmailMailRepository repository.ChangeEmailMail
}

func NewSetEmailChangeToken(userRepository repository.UserRepository, changeEmailMailRepository repository.ChangeEmailMail) *SetEmailChangeToken {
	return &SetEmailChangeToken{
		UserRepository:            userRepository,
		ChangeEmailMailRepository: changeEmailMailRepository,
	}
}

func (s SetEmailChangeToken) SetEmailChangeToken(setEmailChangeTokenDTO *SetEmailChangeTokenDTO) error {

	alreadyRegisteredUser, findAlreadyRegisteredUserErr := s.UserRepository.FindByEmail(setEmailChangeTokenDTO.Email)
	if findAlreadyRegisteredUserErr != nil {
		return findAlreadyRegisteredUserErr
	}
	if alreadyRegisteredUser != nil && alreadyRegisteredUser.ID != setEmailChangeTokenDTO.UserID {
		return nil
	}

	user, userErr := s.UserRepository.FindById(setEmailChangeTokenDTO.UserID)
	if userErr != nil {
		return userErr
	}
	if user == nil {
		return errors.New("No Such User.")
	}

	token, tokenErr := user.SetChangeEmailToken(setEmailChangeTokenDTO.Email)
	if tokenErr != nil {
		return tokenErr
	}

	updateUserErr := s.UserRepository.UpdateUser(user)
	if updateUserErr != nil {
		return updateUserErr
	}

	sendMailErr := s.ChangeEmailMailRepository.Send(setEmailChangeTokenDTO.Email, token)
	if sendMailErr != nil {
		return sendMailErr
	}

	return nil
}

type SetEmailChangeTokenDTO struct {
	UserID model.UserID
	Email  model.UserEmail
}

func NewSetEmailChangeTokenDTO(userId model.UserID, email model.UserEmail) *SetEmailChangeTokenDTO {
	return &SetEmailChangeTokenDTO{
		UserID: userId,
		Email:  email,
	}
}
