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

func (u UserPersistence) GetAll() ([]model.User, error) {
	var allUsers []model.User
	e := u.DatabaseAccessor.Find(&allUsers)
	if e != nil {
		return nil, e.Error
	}
	return allUsers, nil
}

func (u UserPersistence) FindByToken(token model.UserToken) (*model.User, error) {
	var user model.User
	e := u.DatabaseAccessor.Where("token = ?", token).Take(&user)
	if e.Error != nil {
		return nil, e.Error
	}
	return &user, nil
}

func (u UserPersistence) FindByEmail(email model.UserEmail) (*model.User, error) {
	var user model.User
	e := u.DatabaseAccessor.Where("email = ?", email).Take(&user)
	if e != nil {
		return nil, e.Error
	}
	return &user, nil
}

func (u UserPersistence) FindById(id model.UserID) (*model.User, error) {
	var user model.User
	e := u.DatabaseAccessor.First(&user, id)
	if e != nil {
		return nil, e.Error
	}
	return &user, nil
}

func (u UserPersistence) SetUser(newUser *model.User) error {
	e := u.DatabaseAccessor.Create(&newUser)
	if e != nil {
		return e.Error
	}
	return nil
}

func (u UserPersistence) UpdateUser(updateUser *model.User) error{
	e := u.DatabaseAccessor.Save(updateUser)
	if e != nil {
		return e.Error
	}
	return nil
}