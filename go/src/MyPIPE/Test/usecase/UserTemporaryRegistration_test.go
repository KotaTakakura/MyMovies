package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestUserTemporaryRegistrationUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cases := []model.User{
		{
			Email: model.UserEmail("taro@example.jp"),
		},
	}

	for _, c := range cases {

		testUserRepositoryReturnNil := mock_repository.NewMockUserRepository(ctrl)

		testUserRepositoryReturnNil.EXPECT().FindByEmail(c.Email).Return(nil,nil)

		testUserRepositoryReturnNil.EXPECT().
			SetUser(gomock.Any()).
			DoAndReturn(func(arg *model.User)error{
				if arg.Email == c.Email && arg.Birthday == time.Date(1000, 1, 1, 0, 0, 0, 0, time.Local){
					return nil
				}
				t.Fail()
				return nil
		})

		userTemporaryRegistration := usecase.NewUserTemporaryRegistration(testUserRepositoryReturnNil)
		err1 := userTemporaryRegistration.TemporaryRegister(&c)
		if err1 != nil{
			t.Error(err1)
		}
	}
}
