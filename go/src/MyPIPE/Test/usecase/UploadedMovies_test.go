package test

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
	"MyPIPE/domain/model"
	queryService "MyPIPE/domain/queryService/UploadedMovies"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestUploadedMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trueCases := []struct {
		UserID         model.UserID
		UploadedMovies []queryService.UploadedMoviesDTO
	}{
		{
			UserID: model.UserID(10),
			UploadedMovies: []queryService.UploadedMoviesDTO{
				queryService.UploadedMoviesDTO{
					MovieID:            10,
					MovieName:          "MovieName10",
					MovieDescription:   "MovieDescription10",
					MovieStatus:        0,
					MovieThumbnailName: "MovieThumbnailName10",
					MoviePublic:        0,
				},
				queryService.UploadedMoviesDTO{
					MovieID:            10,
					MovieName:          "MovieName10",
					MovieDescription:   "MovieDescription10",
					MovieStatus:        0,
					MovieThumbnailName: "MovieThumbnailName10",
					MoviePublic:        0,
				},
			},
		},
	}

	for _, trueCase := range trueCases {
		uploadedMoviesQueryService := mock_queryService.NewMockUploadedMovies(ctrl)
		uploadedMoviesUsecase := usecase.NewUploadedMovies(uploadedMoviesQueryService)

		uploadedMoviesQueryService.EXPECT().Get(trueCase.UserID).Return(trueCase.UploadedMovies)

		checkUsecaseResult := trueCase.UploadedMovies

		result := uploadedMoviesUsecase.Get(trueCase.UserID)
		if !reflect.DeepEqual(result, checkUsecaseResult) {
			t.Fatal("Usecase Error.")
		}
	}
}
