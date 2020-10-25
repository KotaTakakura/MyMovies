package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestUserTemporaryRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//仮登録・本登録なし
	trueCases := []struct {
		User model.User
	}{
		{
			User: model.User{
				ID:               10,
				Name:             model.UserName(""),
				Password:         model.UserPassword(""),
				Email:            model.UserEmail(""),
				Birthday:         time.Time{},
				ProfileImageName: "",
				Token:            model.UserToken("1234-5678-910"),
			},
		},
	}

	for _, trueCase := range trueCases {
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userTemporaryRegisterUsecase := usecase.NewUserTemporaryRegistration(userRepository)

		userRepository.EXPECT().FindByEmail(trueCase.User.Email).Return(nil, nil)
		userRepository.EXPECT().SetUser(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.User{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.User).Email != trueCase.User.Email {
				t.Fatal("Email Not Match,")
			}
			if data.(*model.User).Birthday != time.Date(1000, 1, 1, 0, 0, 0, 0, time.Local) {
				t.Fatal("Birthday Not Match,")
			}
			if data.(*model.User).Token == model.UserToken("") {
				t.Fatal("Token Not Match,")
			}
			return nil
		})

		result := userTemporaryRegisterUsecase.TemporaryRegister(&trueCase.User)

		if result != nil {
			t.Fatal("Usecase Error.")
		}

	}
}
