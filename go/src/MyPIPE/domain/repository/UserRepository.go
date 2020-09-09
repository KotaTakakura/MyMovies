package repository

import (
	"MyPIPE/domain/model"
)

type UserRepository interface {
	GetAll() []model.User
	FindById(id model.UserID) *model.User
	FindByToken(token model.UserToken) *model.User
	FindByEmail(email model.UserEmail) *model.User
	SetUser(*model.User)
}
