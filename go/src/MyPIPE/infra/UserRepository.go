package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type UserPersistence struct {
	DatabaseAccessor *gorm.DB
}

func NewUserPersistence() *UserPersistence {
	return &UserPersistence{
		DatabaseAccessor: ConnectGorm(),
	}
}

func (u UserPersistence) GetAll() []model.User {
	var allUsers []model.User
	u.DatabaseAccessor.Find(&allUsers)
	return allUsers
}

func (u UserPersistence) FindByToken(token model.UserToken) *model.User {
	var user model.User
	u.DatabaseAccessor.Where("token = ?", token).Take(&user)
	return &user
}

func (u UserPersistence) FindByEmail(email model.UserEmail) *model.User {
	var user model.User
	u.DatabaseAccessor.Where("email = ?", email).Take(&user)
	return &user
}

func (u UserPersistence) FindById(id model.UserID) *model.User {
	var user model.User
	u.DatabaseAccessor.First(&user, id)
	return &user
}

func (u UserPersistence) SetUser(newUser *model.User) *gorm.DB {
	err := u.DatabaseAccessor.Create(&newUser)
	return err
}
