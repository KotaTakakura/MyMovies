package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
	"github.com/google/uuid"
	"time"
)

type IUserTemporaryRegistration interface{
	TemporaryRegister(user *model.User) error
}

type UserTemporaryRegistration struct {
	UserRepository repository.UserRepository
}

func NewUserTemporaryRegistration(userRepository repository.UserRepository) *UserTemporaryRegistration {
	return &UserTemporaryRegistration{
		UserRepository: userRepository,
	}
}

func (u *UserTemporaryRegistration) TemporaryRegister(user *model.User) error {
	registeredUser, _ := u.UserRepository.FindByEmail(user.Email)

	//本登録済み
	if registeredUser != nil && registeredUser.Token == ""{
		return errors.New("Already Registered.")
	}

	//仮登録済み・本登録前
	if registeredUser != nil && registeredUser.Token != ""{
		registeredUser.Token = model.UserToken(uuid.New().String())
		updateError := u.UserRepository.UpdateUser(registeredUser)
		if updateError != nil{
			return errors.New("Update Error.")
		}
		return nil
	}
	newUser := model.NewUser()
	newUser.Email = user.Email
	newUser.Token = model.UserToken(uuid.New().String())
	newUser.Birthday = time.Date(1000, 1, 1, 0, 0, 0, 0, time.Local)
	err2 := u.UserRepository.SetUser(newUser)
	if err2 != nil{
		return err2
	}
	return nil
}
