package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type UserPersistence struct {
	databaseAccessor *gorm.DB
}

func NewUserPersistence() *UserPersistence {
	return &UserPersistence{
		databaseAccessor: ConnectGorm(),
	}
}

func (u UserPersistence) GetAll() []model.User {
	var allUsers []model.User
	u.databaseAccessor.Find(&allUsers)
	return allUsers
}

func (u UserPersistence) FindByToken(token string) *model.User {
	var user model.User
	u.databaseAccessor.Where("token = ?", token).Take(&user)
	return &user
}

func (u UserPersistence) FindByEmail(email string) *model.User {
	var user model.User
	u.databaseAccessor.Where("email = ?", email).Take(&user)
	return &user
}

func (u UserPersistence) FindById(id int) *model.User {
	var user model.User
	u.databaseAccessor.First(&user, id)
	return &user
}

func (u UserPersistence) SetUser(newUser *model.User) {
	u.databaseAccessor.Create(&newUser)
}
