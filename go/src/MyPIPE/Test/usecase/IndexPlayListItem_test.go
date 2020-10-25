package test

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestIndexPlayListItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trueCases := []struct {
		indexPlayListMovieInMyPageUsecaseDTO      usecase.IndexPlayListItemInMyPageDTO       //引数
		indexPlayListMovieInMyPageQueryServiceDTO queryService.IndexPlayListMovieInMyPageDTO //戻り値
	}{
		{
			indexPlayListMovieInMyPageUsecaseDTO: usecase.IndexPlayListItemInMyPageDTO{
				UserID:     model.UserID(10),
				PlayListID: model.PlayListID(10),
			},
			indexPlayListMovieInMyPageQueryServiceDTO: queryService.IndexPlayListMovieInMyPageDTO{
				PlayList: queryService.PlayListForIndexPlayListMovieInMyPageDTO{
					PlayListID:          10,
					PlayListName:        "PlayListName10",
					PlaylistDescription: "PlayListDescription10",
				},
				PlayListMovies: []queryService.PlayListMovieForIndexPlayListMovieInMyPageDTO{
					queryService.PlayListMovieForIndexPlayListMovieInMyPageDTO{
						MovieID:            10,
						MovieTitle:         "MovieTitle10",
						MovieDescription:   "MovieDescription10",
						MovieThumbnailName: "MovieThumbnailName10",
						Order:              1,
					},
				},
			},
		},
	}

	for _, trueCase := range trueCases {
		indexPlayListItemQueryService := mock_queryService.NewMockIndexPlayListMovieQueryService(ctrl)
		indexPlayListItemUsecase := usecase.NewIndexPlayListItemInMyPage(indexPlayListItemQueryService)

		checkUsecaseReturn := trueCase.indexPlayListMovieInMyPageQueryServiceDTO

		indexPlayListItemQueryService.EXPECT().Find(trueCase.indexPlayListMovieInMyPageUsecaseDTO.UserID, trueCase.indexPlayListMovieInMyPageUsecaseDTO.PlayListID).Return(&trueCase.indexPlayListMovieInMyPageQueryServiceDTO)

		result := indexPlayListItemUsecase.Find(&trueCase.indexPlayListMovieInMyPageUsecaseDTO)

		if reflect.DeepEqual(result, checkUsecaseReturn) {
			t.Fatal("Usecase Error.")
		}
	}
}
