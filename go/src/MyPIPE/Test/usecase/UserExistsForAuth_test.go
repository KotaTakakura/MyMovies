package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserExistsForAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	userExistsForAuthUsecase := usecase.NewUserExists(userRepository)

	trueCases := []struct {
		email    model.UserEmail
		password string
	}{
		{email: model.UserEmail("Test@Email.com"), password: "TestPasword"},
	}

	for _, trueCase := range trueCases {
		password, _ := model.NewUserPassword(trueCase.password)
		userRepository.EXPECT().FindByEmail(trueCase.email).Return(&model.User{
			Password: password,
			Email:    trueCase.email,
		}, nil)

		user, err := userExistsForAuthUsecase.CheckUserExistsForAuth(trueCase.email, trueCase.password)
		if err != nil || user == nil {
			t.Fatal("Usecase Error.")
		}
	}

	falseCases := []struct {
		email    model.UserEmail
		password string
	}{
		{email: model.UserEmail("Test@Email.com"), password: "WrongTestPasword"},
	}

	for _, falseCase := range falseCases {
		userRepository.EXPECT().FindByEmail(falseCase.email).Return(&model.User{
			Password: model.UserPassword("Password"),
			Email:    falseCase.email,
		}, nil)

		user, _ := userExistsForAuthUsecase.CheckUserExistsForAuth(falseCase.email, falseCase.password)
		if user != nil {
			t.Fatal("Usecase Error.")
		}
	}
}
