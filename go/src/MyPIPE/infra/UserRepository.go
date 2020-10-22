package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type UserPersistence struct{}

func NewUserPersistence() *UserPersistence {
	return &UserPersistence{}
}

func (u UserPersistence) GetAll() ([]model.User, error) {
	db := ConnectGorm()
	defer db.Close()
	var allUsers []model.User
	e := db.Find(&allUsers)
	if e != nil {
		return nil, e.Error
	}
	return allUsers, nil
}

func (u UserPersistence) FindByToken(token model.UserToken) (*model.User, error) {
	db := ConnectGorm()
	defer db.Close()
	var user model.User
	e := db.Where("token = ?", token).Take(&user)
	if e.Error != nil {
		return nil, e.Error
	}

	return &user, nil
}

func (u UserPersistence) FindByEmail(email model.UserEmail) (*model.User, error) {
	db := ConnectGorm()
	defer db.Close()
	var user model.User
	e := db.Where("email = ?", email).Take(&user)
	if e.Error != nil {
		return nil, e.Error
	}

	return &user, nil
}

func (u UserPersistence) FindById(id model.UserID) (*model.User, error) {
	db := ConnectGorm()
	defer db.Close()
	var user model.User
	e := db.First(&user, uint64(id))
	if e.Error != nil {
		return nil, e.Error
	}
	if e.RowsAffected == 0 {
		return nil, nil
	}

	return &user, nil
}

func (u UserPersistence) FindByName(name model.UserName) (*model.User, error) {
	db := ConnectGorm()
	defer db.Close()
	var user model.User
	e := db.Where("name = ?", name).Take(&user)
	if e.Error != nil {
		return nil, e.Error
	}

	return &user, nil
}

func (u UserPersistence) SetUser(newUser *model.User) error {
	db := ConnectGorm()
	defer db.Close()
	e := db.Create(&newUser)
	if e.Error != nil {
		return e.Error
	}
	return nil
}

func (u UserPersistence) UpdateUser(updateUser *model.User) error {
	db := ConnectGorm()
	defer db.Close()

	transactionErr := db.Transaction(func(tx *gorm.DB) error {
		e := tx.Save(updateUser)
		if e.Error != nil {
			return e.Error
		}
		return nil
	})

	if transactionErr != nil {
		return transactionErr
	}

	return nil
}
