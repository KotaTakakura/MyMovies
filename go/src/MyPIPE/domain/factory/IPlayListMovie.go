package factory

import "MyPIPE/domain/model"

type IPlayListMovie interface{
	CreatePlayListMovie(playListId model.PlayListID,movieId model.MovieID)(*model.PlayListMovie,error)
}
