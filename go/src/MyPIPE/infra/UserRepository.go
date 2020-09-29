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

	e = u.DatabaseAccessor.Table("good_movies").Where("user_id = ?",uint64(user.ID)).Pluck("movie_id",&user.GoodMovies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("bad_movies").Where("user_id = ?",uint64(user.ID)).Pluck("movie_id",&user.BadMovies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("movies").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.Movies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("comments").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.Comments)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("play_lists").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.PlayLists)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("follow_users").Where("user_id = ?",uint64(user.ID)).Pluck("follow_id",&user.Follows)
	if e.Error != nil {
		return nil, e.Error
	}
	return &user, nil
}

func (u UserPersistence) FindByEmail(email model.UserEmail) (*model.User, error) {
	var user model.User
	e := u.DatabaseAccessor.Where("email = ?", email).Take(&user)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("good_movies").Where("user_id = ?",uint64(user.ID)).Pluck("movie_id",&user.GoodMovies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("bad_movies").Where("user_id = ?",uint64(user.ID)).Pluck("movie_id",&user.BadMovies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("movies").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.Movies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("comments").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.Comments)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("play_lists").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.PlayLists)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("follow_users").Where("user_id = ?",uint64(user.ID)).Pluck("follow_id",&user.Follows)
	if e.Error != nil {
		return nil, e.Error
	}

	return &user, nil
}

func (u UserPersistence) FindById(id model.UserID) (*model.User, error) {
	var user model.User
	e := u.DatabaseAccessor.First(&user, uint64(id))
	if e.Error != nil {
		return nil, e.Error
	}
	if e.RowsAffected == 0{
		return nil,nil
	}

	e = u.DatabaseAccessor.Table("good_movies").Where("user_id = ?",uint64(id)).Pluck("movie_id",&user.GoodMovies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("bad_movies").Where("user_id = ?",uint64(id)).Pluck("movie_id",&user.BadMovies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("movies").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.Movies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("comments").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.Comments)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("play_lists").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.PlayLists)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("follow_users").Where("user_id = ?",uint64(user.ID)).Pluck("follow_id",&user.Follows)
	if e.Error != nil {
		return nil, e.Error
	}
	return &user, nil
}

func (u UserPersistence) FindByName(name model.UserName) (*model.User, error) {
	var user model.User
	e := u.DatabaseAccessor.Where("name = ?", name).Take(&user)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("good_movies").Where("user_id = ?",uint64(user.ID)).Pluck("movie_id",&user.GoodMovies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("bad_movies").Where("user_id = ?",uint64(user.ID)).Pluck("movie_id",&user.BadMovies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("movies").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.Movies)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("comments").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.Comments)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("play_lists").Where("user_id = ?",uint64(user.ID)).Pluck("id",&user.PlayLists)
	if e.Error != nil {
		return nil, e.Error
	}

	e = u.DatabaseAccessor.Table("follow_users").Where("user_id = ?",uint64(user.ID)).Pluck("follow_id",&user.Follows)
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

	transactionErr := u.DatabaseAccessor.Transaction(func(tx *gorm.DB) error {
		tx.Exec("Delete From good_movies Where user_id = ?",updateUser.ID)

		for _,goodMovieId := range updateUser.GoodMovies{
			tx.Exec("Insert Into good_movies (user_id,movie_id) Values (?,?)",updateUser.ID,goodMovieId)
		}

		tx.Exec("Delete From bad_movies Where user_id = ?",updateUser.ID)

		for _,badMovieId := range updateUser.BadMovies{
			tx.Exec("Insert Into bad_movies (user_id,movie_id) Values (?,?)",updateUser.ID,badMovieId)
		}

		e := tx.Save(updateUser)
		if e.Error != nil {
			return e.Error
		}
		return nil
	})

	if transactionErr != nil{
		return transactionErr
	}

	return nil
}