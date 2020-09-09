package repository

import (
	"MyPIPE/domain/model"
)

type UserRepository interface {
	GetAll() []model.User
	FindById(id uint64) *model.User
	FindByToken(token string) *model.User
	FindByEmail(email string) *model.User
	SetUser(*model.User)
}
