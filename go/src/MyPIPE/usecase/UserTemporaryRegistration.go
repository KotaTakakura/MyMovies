package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
	"time"
)

type IUserTemporaryRegistration interface {
	TemporaryRegister(user *model.User) error
}

type UserTemporaryRegistration struct {
	UserRepository                  repository.UserRepository
	TemporaryRegisterMailRepository repository.TemporaryRegisterMailRepository
}

func NewUserTemporaryRegistration(userRepository repository.UserRepository, temporaryRegisterMailRepository repository.TemporaryRegisterMailRepository) *UserTemporaryRegistration {
	return &UserTemporaryRegistration{
		UserRepository:                  userRepository,
		TemporaryRegisterMailRepository: temporaryRegisterMailRepository,
	}
}

func (u *UserTemporaryRegistration) TemporaryRegister(user *model.User) error {
	registeredUser, _ := u.UserRepository.FindByEmail(user.Email)
	//本登録済み
	if registeredUser != nil && registeredUser.Token == "" {
		return nil
	}

	//仮登録済み・本登録前
	if registeredUser != nil && registeredUser.Token != "" {
		setTokenErr := registeredUser.SetNewToken()
		if setTokenErr != nil {
			return setTokenErr
		}
		updateError := u.UserRepository.UpdateUser(registeredUser)
		if updateError != nil {
			return errors.New("Update Error.")
		}
		return nil
	}
	newUser := model.NewUser(user.Email, time.Date(1000, 1, 1, 0, 0, 0, 0, time.Local))
	setTokenErr := newUser.SetNewToken()
	if setTokenErr != nil {
		return setTokenErr
	}
	err2 := u.UserRepository.SetUser(newUser)
	if err2 != nil {
		return err2
	}

	temporaryRegisterMail := model.NewTemporaryRegisterMail(newUser.Email, newUser.Token)
	emailErr := u.TemporaryRegisterMailRepository.Send(temporaryRegisterMail)
	if emailErr != nil {
		return emailErr
	}
	return nil
}
