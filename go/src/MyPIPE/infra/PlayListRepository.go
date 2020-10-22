package infra

import (
	"MyPIPE/domain/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type PlayListPersistence struct{}

func NewPlayListPersistence() *PlayListPersistence {
	return &PlayListPersistence{}
}

func (p PlayListPersistence) FindByID(playListID model.PlayListID) (*model.PlayList, error) {
	db := ConnectGorm()
	defer db.Close()
	var playLists model.PlayList
	result := db.Where("id = ?", playListID).Take(&playLists)
	if result.Error != nil {
		return nil, result.Error
	}
	result = db.Table("play_list_movies").Where("play_list_id = ?", playListID).Pluck("movie_id", &playLists.PlayListMovies)
	if result.Error != nil {
		return nil, result.Error
	}
	return &playLists, nil
}

func (p PlayListPersistence) FindByName(playListName model.PlayListName) ([]model.PlayList, error) {
	panic("implement me")
}

func (p PlayListPersistence) FindByUserID(playListUserID model.UserID) ([]model.PlayList, error) {
	panic("implement me")
}

func (p PlayListPersistence) FindByUserIDAndName(playListUserID model.UserID, playListName model.PlayListName) ([]model.PlayList, error) {
	db := ConnectGorm()
	defer db.Close()
	var playLists []model.PlayList
	resultFindPlayList := db.Where("user_id = ? and name = ?", playListUserID, playListName).Find(&playLists)
	if resultFindPlayList.Error != nil {
		return nil, resultFindPlayList.Error
	}
	return playLists, nil
}

func (p *PlayListPersistence) Save(playList *model.PlayList) error {
	db := ConnectGorm()
	defer db.Close()
	if playList.ID == 0 {
		createResult := db.Create(&playList)
		if createResult.Error != nil {
			return createResult.Error
		}
		return nil
	}

	transactionErr := db.Transaction(func(tx *gorm.DB) error {

		saveResult := tx.Save(playList)
		if saveResult.Error != nil {
			return saveResult.Error
		}
		return nil
	})

	if transactionErr != nil {
		return transactionErr
	}

	return nil
}

type playListMovie struct {
	PlayListID model.PlayListID
	MovieID    model.MovieID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (p *PlayListPersistence) Remove(userId model.UserID, playListId model.PlayListID) error {
	db := ConnectGorm()
	defer db.Close()
	transactionErr := db.Transaction(func(tx *gorm.DB) error {

		deletePlayListMoviesResult := tx.Exec("Delete From play_list_movies Where play_list_id in (Select id From play_lists Where user_id = ? and id = ?)", userId, playListId)
		if deletePlayListMoviesResult.Error != nil {
			return deletePlayListMoviesResult.Error
		}
		fmt.Println(userId)
		deletePlayListResult := tx.Exec("Delete From play_lists Where id = ? and user_id = ?", playListId, userId)
		if deletePlayListResult.Error != nil {
			return deletePlayListResult.Error
		}
		fmt.Println(playListId)
		return nil
	})

	if transactionErr != nil {
		return transactionErr
	}

	return nil
}
