package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type UserRegister struct{
	UserRepository	repository.UserRepository
}

func NewUserRegister(u repository.UserRepository) *UserRegister{
	return &UserRegister{
		UserRepository: u,
	}
}

func (u UserRegister)RegisterUser(newUser *model.User) error{
	registeredUserWithToken, _ := u.UserRepository.FindByToken(newUser.Token)

	if registeredUserWithToken == nil{
		return errors.New("Not Temporary Registered.")
	}

	registeredUserWithToken.EmptyToken()
	registeredUserWithToken.Password = newUser.Password
	registeredUserWithToken.Name = newUser.Name
	registeredUserWithToken.Birthday = newUser.Birthday
	e := u.UserRepository.UpdateUser(registeredUserWithToken)

	if e != nil {
		return e
	}

	return nil
}