package infra

import (
	"MyPIPE/domain/model"
	"errors"
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
	if e.RowsAffected == 0 {
		return nil, nil
	}
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
	if e.RowsAffected == 0 {
		return nil, nil
	}
	if e.Error != nil {
		return nil, e.Error
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

func (u UserPersistence) FindByPasswordRememberToken(token model.UserPasswordRememberToken) (*model.User, error) {
	db := ConnectGorm()
	defer db.Close()
	var user model.User
	e := db.Where("password_remember_token = ?", token).Take(&user)
	if e.RowsAffected == 0 {
		return nil, errors.New("No Such User.")
	}
	if e.Error != nil {
		return nil, e.Error
	}

	return &user, nil
}

func (u UserPersistence) FindByEmailChangeToken(token model.UserEmailChangeToken) (*model.User, error) {
	db := ConnectGorm()
	defer db.Close()
	var user model.User
	result := db.Where("email_change_token = ?", token).Take(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
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

func (u UserPersistence) Remove(userId model.UserID) error {
	db := ConnectGorm()
	defer db.Close()

	transactionErr := db.Transaction(func(tx *gorm.DB) error {
		deleteUserResult := tx.Where("id = ?", userId).Delete(model.User{})
		if deleteUserResult.Error != nil {
			return deleteUserResult.Error
		}

		deleteMovieResult := tx.Where("user_id = ?", userId).Delete(model.Movie{})
		if deleteMovieResult.Error != nil {
			return deleteMovieResult.Error
		}

		deleteMovieEvaluationResult := tx.Where("user_id = ?", userId).Delete(model.MovieEvaluation{})
		if deleteMovieEvaluationResult.Error != nil {
			return deleteMovieEvaluationResult.Error
		}

		deleteCommentResult := tx.Where("user_id = ?", userId).Delete(model.Comment{})
		if deleteCommentResult.Error != nil {
			return deleteCommentResult.Error
		}

		deletePlayListMoviesResult := tx.Raw("Delete plm From play_list_movies As plm Left Join play_lists As pl on plm.play_list_id = pl.id Where pl.user_id = ?", userId)
		if deletePlayListMoviesResult.Error != nil {
			return deletePlayListMoviesResult.Error
		}

		deletePlayListsResult := tx.Where("user_id = ?", userId).Delete(model.PlayList{})
		if deletePlayListsResult.Error != nil {
			return deletePlayListsResult.Error
		}
		return nil
	})

	if transactionErr != nil {
		return transactionErr
	}
	return nil
}
