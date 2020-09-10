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

	if registeredUser != nil{
		return errors.New("Already Registered.")
	}

	user.Token = model.UserToken(uuid.New().String())
	user.Birthday = time.Date(1000, 1, 1, 0, 0, 0, 0, time.Local)
	err2 := u.UserRepository.SetUser(user)
	if err2 != nil{
		return err2
	}
	return nil
}
