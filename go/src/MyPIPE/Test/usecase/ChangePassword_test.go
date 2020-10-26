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

func TestChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	changePasswordUsecase := usecase.NewChangePassword(userRepository)

	trueCases := []usecase.ChangePasswordDTO{
		usecase.ChangePasswordDTO{
			UserID:   model.UserID(10),
			Password: model.UserPassword("password"),
		},
	}

	for _, trueCase := range trueCases {
		userRepository.EXPECT().FindById(trueCase.UserID).Return(&model.User{
			ID:       trueCase.UserID,
			Password: "oldPassword",
		}, nil)

		userRepository.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&model.User{}) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.User).Password != trueCase.Password {
				t.Fatal("Password Not Match.")
			}
			return nil
		})

		result := changePasswordUsecase.ChangePassword(&trueCase)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestChangePassword_UserRepository_FindById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	changePasswordUsecase := usecase.NewChangePassword(userRepository)

	cases := []usecase.ChangePasswordDTO{
		usecase.ChangePasswordDTO{
			UserID:   model.UserID(10),
			Password: model.UserPassword("password"),
		},
	}

	for _, Case := range cases {
		userRepository.EXPECT().FindById(Case.UserID).Return(&model.User{
			ID:       Case.UserID,
			Password: "oldPassword",
		}, errors.New("ERROR"))

		result := changePasswordUsecase.ChangePassword(&Case)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}

func TestChangePassword_UserRepository_UpdateUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	changePasswordUsecase := usecase.NewChangePassword(userRepository)

	trueCases := []usecase.ChangePasswordDTO{
		usecase.ChangePasswordDTO{
			UserID:   model.UserID(10),
			Password: model.UserPassword("password"),
		},
	}

	for _, trueCase := range trueCases {
		userRepository.EXPECT().FindById(trueCase.UserID).Return(&model.User{
			ID:       trueCase.UserID,
			Password: "oldPassword",
		}, nil)

		userRepository.EXPECT().UpdateUser(gomock.Any()).Return(errors.New("ERROR"))

		result := changePasswordUsecase.ChangePassword(&trueCase)
		if result == nil {
			t.Fatal("Usecase Error.")
		}
	}
}
