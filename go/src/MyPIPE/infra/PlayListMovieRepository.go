package infra

import (
	"MyPIPE/domain/model"
)

type PlayListMoviePersistence struct{}

func NewPlayListMoviePersistence()*PlayListMoviePersistence{
	return &PlayListMoviePersistence{}
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
//SaveMultiple(playListMovies []model.PlayListMovie)error{}