package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"mime/multipart"
	"os"
	"reflect"
	"testing"
)

func TestChangeThumbnail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := usecase.NewChangeThumbnail(movieRepository, thumbnailUploadRepository)

	cases := []struct {
		FilePath string
	}{
		{
			FilePath: "./../TestThumbnail.jpg",
		},
	}

	for _, Case := range cases {
		file, fileErr := os.Open(Case.FilePath)
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}

		changeThumbnailDTO := &usecase.ChangeThumbnailDTO{
			UserID:  model.UserID(10),
			MovieID: model.MovieID(10),
			Thumbnail: &model.MovieThumbnail{
				Name: "NewThumbnailName",
				File: file,
			},
		}

		movieRepository.EXPECT().FindByUserIdAndMovieId(changeThumbnailDTO.UserID, changeThumbnailDTO.MovieID).Return(&model.Movie{
			ID:            changeThumbnailDTO.MovieID,
			StoreName:     "StoreName",
			DisplayName:   "DisplayName",
			Description:   "Description",
			ThumbnailName: "OldThumbnailName",
			UserID:        changeThumbnailDTO.UserID,
			Public:        0,
			Status:        0,
		}, nil)

		movieRepository.EXPECT().Update(gomock.Any()).DoAndReturn(func(data interface{}) (*model.Movie, error) {
			if reflect.TypeOf(data) != reflect.TypeOf(model.Movie{}) {
				t.Fatal("Type Movie Not Match.")
			}
			if data.(model.Movie).ID != changeThumbnailDTO.MovieID {
				t.Fatal("MovieID Not Match.")
			}
			if data.(model.Movie).UserID != changeThumbnailDTO.UserID {
				t.Fatal("UserID Not Match.")
			}
			return nil, nil
		})

		thumbnailUploadRepository.EXPECT().Upload(gomock.Any(), gomock.Any()).DoAndReturn(func(data1 interface{}, data2 interface{}) error {
			if reflect.DeepEqual(data1.(multipart.File), &changeThumbnailDTO.Thumbnail.File) {
				t.Fatal("File Not Match.")
			}
			return nil
		})

		result := changeThumbnailUsecase.ChangeThumbnail(changeThumbnailDTO)

		if result != nil {
			t.Fatal("Usecase Error")
		}

	}
}

func TestChangeThumbnail_MovieRepository_FindByUserIdAndMovieId_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := usecase.NewChangeThumbnail(movieRepository, thumbnailUploadRepository)

	cases := []struct {
		FilePath string
	}{
		{
			FilePath: "./../TestThumbnail.jpg",
		},
	}

	for _, Case := range cases {
		file, fileErr := os.Open(Case.FilePath)
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}

		changeThumbnailDTO := &usecase.ChangeThumbnailDTO{
			UserID:  model.UserID(10),
			MovieID: model.MovieID(10),
			Thumbnail: &model.MovieThumbnail{
				Name: "NewThumbnailName",
				File: file,
			},
		}

		movieRepository.EXPECT().FindByUserIdAndMovieId(changeThumbnailDTO.UserID, changeThumbnailDTO.MovieID).Return(nil, errors.New("ERROR"))

		result := changeThumbnailUsecase.ChangeThumbnail(changeThumbnailDTO)

		if result == nil {
			t.Fatal("Usecase Error")
		}

	}
}

func TestChangeThumbnail_MovieRepository_Update_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := usecase.NewChangeThumbnail(movieRepository, thumbnailUploadRepository)

	cases := []struct {
		FilePath string
	}{
		{
			FilePath: "./../TestThumbnail.jpg",
		},
	}

	for _, Case := range cases {
		file, fileErr := os.Open(Case.FilePath)
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}

		changeThumbnailDTO := &usecase.ChangeThumbnailDTO{
			UserID:  model.UserID(10),
			MovieID: model.MovieID(10),
			Thumbnail: &model.MovieThumbnail{
				Name: "NewThumbnailName",
				File: file,
			},
		}

		movieRepository.EXPECT().FindByUserIdAndMovieId(changeThumbnailDTO.UserID, changeThumbnailDTO.MovieID).Return(&model.Movie{
			ID:            changeThumbnailDTO.MovieID,
			StoreName:     "StoreName",
			DisplayName:   "DisplayName",
			Description:   "Description",
			ThumbnailName: "OldThumbnailName",
			UserID:        changeThumbnailDTO.UserID,
			Public:        0,
			Status:        0,
		}, nil)

		movieRepository.EXPECT().Update(gomock.Any()).Return(nil, errors.New("ERROR"))

		result := changeThumbnailUsecase.ChangeThumbnail(changeThumbnailDTO)

		if result == nil {
			t.Fatal("Usecase Error")
		}

	}
}

func TestChangeThumbnail_ThumbnailUploadRepository_Upload_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	thumbnailUploadRepository := mock_repository.NewMockThumbnailUploadRepository(ctrl)
	changeThumbnailUsecase := usecase.NewChangeThumbnail(movieRepository, thumbnailUploadRepository)

	cases := []struct {
		FilePath string
	}{
		{
			FilePath: "./../TestThumbnail.jpg",
		},
	}

	for _, Case := range cases {
		file, fileErr := os.Open(Case.FilePath)
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}

		changeThumbnailDTO := &usecase.ChangeThumbnailDTO{
			UserID:  model.UserID(10),
			MovieID: model.MovieID(10),
			Thumbnail: &model.MovieThumbnail{
				Name: "NewThumbnailName",
				File: file,
			},
		}

		movieRepository.EXPECT().FindByUserIdAndMovieId(changeThumbnailDTO.UserID, changeThumbnailDTO.MovieID).Return(&model.Movie{
			ID:            changeThumbnailDTO.MovieID,
			StoreName:     "StoreName",
			DisplayName:   "DisplayName",
			Description:   "Description",
			ThumbnailName: "OldThumbnailName",
			UserID:        changeThumbnailDTO.UserID,
			Public:        0,
			Status:        0,
		}, nil)

		movieRepository.EXPECT().Update(gomock.Any()).DoAndReturn(func(data interface{}) (*model.Movie, error) {
			if reflect.TypeOf(data) != reflect.TypeOf(model.Movie{}) {
				t.Fatal("Type Movie Not Match.")
			}
			if data.(model.Movie).ID != changeThumbnailDTO.MovieID {
				t.Fatal("MovieID Not Match.")
			}
			if data.(model.Movie).UserID != changeThumbnailDTO.UserID {
				t.Fatal("UserID Not Match.")
			}
			return nil, nil
		})

		thumbnailUploadRepository.EXPECT().Upload(gomock.Any(), gomock.Any()).Return(errors.New("ERROR"))

		result := changeThumbnailUsecase.ChangeThumbnail(changeThumbnailDTO)

		if result == nil {
			t.Fatal("Usecase Error")
		}

	}
}
