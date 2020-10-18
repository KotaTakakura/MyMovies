package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type PlayListMoviePersistence struct{}

func NewPlayListMoviePersistence()*PlayListMoviePersistence{
	return &PlayListMoviePersistence{}
}

func (p PlayListMoviePersistence)Find(userId model.UserID,playListId model.PlayListID,movieId model.MovieID)*model.PlayListMovie{
	db := ConnectGorm()
	defer db.Close()
	var playListMovie model.PlayListMovie

	db.Model(&playListMovie).
		Select("play_list_movies.play_list_id as play_list_id, " +
			"play_list_movies.movie_id as movie_id, " +
			"play_list_movies.order as `order`, " +
			"play_list_movies.created_at as created_at, " +
			"play_list_movies.updated_at as updated_at").
		Joins("left join play_lists on play_list_movies.play_list_id = play_lists.id").
		Where("play_lists.user_id = ? and play_list_movies.play_list_id = ? and play_list_movies.movie_id = ?",userId,playListId,movieId).
		First(&playListMovie)

	return &playListMovie
}

func (p PlayListMoviePersistence)FindAll(userId model.UserID,playListId model.PlayListID)[]model.PlayListMovie{
	db := ConnectGorm()
	defer db.Close()
	var playListMovies []model.PlayListMovie
	db.
		Select("play_list_movies.play_list_id as play_list_id, " +
			"play_list_movies.movie_id as movie_id, " +
			"play_list_movies.order as `order`, " +
			"play_list_movies.created_at as created_at, " +
			"play_list_movies.updated_at as updated_at").
		Joins("left join play_lists on play_list_movies.play_list_id = play_lists.id").
		Where("play_lists.user_id = ? and play_list_movies.play_list_id = ?",userId,playListId).
		Find(&playListMovies)
	if len(playListMovies) == 0{
		return nil
	}
	return playListMovies
}

func (p PlayListMoviePersistence)Save(playListMovie *model.PlayListMovie)error{
	db := ConnectGorm()
	defer db.Close()
	var countPlayListMovie int
	db.Table("play_list_movies").Where("play_list_id = ? and movie_id = ?",playListMovie.PlayListID,playListMovie.MovieID).Count(&countPlayListMovie)

	if countPlayListMovie == 0{
		result := db.Create(playListMovie)
		if result.Error != nil{
			return result.Error
		}
		return nil
	}

	return nil
}

func (p PlayListMoviePersistence)SaveAll(playListMovies []model.PlayListMovie)error{
	db := ConnectGorm()
	defer db.Close()

	transactionErr := db.Transaction(func(tx *gorm.DB) error {
		for _,value := range playListMovies{
			result := tx.Table("play_list_movies").Where("play_list_id = ? and movie_id = ?",value.PlayListID,value.MovieID).Update("order",value.Order)
			if result.Error != nil{
				return result.Error
			}
		}
		return nil
	})

	if transactionErr != nil{
		return transactionErr
	}

	return nil
}

func (p PlayListMoviePersistence)Remove(playListMovie *model.PlayListMovie)error{
	db := ConnectGorm()
	defer db.Close()
	result := db.Where("play_list_id = ? and movie_id = ?",playListMovie.PlayListID,playListMovie.MovieID).Delete(&playListMovie)
	if result.Error != nil{
		return result.Error
	}

	return nil
}