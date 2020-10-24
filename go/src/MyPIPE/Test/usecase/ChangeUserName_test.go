package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestChangeUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_repository.NewMockUserRepository(ctrl)
	changeUserNameUsecase := usecase.NewChangeUserName(userRepository)

	trueCases := []usecase.ChangeUserNameDTO{
		usecase.ChangeUserNameDTO{
			UserID:   model.UserID(10),
			UserName: model.UserName("TestUserName"),
		},
	}

	for _, trueCase := range trueCases {
		userRepository.EXPECT().FindById(trueCase.UserID).Return(&model.User{
			ID:   trueCase.UserID,
			Name: model.UserName("oldName"),
		}, nil)

		userRepository.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(func(data interface{}) error {
			if reflect.TypeOf(data) != reflect.TypeOf(&model.User{}) {
				t.Fatal("Type Not Match.")
			}
			if data.(*model.User).ID != trueCase.UserID {
				t.Fatal("UserID Not Match.")
			}
			if data.(*model.User).Name != trueCase.UserName {
				t.Fatal("UserName Not Match.")
			}
			return nil
		})

		result := changeUserNameUsecase.ChangeUserName(&trueCase)
		if result != nil {
			t.Fatal("Usecase Error.")
		}
	}
}
