package test

import (
	mock_factory "MyPIPE/Test/mock/factory"
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"fmt"
	"github.com/golang/mock/gomock"
	"os"
	"reflect"
	"testing"
	"errors"
)

func TestPostMovie(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileUploadRepository := mock_repository.NewMockFileUpload(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	movieFactory := mock_factory.NewMockIMovieModelFactory(ctrl)
	postMovieUsecase := usecase.NewPostMovie(fileUploadRepository,thumbnailUploadRepository,movieRepository,movieFactory)

	cases := []struct{
		ThumbnailFilePath string
		MovieFilePath string
	}{
		{
			ThumbnailFilePath: "./../TestThumbnail.jpg",
			MovieFilePath: "./../TestMovie.mp4",
		},
	}

	for _,Case := range cases{
		thumbnailFile, thumbnailFileErr := os.Open(Case.ThumbnailFilePath)
		if thumbnailFileErr != nil {
			fmt.Println(thumbnailFileErr)
			return
		}

		movieFile, movieFileErr := os.Open(Case.MovieFilePath)
		if movieFileErr != nil {
			fmt.Println(movieFileErr)
			return
		}

		postMovieDTO := &usecase.PostMovieDTO{
			MovieFile: &model.MovieFile{
				StoreName:  "MovieStoreName",
				File:       movieFile,
			},
			Thumbnail: &model.MovieThumbnail{
				Name:       "ThumbnailName",
				File:       thumbnailFile,
			},
			UserID:    model.UserID(10),
		}

		movieFactory.EXPECT().CreateMovieModel(postMovieDTO.UserID,postMovieDTO.MovieFile,postMovieDTO.Thumbnail).Return(&model.Movie{
			ID:            model.MovieID(0),
			StoreName:     postMovieDTO.MovieFile.StoreName,
			DisplayName:   "DisplayName",
			Description:   "Description",
			ThumbnailName: postMovieDTO.Thumbnail.Name,
			UserID:        postMovieDTO.UserID,
			Public:        0,
			Status:        0,
		},nil)

		movieRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(data interface{})(*model.Movie,error){
			if reflect.TypeOf(data) != reflect.TypeOf(model.Movie{}) {
				t.Fatal("Type Movie Not Match.")
			}
			if data.(model.Movie).ID != model.MovieID(0) {
				t.Fatal("MovieID Not Match.")
			}
			if data.(model.Movie).StoreName != postMovieDTO.MovieFile.StoreName {
				t.Fatal("StoreName Not Match.")
			}
			if data.(model.Movie).DisplayName != "DisplayName" {
				t.Fatal("DisplayName Not Match.")
			}
			if data.(model.Movie).Description != "Description" {
				t.Fatal("Description Not Match.")
			}
			if data.(model.Movie).ThumbnailName != postMovieDTO.Thumbnail.Name {
				t.Fatal("ThumbnailName Not Match.")
			}
			if data.(model.Movie).Public != 0 {
				t.Fatal("Publid Not Match.")
			}
			if data.(model.Movie).Status != 0 {
				t.Fatal("Status Not Match.")
			}
			return &model.Movie{
				ID:            model.MovieID(10),
				StoreName:     postMovieDTO.MovieFile.StoreName,
				DisplayName:   "DisplayName",
				Description:   "Description",
				ThumbnailName: postMovieDTO.Thumbnail.Name,
				UserID:        postMovieDTO.UserID,
				Public:        0,
				Status:        0,
			},nil
		})

		fileUploadRepository.EXPECT().Upload(postMovieDTO.MovieFile.File,postMovieDTO.MovieFile.FileHeader,model.MovieID(10))

		thumbnailUploadRepository.EXPECT().Upload(postMovieDTO.Thumbnail.File,gomock.Any()).DoAndReturn(func(data1 interface{},data2 interface{})error{
			if reflect.TypeOf(data2) != reflect.TypeOf(model.Movie{}) {
				t.Fatal("Type Movie Not Match.")
			}
			if data2.(model.Movie).ID != model.MovieID(10) {
				t.Fatal("MovieID Not Match.")
			}
			if data2.(model.Movie).StoreName != postMovieDTO.MovieFile.StoreName {
				t.Fatal("StoreName Not Match.")
			}
			if data2.(model.Movie).DisplayName != "DisplayName" {
				t.Fatal("DisplayName Not Match.")
			}
			if data2.(model.Movie).Description != "Description" {
				t.Fatal("Description Not Match.")
			}
			if data2.(model.Movie).ThumbnailName != postMovieDTO.Thumbnail.Name {
				t.Fatal("ThumbnailName Not Match.")
			}
			if data2.(model.Movie).Public != 0 {
				t.Fatal("Publid Not Match.")
			}
			if data2.(model.Movie).Status != 0 {
				t.Fatal("Status Not Match.")
			}
			return nil
		})

		_,err := postMovieUsecase.PostMovie(postMovieDTO)
		if err != nil{
			t.Fatal("useccase Error")
		}

	}
}

func TestPostMovie_CreateMovieFactory_Error(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileUploadRepository := mock_repository.NewMockFileUpload(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	movieFactory := mock_factory.NewMockIMovieModelFactory(ctrl)
	postMovieUsecase := usecase.NewPostMovie(fileUploadRepository,thumbnailUploadRepository,movieRepository,movieFactory)

	cases := []struct{
		ThumbnailFilePath string
		MovieFilePath string
	}{
		{
			ThumbnailFilePath: "./../TestThumbnail.jpg",
			MovieFilePath: "./../TestMovie.mp4",
		},
	}

	for _,Case := range cases{
		thumbnailFile, thumbnailFileErr := os.Open(Case.ThumbnailFilePath)
		if thumbnailFileErr != nil {
			fmt.Println(thumbnailFileErr)
			return
		}

		movieFile, movieFileErr := os.Open(Case.MovieFilePath)
		if movieFileErr != nil {
			fmt.Println(movieFileErr)
			return
		}

		postMovieDTO := &usecase.PostMovieDTO{
			MovieFile: &model.MovieFile{
				StoreName:  "MovieStoreName",
				File:       movieFile,
			},
			Thumbnail: &model.MovieThumbnail{
				Name:       "ThumbnailName",
				File:       thumbnailFile,
			},
			UserID:    model.UserID(10),
		}

		movieFactory.EXPECT().CreateMovieModel(postMovieDTO.UserID,postMovieDTO.MovieFile,postMovieDTO.Thumbnail).Return(nil,errors.New("ERROR"))

		_,err := postMovieUsecase.PostMovie(postMovieDTO)
		if err == nil{
			t.Fatal("useccase Error")
		}

	}
}

func TestPostMovie_MovieRepository_Save_Error(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileUploadRepository := mock_repository.NewMockFileUpload(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	movieFactory := mock_factory.NewMockIMovieModelFactory(ctrl)
	postMovieUsecase := usecase.NewPostMovie(fileUploadRepository,thumbnailUploadRepository,movieRepository,movieFactory)

	cases := []struct{
		ThumbnailFilePath string
		MovieFilePath string
	}{
		{
			ThumbnailFilePath: "./../TestThumbnail.jpg",
			MovieFilePath: "./../TestMovie.mp4",
		},
	}

	for _,Case := range cases{
		thumbnailFile, thumbnailFileErr := os.Open(Case.ThumbnailFilePath)
		if thumbnailFileErr != nil {
			fmt.Println(thumbnailFileErr)
			return
		}

		movieFile, movieFileErr := os.Open(Case.MovieFilePath)
		if movieFileErr != nil {
			fmt.Println(movieFileErr)
			return
		}

		postMovieDTO := &usecase.PostMovieDTO{
			MovieFile: &model.MovieFile{
				StoreName:  "MovieStoreName",
				File:       movieFile,
			},
			Thumbnail: &model.MovieThumbnail{
				Name:       "ThumbnailName",
				File:       thumbnailFile,
			},
			UserID:    model.UserID(10),
		}

		movieFactory.EXPECT().CreateMovieModel(postMovieDTO.UserID,postMovieDTO.MovieFile,postMovieDTO.Thumbnail).Return(&model.Movie{
			ID:            model.MovieID(0),
			StoreName:     postMovieDTO.MovieFile.StoreName,
			DisplayName:   "DisplayName",
			Description:   "Description",
			ThumbnailName: postMovieDTO.Thumbnail.Name,
			UserID:        postMovieDTO.UserID,
			Public:        0,
			Status:        0,
		},nil)

		movieRepository.EXPECT().Save(gomock.Any()).Return(nil,errors.New("ERROR"))

		_,err := postMovieUsecase.PostMovie(postMovieDTO)
		if err == nil{
			t.Fatal("useccase Error")
		}

	}
}

func TestPostMovie_FileUploadRepository_Upload_Error(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileUploadRepository := mock_repository.NewMockFileUpload(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	movieFactory := mock_factory.NewMockIMovieModelFactory(ctrl)
	postMovieUsecase := usecase.NewPostMovie(fileUploadRepository,thumbnailUploadRepository,movieRepository,movieFactory)

	cases := []struct{
		ThumbnailFilePath string
		MovieFilePath string
	}{
		{
			ThumbnailFilePath: "./../TestThumbnail.jpg",
			MovieFilePath: "./../TestMovie.mp4",
		},
	}

	for _,Case := range cases{
		thumbnailFile, thumbnailFileErr := os.Open(Case.ThumbnailFilePath)
		if thumbnailFileErr != nil {
			fmt.Println(thumbnailFileErr)
			return
		}

		movieFile, movieFileErr := os.Open(Case.MovieFilePath)
		if movieFileErr != nil {
			fmt.Println(movieFileErr)
			return
		}

		postMovieDTO := &usecase.PostMovieDTO{
			MovieFile: &model.MovieFile{
				StoreName:  "MovieStoreName",
				File:       movieFile,
			},
			Thumbnail: &model.MovieThumbnail{
				Name:       "ThumbnailName",
				File:       thumbnailFile,
			},
			UserID:    model.UserID(10),
		}

		movieFactory.EXPECT().CreateMovieModel(postMovieDTO.UserID,postMovieDTO.MovieFile,postMovieDTO.Thumbnail).Return(&model.Movie{
			ID:            model.MovieID(0),
			StoreName:     postMovieDTO.MovieFile.StoreName,
			DisplayName:   "DisplayName",
			Description:   "Description",
			ThumbnailName: postMovieDTO.Thumbnail.Name,
			UserID:        postMovieDTO.UserID,
			Public:        0,
			Status:        0,
		},nil)

		movieRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(data interface{})(*model.Movie,error){
			if reflect.TypeOf(data) != reflect.TypeOf(model.Movie{}) {
				t.Fatal("Type Movie Not Match.")
			}
			if data.(model.Movie).ID != model.MovieID(0) {
				t.Fatal("MovieID Not Match.")
			}
			if data.(model.Movie).StoreName != postMovieDTO.MovieFile.StoreName {
				t.Fatal("StoreName Not Match.")
			}
			if data.(model.Movie).DisplayName != "DisplayName" {
				t.Fatal("DisplayName Not Match.")
			}
			if data.(model.Movie).Description != "Description" {
				t.Fatal("Description Not Match.")
			}
			if data.(model.Movie).ThumbnailName != postMovieDTO.Thumbnail.Name {
				t.Fatal("ThumbnailName Not Match.")
			}
			if data.(model.Movie).Public != 0 {
				t.Fatal("Publid Not Match.")
			}
			if data.(model.Movie).Status != 0 {
				t.Fatal("Status Not Match.")
			}
			return &model.Movie{
				ID:            model.MovieID(10),
				StoreName:     postMovieDTO.MovieFile.StoreName,
				DisplayName:   "DisplayName",
				Description:   "Description",
				ThumbnailName: postMovieDTO.Thumbnail.Name,
				UserID:        postMovieDTO.UserID,
				Public:        0,
				Status:        0,
			},nil
		})

		fileUploadRepository.EXPECT().Upload(postMovieDTO.MovieFile.File,postMovieDTO.MovieFile.FileHeader,model.MovieID(10)).Return(errors.New("ERROR"))

		_,err := postMovieUsecase.PostMovie(postMovieDTO)
		if err == nil{
			t.Fatal("useccase Error")
		}

	}
}

func TestPostMovie_ThumbnailUploadRepository_Upload_Error(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileUploadRepository := mock_repository.NewMockFileUpload(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	movieFactory := mock_factory.NewMockIMovieModelFactory(ctrl)
	postMovieUsecase := usecase.NewPostMovie(fileUploadRepository,thumbnailUploadRepository,movieRepository,movieFactory)

	cases := []struct{
		ThumbnailFilePath string
		MovieFilePath string
	}{
		{
			ThumbnailFilePath: "./../TestThumbnail.jpg",
			MovieFilePath: "./../TestMovie.mp4",
		},
	}

	for _,Case := range cases{
		thumbnailFile, thumbnailFileErr := os.Open(Case.ThumbnailFilePath)
		if thumbnailFileErr != nil {
			fmt.Println(thumbnailFileErr)
			return
		}

		movieFile, movieFileErr := os.Open(Case.MovieFilePath)
		if movieFileErr != nil {
			fmt.Println(movieFileErr)
			return
		}

		postMovieDTO := &usecase.PostMovieDTO{
			MovieFile: &model.MovieFile{
				StoreName:  "MovieStoreName",
				File:       movieFile,
			},
			Thumbnail: &model.MovieThumbnail{
				Name:       "ThumbnailName",
				File:       thumbnailFile,
			},
			UserID:    model.UserID(10),
		}

		movieFactory.EXPECT().CreateMovieModel(postMovieDTO.UserID,postMovieDTO.MovieFile,postMovieDTO.Thumbnail).Return(&model.Movie{
			ID:            model.MovieID(0),
			StoreName:     postMovieDTO.MovieFile.StoreName,
			DisplayName:   "DisplayName",
			Description:   "Description",
			ThumbnailName: postMovieDTO.Thumbnail.Name,
			UserID:        postMovieDTO.UserID,
			Public:        0,
			Status:        0,
		},nil)

		movieRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(data interface{})(*model.Movie,error){
			if reflect.TypeOf(data) != reflect.TypeOf(model.Movie{}) {
				t.Fatal("Type Movie Not Match.")
			}
			if data.(model.Movie).ID != model.MovieID(0) {
				t.Fatal("MovieID Not Match.")
			}
			if data.(model.Movie).StoreName != postMovieDTO.MovieFile.StoreName {
				t.Fatal("StoreName Not Match.")
			}
			if data.(model.Movie).DisplayName != "DisplayName" {
				t.Fatal("DisplayName Not Match.")
			}
			if data.(model.Movie).Description != "Description" {
				t.Fatal("Description Not Match.")
			}
			if data.(model.Movie).ThumbnailName != postMovieDTO.Thumbnail.Name {
				t.Fatal("ThumbnailName Not Match.")
			}
			if data.(model.Movie).Public != 0 {
				t.Fatal("Publid Not Match.")
			}
			if data.(model.Movie).Status != 0 {
				t.Fatal("Status Not Match.")
			}
			return &model.Movie{
				ID:            model.MovieID(10),
				StoreName:     postMovieDTO.MovieFile.StoreName,
				DisplayName:   "DisplayName",
				Description:   "Description",
				ThumbnailName: postMovieDTO.Thumbnail.Name,
				UserID:        postMovieDTO.UserID,
				Public:        0,
				Status:        0,
			},nil
		})

		fileUploadRepository.EXPECT().Upload(postMovieDTO.MovieFile.File,postMovieDTO.MovieFile.FileHeader,model.MovieID(10))

		thumbnailUploadRepository.EXPECT().Upload(postMovieDTO.Thumbnail.File,gomock.Any()).Return(errors.New("Error"))

		_,err := postMovieUsecase.PostMovie(postMovieDTO)
		if err == nil{
			t.Fatal("useccase Error")
		}

	}
}