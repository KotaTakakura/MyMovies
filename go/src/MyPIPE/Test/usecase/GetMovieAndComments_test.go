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

func TestGetMovieAndComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentQueryService := mock_queryService.NewMockCommentQueryService(ctrl)
	getMovieAndCommentsUsecase := usecase.NewGetMovieAndComments(commentQueryService)

	trueCases := []usecase.MovieAndGetCommentsDTO{
		usecase.MovieAndGetCommentsDTO{
			MovieID: model.MovieID(10),
		},
	}

	for _, trueCase := range trueCases {

		usecaseResult := queryService.FindByMovieIdDTO{
			Movie: queryService.MovieForFindByMovieIdDTO{
				ID:          uint64(trueCase.MovieID),
				UserID:      10,
				DisplayName: "TestDisplayName",
				Description: "TestDescription",
			},
			MovieLikeCount: 100,
			User: queryService.UserForFindByMovieIdDTO{
				ID:   10,
				Name: "TestUserName",
			},
			Comments: []queryService.CommentsFoundByMovieId{
				queryService.CommentsFoundByMovieId{
					CommentID:   10,
					CommentBody: "TestComment10",
					UserName:    "TestUserName20",
					UserID:      20,
					MovieID:     uint64(trueCase.MovieID),
				},
				queryService.CommentsFoundByMovieId{
					CommentID:   20,
					CommentBody: "TestComment20",
					UserName:    "TestUserName30",
					UserID:      30,
					MovieID:     uint64(trueCase.MovieID),
				},
				queryService.CommentsFoundByMovieId{
					CommentID:   30,
					CommentBody: "TestComment30",
					UserName:    "TestUserName40",
					UserID:      40,
					MovieID:     uint64(trueCase.MovieID),
				},
			},
		}

		checkUsecaseResult := usecaseResult

		commentQueryService.EXPECT().FindByMovieId(trueCase.MovieID).Return(usecaseResult)

		result := getMovieAndCommentsUsecase.Get(&trueCase)

		if !reflect.DeepEqual(result, checkUsecaseResult) {
			t.Fatal("Usecase Error.")
		}
	}
}
