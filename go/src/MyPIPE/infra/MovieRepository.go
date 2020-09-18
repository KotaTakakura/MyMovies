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
	var movies model.Movie
	result := m.databaseAccessor.First(&movies, uint64(id))
	if result != nil {
		return nil, result.Error
	}
	return &movies, nil
}

func (m *MoviePersistence) FindByUserId(userId model.MovieID) (*model.Movie, error) {
	var movies *model.Movie
	result := m.databaseAccessor.Where("user_id = ?", userId).Find(movies)
	if result != nil {
		return nil, result.Error
	}
	return movies, nil
}

func (m *MoviePersistence)Save(movie model.Movie)(*model.Movie,error){
	result := m.databaseAccessor.Create(&movie)
	if result.Error != nil{
		return nil,result.Error
	}
	return &movie,nil
}