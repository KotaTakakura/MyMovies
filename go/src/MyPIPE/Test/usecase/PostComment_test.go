package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
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
		UserRepository := mock_repository.NewMockUserRepository(ctrl)
		UserRepository.EXPECT().FindById(c.UserID).Return(&model.User{ID: c.UserID},nil)
		UserRepository.EXPECT().UpdateUser(gomock.Any()).
			DoAndReturn(func(user *model.User) error{
				if user.ID != c.UserID{
					t.Error("ID Doesn't Match.")
				}

				if !reflect.DeepEqual(user.CommentsToAppend[0],c){
					t.Error("Comment Struct Invalid.")
				}

				return nil
		})

		postCommentUsecase := usecase.NewPostComment(UserRepository)
		err := postCommentUsecase.PostComment(c)
		if err != nil{
			t.Error("Comment Post Failed.")
		}
	}
}
