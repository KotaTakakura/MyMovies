package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type ISetPasswordRememberToken interface {
	SetPasswordRememberToken(setPasswordRememberTokenDTO *SetPasswordRememberTokenDTO) error
}

type SetPasswordRememberToken struct {
	UserRepository               repository.UserRepository
	ResetPasswordEmailRepository repository.ResetPasswordEmail
}

func NewSetPasswordRememberToken(userRepository repository.UserRepository, resetPasswordEmailRepository repository.ResetPasswordEmail) *SetPasswordRememberToken {
	return &SetPasswordRememberToken{
		UserRepository:               userRepository,
		ResetPasswordEmailRepository: resetPasswordEmailRepository,
	}
}

func (r SetPasswordRememberToken) SetPasswordRememberToken(setPasswordRememberTokenDTO *SetPasswordRememberTokenDTO) error {
	user, findUserErr := r.UserRepository.FindByEmail(setPasswordRememberTokenDTO.Email)
	if findUserErr != nil {
		return findUserErr
	}
	if user == nil {
		return errors.New("No Such User.")
	}

	rememberPasswordToken, rememberPasswordTokenErr := user.SetPasswordRememberToken()
	if rememberPasswordTokenErr != nil {
		return rememberPasswordTokenErr
	}

	setUserErr := r.UserRepository.UpdateUser(user)
	if setUserErr != nil {
		return setUserErr
	}

	sendMailErr := r.ResetPasswordEmailRepository.Send(user.Email, rememberPasswordToken)
	if sendMailErr != nil {
		return sendMailErr
	}

	return nil
}

type SetPasswordRememberTokenDTO struct {
	Email model.UserEmail
}

func NewSetPasswordRememberTokenDTO(email model.UserEmail) *SetPasswordRememberTokenDTO {
	return &SetPasswordRememberTokenDTO{
		Email: email,
	}
}
