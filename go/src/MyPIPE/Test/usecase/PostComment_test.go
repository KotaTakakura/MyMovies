package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestPostComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []model.Comment{
		model.Comment{
			Body:    model.CommentBody("こんにちは！！！"),
			MovieID: model.MovieID(100),
			UserID:  model.UserID(1010),
		},
		model.Comment{
			Body:    model.CommentBody("Good Morning！！！"),
			MovieID: model.MovieID(11),
			UserID:  model.UserID(1200),
		},
	}

	for _, c := range cases {
		CommentRepository := mock_repository.NewMockCommentRepository(ctrl)
		CommentRepository.EXPECT().Save(&c).Return(nil)

		MovieRepository := mock_repository.NewMockMovieRepository(ctrl)
		MovieRepository.EXPECT().FindById(c.MovieID).Return(&model.Movie{
			ID:          c.MovieID,
			StoreName:   "StoreNameTest",
			DisplayName: "DisplayNameTest",
			UserID:      c.UserID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}, nil)

		postCommentUsecase := usecase.NewPostComment(CommentRepository, MovieRepository)
		err := postCommentUsecase.PostComment(c)
		if err != nil {
			t.Error("PostComment Usecase Test Failed.")
		}
	}
}
