package repository

import (
	"MyPIPE/domain/model"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	FindById(id model.UserID) (*model.User, error)
	FindByToken(token model.UserToken) (*model.User, error)
	FindByEmail(email model.UserEmail) (*model.User, error)
	FindByName(name model.UserName) (*model.User, error)
	FindByPasswordRememberToken(token model.UserPasswordRememberToken) (*model.User, error)
	FindByEmailChangeToken(token model.UserEmailChangeToken) (*model.User, error)
	SetUser(*model.User) error
	UpdateUser(*model.User) error
}
