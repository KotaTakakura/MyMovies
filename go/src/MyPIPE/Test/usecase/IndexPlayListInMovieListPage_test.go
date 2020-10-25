package test

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
	"MyPIPE/domain/queryService"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestIndexPlayListInMovieListPage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trueCases := []struct {
		findDTO                         usecase.FindDTO                              //Usecaseの引数
		indexPlayListInMovieListPageDTO queryService.IndexPlayListInMovieListPageDTO //戻り値
	}{
		{
			findDTO: usecase.FindDTO{
				UserID:  10,
				MovieID: 10,
			},
			indexPlayListInMovieListPageDTO: queryService.IndexPlayListInMovieListPageDTO{
				PlayLists: []queryService.PlayListForIndexPlayListInMovieListPageDTO{
					queryService.PlayListForIndexPlayListInMovieListPageDTO{
						PlayListID:   10,
						PlayListName: "TestPlayListName10",
						PlayListMovies: []queryService.PlayListMoviesForIndexPlayListInMovieListPageDTO{
							queryService.PlayListMoviesForIndexPlayListInMovieListPageDTO{
								MovieID:    10,
								PlayListID: 10,
							},
							queryService.PlayListMoviesForIndexPlayListInMovieListPageDTO{
								MovieID:    20,
								PlayListID: 10,
							},
							queryService.PlayListMoviesForIndexPlayListInMovieListPageDTO{
								MovieID:    30,
								PlayListID: 10,
							},
						},
					},
					queryService.PlayListForIndexPlayListInMovieListPageDTO{
						PlayListID:   20,
						PlayListName: "TestPlayListName20",
						PlayListMovies: []queryService.PlayListMoviesForIndexPlayListInMovieListPageDTO{
							queryService.PlayListMoviesForIndexPlayListInMovieListPageDTO{
								MovieID:    10,
								PlayListID: 10,
							},
							queryService.PlayListMoviesForIndexPlayListInMovieListPageDTO{
								MovieID:    20,
								PlayListID: 10,
							},
							queryService.PlayListMoviesForIndexPlayListInMovieListPageDTO{
								MovieID:    30,
								PlayListID: 10,
							},
						},
					},
				},
			},
		},
	}

	for _, trueCase := range trueCases {
		indexPlayListInMovieListPageQueryService := mock_queryService.NewMockIndexPlayListInMovieListPageQueryService(ctrl)
		indexPlayListInMovieListPageUsecase := usecase.NewIndexPlayListInMovieListPage(indexPlayListInMovieListPageQueryService)
		checkUsecaseReturn := trueCase.indexPlayListInMovieListPageDTO
		indexPlayListInMovieListPageQueryService.EXPECT().Find(trueCase.findDTO.UserID, trueCase.findDTO.MovieID).Return(&trueCase.indexPlayListInMovieListPageDTO)

		result := indexPlayListInMovieListPageUsecase.Find(&trueCase.findDTO)
		if reflect.DeepEqual(result, checkUsecaseReturn) {
			t.Fatal("usecase Error")
		}
	}
}
