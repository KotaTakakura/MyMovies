package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
	"errors"
)

func TestUserTemporaryRegistrationUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cases := []model.User{
		{
			Email: model.NewUserEmail("taro@example.jp"),
		},
	}

	for _, c := range cases {

		testUserRepositoryReturnNil := mock_repository.NewMockUserRepository(ctrl)
		testUserRepositoryReturnNil.EXPECT().FindByEmail(c.Email).Return(nil,nil)
		testUserRepositoryReturnNil.EXPECT().SetUser(&c).Return(nil)
		userTemporaryRegistration := usecase.NewUserTemporaryRegistration(testUserRepositoryReturnNil)
		err1 := userTemporaryRegistration.TemporaryRegister(&c)
		if err1 != nil{
			t.Error(err1)
		}

		testUserRepositoryFindByEmailError := mock_repository.NewMockUserRepository(ctrl)
		testUserRepositoryFindByEmailError.EXPECT().FindByEmail(c.Email).Return(&c,nil)
		userTemporaryRegistration = usecase.NewUserTemporaryRegistration(testUserRepositoryFindByEmailError)
		err2 := userTemporaryRegistration.TemporaryRegister(&c)
		if err2 == nil{
			t.Error(err2)
		}

		testUserRepositorySetUserError := mock_repository.NewMockUserRepository(ctrl)
		testUserRepositorySetUserError.EXPECT().FindByEmail(c.Email).Return(nil,nil)
		testUserRepositorySetUserError.EXPECT().SetUser(&c).Return(errors.New("ERROR!!!"))
		userTemporaryRegistration = usecase.NewUserTemporaryRegistration(testUserRepositorySetUserError)
		err3 := userTemporaryRegistration.TemporaryRegister(&c)
		if err3 == nil{
			t.Error(err3)
		}
	}
}
