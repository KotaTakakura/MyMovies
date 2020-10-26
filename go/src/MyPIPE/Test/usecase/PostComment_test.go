package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestPostComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepository := mock_repository.NewMockCommentRepository(ctrl)
	movieRepository := mock_repository.NewMockMovieRepository(ctrl)
	postCommentUsecase := usecase.NewPostComment(commentRepository, movieRepository)

	trueCases := []usecase.PostCommentDTO{
		usecase.PostCommentDTO{
			UserID:  model.UserID(10),
			MovieID: model.MovieID(20),
			Body:    model.CommentBody("TestComment."),
		},
	}

	for _, trueCase := range trueCases {
		movieRepository.EXPECT().FindById(trueCase.MovieID).Return(&model.Movie{
			ID: trueCase.MovieID,
		}, nil)
		commentRepository.EXPECT().Save(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&model.Comment{}) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.Comment).UserID != trueCase.UserID {
				t.Fatal("UserID Not Match.")
			}
			if data.(*model.Comment).MovieID != trueCase.MovieID {
				t.Fatal("MovieID Not Match.")
			}
			if data.(*model.Comment).Body != trueCase.Body {
				t.Fatal("CommentBody Not Match.")
			}
			return nil
		})

		result := postCommentUsecase.PostComment(&trueCase)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestPostComment_MovieRepository_FindById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []usecase.PostCommentDTO{
		usecase.PostCommentDTO{
			UserID:  model.UserID(10),
			MovieID: model.MovieID(20),
			Body:    model.CommentBody("TestComment."),
		},
	}

	for _, Case := range cases {
		commentRepository := mock_repository.NewMockCommentRepository(ctrl)
		movieRepository := mock_repository.NewMockMovieRepository(ctrl)
		postCommentUsecase := usecase.NewPostComment(commentRepository, movieRepository)

		movieRepository.EXPECT().FindById(Case.MovieID).Return(&model.Movie{
			ID: Case.MovieID,
		}, errors.New("ERROR"))

		result := postCommentUsecase.PostComment(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}

	for _, Case := range cases {
		commentRepository := mock_repository.NewMockCommentRepository(ctrl)
		movieRepository := mock_repository.NewMockMovieRepository(ctrl)
		postCommentUsecase := usecase.NewPostComment(commentRepository, movieRepository)

		movieRepository.EXPECT().FindById(Case.MovieID).Return(nil, nil)

		result := postCommentUsecase.PostComment(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestPostComment_MovieRepository_Save_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []usecase.PostCommentDTO{
		usecase.PostCommentDTO{
			UserID:  model.UserID(10),
			MovieID: model.MovieID(20),
			Body:    model.CommentBody("TestComment."),
		},
	}

	for _, Case := range cases {
		commentRepository := mock_repository.NewMockCommentRepository(ctrl)
		movieRepository := mock_repository.NewMockMovieRepository(ctrl)
		postCommentUsecase := usecase.NewPostComment(commentRepository, movieRepository)

		movieRepository.EXPECT().FindById(Case.MovieID).Return(&model.Movie{
			ID: Case.MovieID,
		}, nil)
		commentRepository.EXPECT().Save(gomock.Any()).Return(errors.New("ERROR"))

		result := postCommentUsecase.PostComment(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}
