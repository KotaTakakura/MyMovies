package infra

import (
	"MyPIPE/domain/model"
	"github.com/jinzhu/gorm"
)

type MoviePersistence struct {
	databaseAccessor *gorm.DB
}

func NewMoviePersistence() *MoviePersistence {
	return &MoviePersistence{
		databaseAccessor: ConnectGorm(),
	}
}

func (m *MoviePersistence) GetAll() ([]model.Movie, error) {
	var movies []model.Movie
	result := m.databaseAccessor.Find(movies)
	if result != nil {
		return nil, result.Error
	}
	return movies, nil
}

func (m *MoviePersistence) FindById(id model.MovieID) (*model.Movie, error) {
	var movies *model.Movie
	result := m.databaseAccessor.Find(movies, id)
	if result != nil {
		return nil, result.Error
	}
	return movies, nil
}

func (m *MoviePersistence) FindByUserId(userId model.MovieID) (*model.Movie, error) {
	var movies *model.Movie
	result := m.databaseAccessor.Where("user_id = ?", userId).Find(movies)
	if result != nil {
		return nil, result.Error
	}
	return movies, nil
}

func (m *MoviePersistence)Save(movie model.Movie)error{
	result := m.databaseAccessor.Save(&movie)
	if result.Error != nil{
		return result.Error
	}
	return nil
}