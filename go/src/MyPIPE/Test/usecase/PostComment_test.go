package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPostComment(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []model.Comment{
		model.Comment{
			Body:	model.CommentBody("こんにちは！！！"),
			MovieID: model.MovieID(100),
			UserID: model.UserID(1010),
		},
		model.Comment{
			Body:	model.CommentBody("おはよう！！！"),
			MovieID: model.MovieID(11),
			UserID: model.UserID(1011110),
		},
	}

	for _, c := range cases {
		CommentRepository := mock_repository.NewMockCommentRepository(ctrl)
		CommentRepository.EXPECT().Save(&c).Return(nil)

		MovieRepository := mock_repository.NewMockMovieRepository(ctrl)
		MovieRepository.EXPECT().FindById(c.MovieID).Return(nil,nil)

		postCommentUsecase := usecase.NewPostComment(CommentRepository,MovieRepository)
		err := postCommentUsecase.PostComment(c)
		if err != nil{
			t.Error("Comment Post Failed.")
		}
	}
}
