package infra

import (
	"MyPIPE/domain/model"
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

func (p PlayListMoviePersistence)Remove(playListMovie *model.PlayListMovie)error{
	db := ConnectGorm()
	defer db.Close()
	result := db.Where("play_list_id = ? and movie_id = ?",playListMovie.PlayListID,playListMovie.MovieID).Delete(&playListMovie)
	if result.Error != nil{
		return result.Error
	}

	return nil
}