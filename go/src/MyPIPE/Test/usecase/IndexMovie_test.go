package test

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
	"MyPIPE/domain/queryService"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestIndexMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trueCases := []struct {
		numberOfIndexMovieQueryServiceAllMethodCalled    int
		numberOfIndexMovieQueryServiceSearchMethodCalled int
		data                                             usecase.IndexMovieSearchDTO
	}{
		{
			numberOfIndexMovieQueryServiceAllMethodCalled:    0,
			numberOfIndexMovieQueryServiceSearchMethodCalled: 1,
			data: usecase.IndexMovieSearchDTO{
				Page:    queryService.IndexMovieQueryServicePage(5),
				KeyWord: "TestKeyword",
				Order:   queryService.IndexMovieQueryServiceOrder("desc"),
			},
		},
		{
			numberOfIndexMovieQueryServiceAllMethodCalled:    1,
			numberOfIndexMovieQueryServiceSearchMethodCalled: 0,
			data: usecase.IndexMovieSearchDTO{
				Page:    queryService.IndexMovieQueryServicePage(5),
				KeyWord: "",
				Order:   queryService.IndexMovieQueryServiceOrder("desc"),
			},
		},
	}

	for _, trueCase := range trueCases {

		usecaseReturn := queryService.IndexMovieDTO{
			Movies: []queryService.MoviesForIndexMovieDTO{
				queryService.MoviesForIndexMovieDTO{
					MovieID:          10,
					MovieDisplayName: "TestMovieDisplayName10",
					UserID:           10,
					UserName:         "TestUserName10",
					ThumbnailName:    "TestThumbnailName10",
				},
			},
			Count: 10,
		}
		usecaseReturnCheck := usecaseReturn

		indexMovieQueryService := mock_queryService.NewMockIndexMovieQueryService(ctrl)
		indexMovieUsecase := usecase.NewIndexMovie(indexMovieQueryService)

		indexMovieQueryService.EXPECT().All(trueCase.data.Page, trueCase.data.Order).Return(usecaseReturn).Times(trueCase.numberOfIndexMovieQueryServiceAllMethodCalled)
		indexMovieQueryService.EXPECT().Search(trueCase.data.Page, trueCase.data.KeyWord, trueCase.data.Order).Return(usecaseReturn).Times(trueCase.numberOfIndexMovieQueryServiceSearchMethodCalled)

		result := indexMovieUsecase.Search(&trueCase.data)
		if !reflect.DeepEqual(result, usecaseReturnCheck) {
			t.Fatal("Usecase Error")
		}

	}
}
