package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type IUserRegister interface {
	RegisterUser(newUser *model.User) error
}

type UserRegister struct {
	UserRepository repository.UserRepository
}

func NewUserRegister(u repository.UserRepository) *UserRegister {
	return &UserRegister{
		UserRepository: u,
	}
}

func (u UserRegister) RegisterUser(newUser *model.User) error {
	registeredUserWithToken, findUserErr := u.UserRepository.FindByToken(newUser.Token)
	if findUserErr != nil{
		return findUserErr
	}

	if registeredUserWithToken == nil {
		return errors.New("Invalid Token.")
	}

	registerErr := registeredUserWithToken.Register(newUser.Name, newUser.Password, newUser.Birthday)
	if registerErr != nil {
		return registerErr
	}
	e := u.UserRepository.UpdateUser(registeredUserWithToken)

	if e != nil {
		return e
	}

	return nil
}
