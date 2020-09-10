package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
	"github.com/google/uuid"
	"time"
)

type UserTemporaryRegistration struct {
	UserRepository repository.UserRepository
}

func NewUserTemporaryRegistration(userRepository repository.UserRepository) *UserTemporaryRegistration {
	return &UserTemporaryRegistration{
		UserRepository: userRepository,
	}
}

func (u *UserTemporaryRegistration) TemporaryRegister(user *model.User) error {
	registeredUser, err := u.UserRepository.FindByEmail(user.Email)

	if err != nil {
		return err
	}

	if registeredUser != nil {
		return errors.New("Already Registered.")
	}

	user.Token = model.UserToken(uuid.New().String())
	user.Birthday = time.Date(1000, 1, 1, 0, 0, 0, 0, time.Local)
	u.UserRepository.SetUser(user)
	return nil
}
