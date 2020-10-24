package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestCreatePlayList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
	createPlayListUsecase := usecase.NewCreatePlayList(userRepository, playListRepository)

	trueCases := []usecase.CreatePlayListDTO{
		usecase.CreatePlayListDTO{
			UserID:              model.UserID(10),
			PlayListName:        model.PlayListName("TestPlayListName"),
			PlayListDescription: model.PlayListDescription("TestPlayListDescription"),
		},
	}

	for _, trueCase := range trueCases {
		playListRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&model.PlayList{}) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.PlayList).UserID != trueCase.UserID {
				t.Fatal("UserID Not Match.")
			}
			if data.(*model.PlayList).Name != trueCase.PlayListName {
				t.Fatal("PlayListName Not Match.")
			}
			if data.(*model.PlayList).Description != trueCase.PlayListDescription {
				t.Fatal("PlayListDescription Not Match.")
			}
			return nil
		})
		result := createPlayListUsecase.CreatePlayList(&trueCase)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}
}
