package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestUserRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trueCases := []struct {
		User model.User
	}{
		{
			User: model.User{
				ID:               model.UserID(10),
				Name:             model.UserName("UserName"),
				Password:         model.UserPassword("userPassword"),
				Email:            model.UserEmail("user@email.com"),
				Birthday:         time.Now(),
				ProfileImageName: "",
				Token:            model.UserToken("123456-7891011"),
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
		},
	}

	for _, trueCase := range trueCases {
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRegisterUsecase := usecase.NewUserRegister(userRepository)
		userRepository.EXPECT().FindByToken(trueCase.User.Token).Return(&trueCase.User, nil)

		userRepository.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&(model.User{})) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.User).ID != trueCase.User.ID {
				t.Fatal("UserID Not Match,")
			}
			if data.(*model.User).Name != trueCase.User.Name {
				t.Fatal("Name Not Match,")
			}
			if data.(*model.User).Email != trueCase.User.Email {
				t.Fatal("Email Not Match,")
			}
			if data.(*model.User).Birthday != trueCase.User.Birthday {
				t.Fatal("Birthday Not Match,")
			}
			if data.(*model.User).Password != trueCase.User.Password {
				t.Fatal("Password Not Match,")
			}
			if data.(*model.User).Token != model.UserToken("") {
				t.Fatal("Token Error,")
			}
			if data.(*model.User).CreatedAt != trueCase.User.CreatedAt {
				t.Fatal("CreatedAt Error,")
			}
			return nil
		})

		result := userRegisterUsecase.RegisterUser(&trueCase.User)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestUserRegister_UserRepository_FindByToken_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		User model.User
	}{
		{
			User: model.User{
				ID:               model.UserID(10),
				Name:             model.UserName("UserName"),
				Password:         model.UserPassword("userPassword"),
				Email:            model.UserEmail("user@email.com"),
				Birthday:         time.Now(),
				ProfileImageName: "",
				Token:            model.UserToken("123456-7891011"),
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
		},
	}

	for _, Case := range cases {
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRegisterUsecase := usecase.NewUserRegister(userRepository)
		userRepository.EXPECT().FindByToken(Case.User.Token).Return(nil, errors.New("ERROR"))

		result := userRegisterUsecase.RegisterUser(&Case.User)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}

	for _, Case := range cases {
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRegisterUsecase := usecase.NewUserRegister(userRepository)
		userRepository.EXPECT().FindByToken(Case.User.Token).Return(nil, nil)

		result := userRegisterUsecase.RegisterUser(&Case.User)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestUserRegister_User_Register_Expired_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trueCases := []struct {
		User model.User
	}{
		{
			User: model.User{
				ID:               model.UserID(10),
				Name:             model.UserName("UserName"),
				Password:         model.UserPassword("userPassword"),
				Email:            model.UserEmail("user@email.com"),
				Birthday:         time.Now(),
				ProfileImageName: "",
				Token:            model.UserToken("123456-7891011"),
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
		},
	}

	for _, trueCase := range trueCases {
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRegisterUsecase := usecase.NewUserRegister(userRepository)
		userRepository.EXPECT().FindByToken(trueCase.User.Token).Return(&model.User{
			ID:        model.UserID(10),
			Email:     model.UserEmail("user@email.com"),
			Birthday:  time.Now(),
			Token:     model.UserToken("123456-7891011"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Date(2001, 5, 20, 23, 59, 59, 0, time.Local),
		}, nil)

		result := userRegisterUsecase.RegisterUser(&trueCase.User)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestUserRegister_UserRepository_UpdateUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trueCases := []struct {
		User model.User
	}{
		{
			User: model.User{
				ID:               model.UserID(10),
				Name:             model.UserName("UserName"),
				Password:         model.UserPassword("userPassword"),
				Email:            model.UserEmail("user@email.com"),
				Birthday:         time.Now(),
				ProfileImageName: "",
				Token:            model.UserToken("123456-7891011"),
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			},
		},
	}

	for _, trueCase := range trueCases {
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRegisterUsecase := usecase.NewUserRegister(userRepository)
		userRepository.EXPECT().FindByToken(trueCase.User.Token).Return(&trueCase.User, nil)

		userRepository.EXPECT().UpdateUser(gomock.Any()).Return(errors.New("ERROR"))

		result := userRegisterUsecase.RegisterUser(&trueCase.User)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}
