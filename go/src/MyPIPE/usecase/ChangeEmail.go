package usecase

import (
	"MyPIPE/domain/model"
	"MyPIPE/domain/repository"
	"errors"
)

type IChangeEmail interface {
	ChangeEmail(changeEmailDTO *ChangeEmailDTO) error
}

type ChangeEmail struct {
	UserRepository repository.UserRepository
}

func NewChangeEmail(userRepository repository.UserRepository) *ChangeEmail {
	return &ChangeEmail{
		UserRepository: userRepository,
	}
}

func (c ChangeEmail) ChangeEmail(changeEmailDTO *ChangeEmailDTO) error {
	user, findUserErr := c.UserRepository.FindByEmailChangeToken(changeEmailDTO.Token)
	if findUserErr != nil {
		return findUserErr
	}
	if user == nil {
		return errors.New("No Such User.")
	}

	changeEmailErr := user.ChangeEmail()
	if changeEmailErr != nil {
		return changeEmailErr
	}

	updateUserErr := c.UserRepository.UpdateUser(user)
	if updateUserErr != nil {
		return updateUserErr
	}

	return nil
}

type ChangeEmailDTO struct {
	Token model.UserEmailChangeToken
}

func NewChangeEmailDTO(token model.UserEmailChangeToken) *ChangeEmailDTO {
	return &ChangeEmailDTO{
		Token: token,
	}
}
