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
	e := u.DatabaseAccessor.
		Preload("Movies").
		Preload("Comments").
		Preload("GoodMovies").
		Preload("BadMovies").
		Preload("PlayLists").
		Preload("Follows").Find(&allUsers)
	if e != nil {
		return nil, e.Error
	}
	return allUsers, nil
}

func (u UserPersistence) FindByToken(token model.UserToken) (*model.User, error) {
	var user model.User
	e := u.DatabaseAccessor.
		Where("token = ?", token).
		Preload("Movies").
		Preload("Comments").
		Preload("GoodMovies").
		Preload("BadMovies").
		Preload("PlayLists").
		Preload("Follows").
		Take(&user)
	if e.Error != nil {
		return nil, e.Error
	}
	return &user, nil
}

func (u UserPersistence) FindByEmail(email model.UserEmail) (*model.User, error) {
	var user model.User
	e := u.DatabaseAccessor.Where("email = ?", email).
		//Preload("Movies").
		//Preload("Comments").
		//Preload("GoodMovies").
		//Preload("BadMovies").
		//Preload("PlayLists").
		//Preload("Follows").
		Take(&user)
	if e.Error != nil {
		return nil, e.Error
	}
	return &user, nil
}

func (u UserPersistence) FindById(id model.UserID) (*model.User, error) {
	var user model.User
	e := u.DatabaseAccessor.
		Preload("Movies").
		Preload("Comments").
		Preload("GoodMovies").
		Preload("BadMovies").
		Preload("PlayLists").
		Preload("Follows").
		First(&user, uint64(id))
	if e.Error != nil {
		return nil, e.Error
	}
	return &user, nil
}

func (u UserPersistence) FindByName(name model.UserName) (*model.User, error) {
	var user model.User
	e := u.DatabaseAccessor.Where("name = ?", name).
		Preload("Movies").
		Preload("Comments").
		Preload("GoodMovies").
		Preload("BadMovies").
		Preload("PlayLists").
		Preload("Follows").
		Take(&user)
	if e.Error != nil {
		return nil, e.Error
	}
	return &user, nil
}

func (u UserPersistence) SetUser(newUser *model.User) error {
	e := u.DatabaseAccessor.Create(&newUser)
	if e.Error != nil {
		return e.Error
	}
	return nil
}

func (u UserPersistence) UpdateUser(updateUser *model.User) error{

	for _, commentToAppend := range updateUser.CommentsToAppend {
		u.DatabaseAccessor.Create(&commentToAppend)
	}

	for _, commentToDelete := range updateUser.CommentsToDelete {
		u.DatabaseAccessor.Delete(&commentToDelete)
	}

	e := u.DatabaseAccessor.Save(updateUser)
	if e.Error != nil {
		return e.Error
	}
	return nil
}