package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type IResetPassword interface {
	ResetPassword(resetPasswordDTO *ResetPasswordDTO)error
}

type ResetPassword struct{
	UserRepository repository.UserRepository
}

func NewResetPassword(userRepository repository.UserRepository)*ResetPassword{
	return &ResetPassword{
		UserRepository: userRepository,
	}
}

func (r ResetPassword)ResetPassword(resetPasswordDTO *ResetPasswordDTO)error{
	user,findUserErr := r.UserRepository.FindByPasswordRememberToken(resetPasswordDTO.PasswordRememberToken)
	if findUserErr != nil{
		return findUserErr
	}
	if user == nil{
		return errors.New("No Such User.")
	}

	resetPasswordErr := user.ResetPassword(resetPasswordDTO.Password)
	if resetPasswordErr != nil{
		return resetPasswordErr
	}

	updateUserErr := r.UserRepository.UpdateUser(user)
	if updateUserErr != nil{
		return updateUserErr
	}
	
	return nil
}

type ResetPasswordDTO struct{
	PasswordRememberToken model.UserPasswordRememberToken
	Password	model.UserPassword
}

func NewResetPasswordDTO(passwordRememberToken model.UserPasswordRememberToken,password model.UserPassword)*ResetPasswordDTO{
	return &ResetPasswordDTO{
		PasswordRememberToken: passwordRememberToken,
		Password:	password,
	}
}