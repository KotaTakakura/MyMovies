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

func TestIndexPlayListInMyPage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trueCases := []struct {
		indexPlayListsInMyPageUsecaseDTO      usecase.IndexPlayListsInMyPageDTO
		indexPlayListsInMyPageQueryServiceDTO queryService.IndexPlayListsInMyPageDTO
	}{
		{
			indexPlayListsInMyPageUsecaseDTO: usecase.IndexPlayListsInMyPageDTO{
				UserID: model.UserID(10),
			},
			indexPlayListsInMyPageQueryServiceDTO: queryService.IndexPlayListsInMyPageDTO{
				PlayLists: []queryService.PlayListForIndexPlayListsInMyPageDTO{
					queryService.PlayListForIndexPlayListsInMyPageDTO{
						PlayListID:                      10,
						PlayListName:                    "PlayListName10",
						PlayListDescription:             "PlayListDescription10",
						PlayListFirstMovieID:            10,
						PlayListFirstMovieThumbnailName: "PlayListThumbNailName10",
					},
					queryService.PlayListForIndexPlayListsInMyPageDTO{
						PlayListID:                      20,
						PlayListName:                    "PlayListName20",
						PlayListDescription:             "PlayListDescription20",
						PlayListFirstMovieID:            20,
						PlayListFirstMovieThumbnailName: "PlayListThumbNailName20",
					},
				},
				PlayListsCount: 2,
			},
		},
	}

	for _, trueCase := range trueCases {
		indexPlayListsInMyPageQueryService := mock_queryService.NewMockIndexPlayListsInMyPageQueryService(ctrl)
		indexPlayListsInMyPageUsecase := usecase.NewIndexPlayListsInMyPage(indexPlayListsInMyPageQueryService)
		checkUsecaseResult := trueCase.indexPlayListsInMyPageQueryServiceDTO
		indexPlayListsInMyPageQueryService.EXPECT().All(trueCase.indexPlayListsInMyPageUsecaseDTO.UserID).Return(&trueCase.indexPlayListsInMyPageQueryServiceDTO)

		result := indexPlayListsInMyPageUsecase.All(&trueCase.indexPlayListsInMyPageUsecaseDTO)
		if reflect.DeepEqual(result, checkUsecaseResult) {
			t.Fatal("Usecase Error.")
		}
	}
}
