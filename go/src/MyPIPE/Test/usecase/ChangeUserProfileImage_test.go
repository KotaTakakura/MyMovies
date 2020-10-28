package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"mime/multipart"
	"os"
	"reflect"
	"testing"
)

func TestChangeUserProfileImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	userProfileImageRepository := mock_repository.NewMockUserProfileImageRepository(ctrl)
	changeUserProfileImageUsecase := usecase.NewChangeUserProfileImage(userRepository, userProfileImageRepository)

	cases := []struct {
		FilePath string
	}{
		{
			FilePath: "./../TestProfileImage.jpg",
		},
	}

	for _, Case := range cases {
		file, fileErr := os.Open(Case.FilePath)
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}

		changeUserProfileImageDTO := &usecase.ChangeUserProfileImageDTO{
			UserID: 10,
			ProfileImage: model.UserProfileImage{
				Name: "TestThumbnail.jpg",
				File: file,
			},
		}

		userRepository.EXPECT().FindById(changeUserProfileImageDTO.UserID).Return(&model.User{
			ID:               changeUserProfileImageDTO.UserID,
			Name:             "Name10",
			Password:         "Password10",
			Email:            "email@test.jp",
			ProfileImageName: "OldProfileImageName",
			Token:            "",
		}, nil)

		userProfileImageRepository.EXPECT().Upload(gomock.Any(), gomock.Any()).DoAndReturn(func(data1 interface{}, data2 interface{}) error {
			if reflect.DeepEqual(data1.(multipart.File), &changeUserProfileImageDTO.ProfileImage.File) {
				t.Fatal("File Not Match.")
			}
			if reflect.TypeOf(data2) != reflect.TypeOf(&model.User{}) {
				t.Fatal("Type User Not Match.")
			}
			if data2.(*model.User).ID != changeUserProfileImageDTO.UserID {
				t.Fatal("UserID Not Match.")
			}
			if data2.(*model.User).ProfileImageName != changeUserProfileImageDTO.ProfileImage.Name {
				t.Fatal("Profile Image Name Not Match.")
			}
			return nil
		})

		userRepository.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&model.User{}) {
				t.Fatal("Type User Not Match.")
			}
			if data.(*model.User).ID != changeUserProfileImageDTO.UserID {
				t.Fatal("UserID Not Match.")
			}
			if data.(*model.User).ProfileImageName != changeUserProfileImageDTO.ProfileImage.Name {
				t.Fatal("Profile Image Name Not Match.")
			}
			return nil
		})

		result := changeUserProfileImageUsecase.ChangeUserProfileImage(changeUserProfileImageDTO)

		if result != nil {
			t.Fatal("Usecase Error")
		}
	}
}

func TestChangeUserProfileImage_UserRepository_FindById_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	userProfileImageRepository := mock_repository.NewMockUserProfileImageRepository(ctrl)
	changeUserProfileImageUsecase := usecase.NewChangeUserProfileImage(userRepository, userProfileImageRepository)

	cases := []struct {
		FilePath string
	}{
		{
			FilePath: "./../TestProfileImage.jpg",
		},
	}

	for _, Case := range cases {
		file, fileErr := os.Open(Case.FilePath)
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}

		changeUserProfileImageDTO := &usecase.ChangeUserProfileImageDTO{
			UserID: 10,
			ProfileImage: model.UserProfileImage{
				Name: "TestThumbnail.jpg",
				File: file,
			},
		}

		userRepository.EXPECT().FindById(changeUserProfileImageDTO.UserID).Return(nil, errors.New("ERROR"))

		result := changeUserProfileImageUsecase.ChangeUserProfileImage(changeUserProfileImageDTO)

		if result == nil {
			t.Fatal("Usecase Error")
		}
	}
}

func TestChangeUserProfileImage_UserProfileImageRepository_Upload_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	userProfileImageRepository := mock_repository.NewMockUserProfileImageRepository(ctrl)
	changeUserProfileImageUsecase := usecase.NewChangeUserProfileImage(userRepository, userProfileImageRepository)

	cases := []struct {
		FilePath string
	}{
		{
			FilePath: "./../TestProfileImage.jpg",
		},
	}

	for _, Case := range cases {
		file, fileErr := os.Open(Case.FilePath)
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}

		changeUserProfileImageDTO := &usecase.ChangeUserProfileImageDTO{
			UserID: 10,
			ProfileImage: model.UserProfileImage{
				Name: "TestThumbnail.jpg",
				File: file,
			},
		}

		userRepository.EXPECT().FindById(changeUserProfileImageDTO.UserID).Return(&model.User{
			ID:               changeUserProfileImageDTO.UserID,
			Name:             "Name10",
			Password:         "Password10",
			Email:            "email@test.jp",
			ProfileImageName: "OldProfileImageName",
			Token:            "",
		}, nil)

		userProfileImageRepository.EXPECT().Upload(gomock.Any(), gomock.Any()).Return(errors.New("ERROR"))

		result := changeUserProfileImageUsecase.ChangeUserProfileImage(changeUserProfileImageDTO)

		if result == nil {
			t.Fatal("Usecase Error")
		}
	}
}

func TestChangeUserProfileImage_UserRepository_UpdateUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	userProfileImageRepository := mock_repository.NewMockUserProfileImageRepository(ctrl)
	changeUserProfileImageUsecase := usecase.NewChangeUserProfileImage(userRepository, userProfileImageRepository)

	cases := []struct {
		FilePath string
	}{
		{
			FilePath: "./../TestProfileImage.jpg",
		},
	}

	for _, Case := range cases {
		file, fileErr := os.Open(Case.FilePath)
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}

		changeUserProfileImageDTO := &usecase.ChangeUserProfileImageDTO{
			UserID: 10,
			ProfileImage: model.UserProfileImage{
				Name: "TestThumbnail.jpg",
				File: file,
			},
		}

		userRepository.EXPECT().FindById(changeUserProfileImageDTO.UserID).Return(&model.User{
			ID:               changeUserProfileImageDTO.UserID,
			Name:             "Name10",
			Password:         "Password10",
			Email:            "email@test.jp",
			ProfileImageName: "OldProfileImageName",
			Token:            "",
		}, nil)

		userProfileImageRepository.EXPECT().Upload(gomock.Any(), gomock.Any()).DoAndReturn(func(data1 interface{}, data2 interface{}) error {
			if reflect.DeepEqual(data1.(multipart.File), &changeUserProfileImageDTO.ProfileImage.File) {
				t.Fatal("File Not Match.")
			}
			if reflect.TypeOf(data2) != reflect.TypeOf(&model.User{}) {
				t.Fatal("Type User Not Match.")
			}
			if data2.(*model.User).ID != changeUserProfileImageDTO.UserID {
				t.Fatal("UserID Not Match.")
			}
			if data2.(*model.User).ProfileImageName != changeUserProfileImageDTO.ProfileImage.Name {
				t.Fatal("Profile Image Name Not Match.")
			}
			return nil
		})

		userRepository.EXPECT().UpdateUser(gomock.Any()).Return(errors.New("ERROR"))

		result := changeUserProfileImageUsecase.ChangeUserProfileImage(changeUserProfileImageDTO)

		if result == nil {
			t.Fatal("Usecase Error")
		}
	}
}
