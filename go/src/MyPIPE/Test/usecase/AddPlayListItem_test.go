package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestAddPlayListItem(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playListItemAddJson := &usecase.PlayListItemAddJson{
		PlayListID: 100,
		UserID:     50,
		MovieID:    30,
	}

	//MovieRepository Mock
	MovieRepository := mock_repository.NewMockMovieRepository(ctrl)
	MovieRepository.EXPECT().FindById(model.MovieID(playListItemAddJson.MovieID)).Return(&model.Movie{},nil)

	//PlayListRepository Mock
	PlayListRepository := mock_repository.NewMockPlayListRepository(ctrl)
	PlayListRepository.EXPECT().FindByID(model.PlayListID(playListItemAddJson.PlayListID)).
		Return(&model.PlayList{
		ID:            model.PlayListID(100),
		UserID:        model.UserID(playListItemAddJson.UserID),
		Name:          "TestPlayList",
		PlayListItems: []model.MovieID{},
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	},nil)
	PlayListRepository.EXPECT().Save(gomock.Any()).DoAndReturn(
		func(playList *model.PlayList)error{
			if playList.UserID != model.UserID(playListItemAddJson.UserID){
				t.Error("UserID Invalid.")
			}

			for _,v := range playList.PlayListItems{
				if v == model.MovieID(playListItemAddJson.MovieID){
					return nil
				}
			}

			t.Error("Item Doesn't Contained.")
			return nil
		})

	playListItemAddUsecase := usecase.NewAddPlayListItem(PlayListRepository,MovieRepository)
	err := playListItemAddUsecase.AddPlayListItem(*playListItemAddJson)
	if err != nil{
		t.Error("AddPlayListItemUsecase Error.")
	}

}