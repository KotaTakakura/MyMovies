package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type MoviePersistence struct{
	databaseAccessor	*gorm.DB
}

func NewMoviePersistence() *MoviePersistence{
	return &MoviePersistence{
		databaseAccessor: ConnectGorm(),
	}
}

func (m *MoviePersistence)GetAll() []model.Movie{
	var movies []model.Movie
	m.databaseAccessor.Find(movies)
	return movies
}

func (m *MoviePersistence)FindById(id uint64) *model.Movie{
	var movies *model.Movie
	m.databaseAccessor.Find(movies, id)
	return movies
}

func (m *MoviePersistence)FindByUserId(userId uint64) *model.Movie{
	var movies *model.Movie
	m.databaseAccessor.Where("user_id = ?", userId).Find(movies)
	return movies
}