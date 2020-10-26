package test

import (
	mock_factory "MyPIPE/Test/mock/factory"
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestAddPlayListItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepsitory := mock_repository.NewMockPlayListRepository(ctrl)
	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	playListMovieFactory := mock_factory.NewMockIPlayListMovie(ctrl)
	addPlayListItemUsecase := usecase.NewAddPlayListItem(playListRepsitory, playListMovieRepository, playListMovieFactory)

	trueCases := []usecase.AddPlayListItemAddJson{
		usecase.AddPlayListItemAddJson{
			PlayListID: model.PlayListID(10),
			UserID:     model.UserID(20),
			MovieID:    model.MovieID(30),
		},
	}

	for _, trueCase := range trueCases {
		playListRepsitory.EXPECT().FindByID(trueCase.PlayListID).Return(&model.PlayList{
			ID:             trueCase.PlayListID,
			UserID:         trueCase.UserID,
			Name:           "TestName",
			Description:    "TestDescription",
			PlayListMovies: nil,
		}, nil)

		playListMovieFactory.EXPECT().CreatePlayListMovie(trueCase.PlayListID, trueCase.MovieID).Return(&model.PlayListMovie{
			PlayListID: trueCase.PlayListID,
			MovieID:    trueCase.MovieID,
			Order:      0,
		}, nil)

		playListMovieRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.PlayListMovie{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.PlayListMovie).PlayListID != trueCase.PlayListID {
				t.Fatal("PlayListID Not Match.")
			}
			if data.(*model.PlayListMovie).MovieID != trueCase.MovieID {
				t.Fatal("MovieID Not Match.")
			}
			return nil
		})

		result := addPlayListItemUsecase.AddPlayListItem(&trueCase)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestAddPlayListItem_PlayListRepository_FindByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepsitory := mock_repository.NewMockPlayListRepository(ctrl)
	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	playListMovieFactory := mock_factory.NewMockIPlayListMovie(ctrl)
	addPlayListItemUsecase := usecase.NewAddPlayListItem(playListRepsitory, playListMovieRepository, playListMovieFactory)

	cases := []usecase.AddPlayListItemAddJson{
		usecase.AddPlayListItemAddJson{
			PlayListID: model.PlayListID(10),
			UserID:     model.UserID(20),
			MovieID:    model.MovieID(30),
		},
	}

	for _, Case := range cases {
		playListRepsitory.EXPECT().FindByID(Case.PlayListID).Return(nil, nil)

		result := addPlayListItemUsecase.AddPlayListItem(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}

	for _, Case := range cases {
		playListRepsitory.EXPECT().FindByID(Case.PlayListID).Return(&model.PlayList{
			ID:             Case.PlayListID,
			UserID:         Case.UserID + 1,
			Name:           "TestName",
			Description:    "TestDescription",
			PlayListMovies: nil,
		}, nil)

		result := addPlayListItemUsecase.AddPlayListItem(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}

	for _, Case := range cases {
		playListRepsitory.EXPECT().FindByID(Case.PlayListID).Return(&model.PlayList{
			ID:             Case.PlayListID,
			UserID:         Case.UserID,
			Name:           "TestName",
			Description:    "TestDescription",
			PlayListMovies: nil,
		}, errors.New("ERROR!!"))

		result := addPlayListItemUsecase.AddPlayListItem(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestAddPlayListItem_PlaylistMovieFactory_CreatePlayListMovie_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []usecase.AddPlayListItemAddJson{
		usecase.AddPlayListItemAddJson{
			PlayListID: model.PlayListID(10),
			UserID:     model.UserID(20),
			MovieID:    model.MovieID(30),
		},
	}

	for _, Case := range cases {

		playListRepsitory := mock_repository.NewMockPlayListRepository(ctrl)
		playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
		playListMovieFactory := mock_factory.NewMockIPlayListMovie(ctrl)
		addPlayListItemUsecase := usecase.NewAddPlayListItem(playListRepsitory, playListMovieRepository, playListMovieFactory)

		playListRepsitory.EXPECT().FindByID(Case.PlayListID).Return(&model.PlayList{
			ID:             Case.PlayListID,
			UserID:         Case.UserID,
			Name:           "TestName",
			Description:    "TestDescription",
			PlayListMovies: nil,
		}, nil)

		playListMovieFactory.EXPECT().CreatePlayListMovie(Case.PlayListID, Case.MovieID).Return(&model.PlayListMovie{
			PlayListID: Case.PlayListID,
			MovieID:    Case.MovieID,
			Order:      0,
		}, errors.New("ERROR!!"))

		result := addPlayListItemUsecase.AddPlayListItem(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestAddPlayListItem_PlayListMovieRepository_Save_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepsitory := mock_repository.NewMockPlayListRepository(ctrl)
	playListMovieRepository := mock_repository.NewMockPlayListMovieRepository(ctrl)
	playListMovieFactory := mock_factory.NewMockIPlayListMovie(ctrl)
	addPlayListItemUsecase := usecase.NewAddPlayListItem(playListRepsitory, playListMovieRepository, playListMovieFactory)

	trueCases := []usecase.AddPlayListItemAddJson{
		usecase.AddPlayListItemAddJson{
			PlayListID: model.PlayListID(10),
			UserID:     model.UserID(20),
			MovieID:    model.MovieID(30),
		},
	}

	for _, trueCase := range trueCases {
		playListRepsitory.EXPECT().FindByID(trueCase.PlayListID).Return(&model.PlayList{
			ID:             trueCase.PlayListID,
			UserID:         trueCase.UserID,
			Name:           "TestName",
			Description:    "TestDescription",
			PlayListMovies: nil,
		}, nil)

		playListMovieFactory.EXPECT().CreatePlayListMovie(trueCase.PlayListID, trueCase.MovieID).Return(&model.PlayListMovie{
			PlayListID: trueCase.PlayListID,
			MovieID:    trueCase.MovieID,
			Order:      0,
		}, nil)

		playListMovieRepository.EXPECT().Save(gomock.Any()).Return(errors.New("ERROR"))

		result := addPlayListItemUsecase.AddPlayListItem(&trueCase)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}
