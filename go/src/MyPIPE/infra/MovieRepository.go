package infra

import (
	"MyPIPE/domain/model"
	"fmt"
)

type MoviePersistence struct{}

func NewMoviePersistence() *MoviePersistence {
	return &MoviePersistence{}
}

func (m *MoviePersistence) GetAll() ([]model.Movie, error) {
	db := ConnectGorm()
	defer db.Close()
	var movies []model.Movie
	result := db.Find(movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}

func (m *MoviePersistence) FindById(id model.MovieID) (*model.Movie, error) {
	db := ConnectGorm()
	defer db.Close()
	var movies model.Movie
	result := db.First(&movies, uint64(id))
	if result.Error != nil {
		return nil, result.Error
	}
	return &movies, nil
}

func (m *MoviePersistence) FindByUserId(userId model.MovieID) (*model.Movie, error) {
	db := ConnectGorm()
	defer db.Close()
	var movies *model.Movie
	result := db.Where("user_id = ?", userId).Find(movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}

func (m *MoviePersistence) FindByUserIdAndMovieId(userId model.UserID, movieId model.MovieID) (*model.Movie, error) {
	db := ConnectGorm()
	defer db.Close()
	var movies model.Movie
	fmt.Println(userId)
	fmt.Println(movieId)
	result := db.Table("movies").Where("user_id = ? AND id = ?", userId, movieId).First(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(movies)
	return &movies, nil
}

func (m *MoviePersistence) Save(movie model.Movie) (*model.Movie, error) {
	db := ConnectGorm()
	defer db.Close()
	result := db.Create(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}

func (m *MoviePersistence) Update(movie model.Movie) (*model.Movie, error) {
	db := ConnectGorm()
	defer db.Close()
	result := db.Save(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}
