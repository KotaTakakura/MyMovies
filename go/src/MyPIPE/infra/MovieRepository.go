package infra

import (
	"MyPIPE/domain/model"
	"fmt"
	"github.com/jinzhu/gorm"
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
	var count int
	result := db.First(&movies, uint64(id)).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	if count == 0 {
		return nil, nil
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

func (m *MoviePersistence) Remove(userId model.UserID, movieId model.MovieID) error {
	db := ConnectGorm()
	defer db.Close()

	transactionErr := db.Transaction(func(tx *gorm.DB) error {

		deleteMovieEvaluationResult := tx.Where("movie_id = ?", movieId).Delete(model.MovieEvaluation{})
		if deleteMovieEvaluationResult.Error != nil {
			return deleteMovieEvaluationResult.Error
		}

		deleteCommentResult := tx.Where("movie_id = ?", movieId).Delete(model.Comment{})
		if deleteCommentResult.Error != nil {
			return deleteCommentResult.Error
		}

		deletePlayListMoviesResult := tx.Where("movie_id = ?", movieId).Delete(model.PlayListMovie{})
		if deletePlayListMoviesResult.Error != nil {
			return deletePlayListMoviesResult.Error
		}

		deleteMovieResult := tx.Where("id = ? and user_id = ?", movieId, userId).Delete(model.Movie{})
		if deleteMovieResult.Error != nil {
			return deleteMovieResult.Error
		}

		return nil
	})

	if transactionErr != nil {
		return transactionErr
	}

	return nil
}
