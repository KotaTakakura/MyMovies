package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestDeletePlayList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
	deletePlayListUsecase := usecase.NewDeletePlayList(playListRepository)

	trueCases := []usecase.DeletePlayListDTO{
		usecase.DeletePlayListDTO{
			UserID:     model.UserID(10),
			PlaylistID: model.PlayListID(20),
		},
	}

	for _, trueCase := range trueCases {
		playListRepository.EXPECT().Remove(trueCase.UserID, trueCase.PlaylistID).Return(nil)

		result := deletePlayListUsecase.Delete(&trueCase)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestDeletePlayList_PlayListRepository_Remove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListRepository := mock_repository.NewMockPlayListRepository(ctrl)
	deletePlayListUsecase := usecase.NewDeletePlayList(playListRepository)

	trueCases := []usecase.DeletePlayListDTO{
		usecase.DeletePlayListDTO{
			UserID:     model.UserID(10),
			PlaylistID: model.PlayListID(20),
		},
	}

	for _, trueCase := range trueCases {
		playListRepository.EXPECT().Remove(trueCase.UserID, trueCase.PlaylistID).Return(errors.New("ERROR"))

		result := deletePlayListUsecase.Delete(&trueCase)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}
