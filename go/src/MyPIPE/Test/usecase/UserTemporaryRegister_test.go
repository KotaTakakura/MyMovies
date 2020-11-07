package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"errors"
	"fmt"
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
		temporaryRegisterMailRepository := mock_repository.NewMockTemporaryRegisterMailRepository(ctrl)
		userTemporaryRegisterUsecase := usecase.NewUserTemporaryRegistration(userRepository, temporaryRegisterMailRepository)

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

		temporaryRegisterMailRepository.EXPECT().Send(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.TemporaryRegisterMail{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.TemporaryRegisterMail).To != trueCase.User.Email {
				t.Fatal("Email Not Match,")
			}
			fmt.Println(data.(*model.TemporaryRegisterMail).Token)
			if data.(*model.TemporaryRegisterMail).Token == "" {
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

func TestUserTemporaryRegister_User_Already_Registered(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
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

	for _, Case := range cases {
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		temporaryRegisterMailRepository := mock_repository.NewMockTemporaryRegisterMailRepository(ctrl)
		userTemporaryRegisterUsecase := usecase.NewUserTemporaryRegistration(userRepository, temporaryRegisterMailRepository)

		userRepository.EXPECT().FindByEmail(Case.User.Email).Return(&model.User{
			Token: "",
		}, nil)

		result := userTemporaryRegisterUsecase.TemporaryRegister(&Case.User)

		if result != nil {
			t.Fatal("Usecase Error.")
		}

	}
}

func TestUserTemporaryRegister_User_Already_Temporary_Registered(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
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

	for _, Case := range cases {
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		temporaryRegisterMailRepository := mock_repository.NewMockTemporaryRegisterMailRepository(ctrl)
		userTemporaryRegisterUsecase := usecase.NewUserTemporaryRegistration(userRepository, temporaryRegisterMailRepository)

		userRepository.EXPECT().FindByEmail(Case.User.Email).Return(&model.User{
			ID:    model.UserID(20),
			Token: "1234-5678-910",
			Email: "test@email.com",
		}, nil)

		temporaryRegisterMailRepository.EXPECT().Send(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.TemporaryRegisterMail{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.TemporaryRegisterMail).To != "test@email.com" {
				t.Fatal("Email Not Match,")
			}
			fmt.Println(data.(*model.TemporaryRegisterMail).Token)
			if data.(*model.TemporaryRegisterMail).Token == "" {
				t.Fatal("Token Not Match,")
			}
			return nil
		})

		userRepository.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.User{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.User).ID != model.UserID(20) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*model.User).Token == "1234-5678-910" || data.(*model.User).Token == "" {
				t.Fatal("Token Error,")
			}
			return nil
		})

		result := userTemporaryRegisterUsecase.TemporaryRegister(&Case.User)

		if result != nil {
			t.Fatal("Usecase Error.")
		}

	}

	//ユーザー再仮登録
	for _, Case := range cases {
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		temporaryRegisterMailRepository := mock_repository.NewMockTemporaryRegisterMailRepository(ctrl)
		userTemporaryRegisterUsecase := usecase.NewUserTemporaryRegistration(userRepository, temporaryRegisterMailRepository)

		userRepository.EXPECT().FindByEmail(Case.User.Email).Return(&model.User{
			ID:    model.UserID(20),
			Token: "1234-5678-910",
		}, nil)

		userRepository.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.User{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.User).ID != model.UserID(20) {
				t.Fatal("UserID Not Match,")
			}
			if data.(*model.User).Token == "1234-5678-910" || data.(*model.User).Token == "" {
				t.Fatal("Token Error,")
			}
			return nil
		})

		temporaryRegisterMailRepository.EXPECT().Send(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.TemporaryRegisterMail{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.TemporaryRegisterMail).To != Case.User.Email {
				t.Fatal("Email Not Match,")
			}
			fmt.Println(data.(*model.TemporaryRegisterMail).Token)
			if data.(*model.TemporaryRegisterMail).Token == "" {
				t.Fatal("Token Not Match,")
			}
			return nil
		})

		result := userTemporaryRegisterUsecase.TemporaryRegister(&Case.User)

		if result != nil {
			t.Fatal("Usecase Error.")
		}

	}

	//Updateエラー
	for _, Case := range cases {
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		temporaryRegisterMailRepository := mock_repository.NewMockTemporaryRegisterMailRepository(ctrl)
		userTemporaryRegisterUsecase := usecase.NewUserTemporaryRegistration(userRepository, temporaryRegisterMailRepository)

		userRepository.EXPECT().FindByEmail(Case.User.Email).Return(&model.User{
			ID:    model.UserID(20),
			Token: "1234-5678-910",
		}, nil)

		userRepository.EXPECT().UpdateUser(gomock.Any()).Return(errors.New("ERROR"))

		result := userTemporaryRegisterUsecase.TemporaryRegister(&Case.User)

		if result == nil {
			t.Fatal("Usecase Error.")
		}

	}
}

func TestUserTemporaryRegister_UserRepository_SetUser_Error(t *testing.T) {
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
		temporaryRegisterMailRepository := mock_repository.NewMockTemporaryRegisterMailRepository(ctrl)
		userTemporaryRegisterUsecase := usecase.NewUserTemporaryRegistration(userRepository, temporaryRegisterMailRepository)

		userRepository.EXPECT().FindByEmail(trueCase.User.Email).Return(nil, nil)
		userRepository.EXPECT().SetUser(gomock.Any()).Return(errors.New("ERROR"))

		result := userTemporaryRegisterUsecase.TemporaryRegister(&trueCase.User)

		if result == nil {
			t.Fatal("Usecase Error.")
		}

	}
}

func TestUserTemporaryRegister_SendMail_Error(t *testing.T) {
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
		temporaryRegisterMailRepository := mock_repository.NewMockTemporaryRegisterMailRepository(ctrl)
		userTemporaryRegisterUsecase := usecase.NewUserTemporaryRegistration(userRepository, temporaryRegisterMailRepository)

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

		temporaryRegisterMailRepository.EXPECT().Send(gomock.Any()).Return(errors.New("ERROR"))

		result := userTemporaryRegisterUsecase.TemporaryRegister(&trueCase.User)

		if result == nil {
			t.Fatal("Usecase Error.")
		}

	}
}
