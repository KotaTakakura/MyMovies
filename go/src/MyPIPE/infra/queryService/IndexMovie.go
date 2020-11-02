package queryService_infra

import (
	"MyPIPE/domain/queryService"
	"MyPIPE/infra"
)

type IndexMovie struct{}

func NewIndexMovie() *IndexMovie {
	return &IndexMovie{}
}

func (i IndexMovie) Search(page queryService.IndexMovieQueryServicePage, keyWord string, order queryService.IndexMovieQueryServiceOrder) queryService.IndexMovieDTO {
	db := infra.ConnectGorm()
	defer db.Close()
	var movies []queryService.MoviesForIndexMovieDTO
	var count uint64
	db.Table("movies").
		Select("movies.id as movie_id, movies.display_name as movie_display_name,movies.thumbnail_name as thumbnail_name,users.id as user_id,users.name as user_name, users.profile_image_name as user_profile_image_name").
		Joins("inner join users on movies.user_id = users.id").
		Where(`movies.display_name like ? and movies.public = 1`, "%" + keyWord + "%").
		Limit(24).
		Offset((page - 1) * 24).
		Order("movies.created_at " + string(order)).
		Scan(&movies).Count(&count)

	result := queryService.IndexMovieDTO{
		Movies: movies,
		Count:  count,
	}

	return result
}

func (i IndexMovie) All(page queryService.IndexMovieQueryServicePage, order queryService.IndexMovieQueryServiceOrder) queryService.IndexMovieDTO {
	db := infra.ConnectGorm()
	defer db.Close()
	var movies []queryService.MoviesForIndexMovieDTO
	var count uint64

	db.Table("movies").
		Select("movies.id as movie_id, movies.display_name as movie_display_name,movies.thumbnail_name as thumbnail_name,users.id as user_id,users.name as user_name, users.profile_image_name as user_profile_image_name").
		Joins("inner join users on movies.user_id = users.id").
		Where("movies.public = 1").
		Limit(24).
		Offset((page - 1) * 24).
		Order("movies.created_at " + string(order)).
		Scan(&movies).Count(&count)

	//db.Table("movies").
	//	Select("movies.id as movie_id, movies.display_name as movie_display_name,movies.thumbnail_name as thumbnail_name,users.id as user_id,users.name as user_name, users.profile_image_name as user_profile_image_name").
	//	Where("movies.public = 1").
	//	Joins("inner join users on movies.user_id = users.id").
	//	Count(&count)

	result := queryService.IndexMovieDTO{
		Movies: movies,
		Count:  count,
	}

	return result
}
