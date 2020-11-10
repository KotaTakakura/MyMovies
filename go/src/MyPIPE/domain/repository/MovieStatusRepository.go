package repository

import "MyPIPE/domain/model"

type MovieStatusRepository interface {
	Find(movieId model.MovieID)(*model.MovieStatusModel,error)
	Save(movieStatusModel *model.MovieStatusModel)error
}
