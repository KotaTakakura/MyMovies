package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	domain_service "MyPIPE/domain/service/User"
	"MyPIPE/infra"
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
		return errors.New("Invalid Token.")
	}

	if !registeredUserWithToken.TemporaryRegisteredWithinOneHour() {
		return errors.New("Invalid Tokene.")
	}

	checkNameService := domain_service.NewCheckNameExists(infra.NewUserPersistence())
	if checkNameService.CheckNameExists(newUser.Name) {
		return errors.New("User Name Already Exists.")
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