package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestUserRegister(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	birthday, _ := time.Parse("2006-01-02", "2000-01-01")

	cases := []model.User{
		{
			Name:	model.UserName("TestTaro"),
			Password: model.UserPassword("password"),
			Token: model.UserToken("ewpiurq-fejas-faes-fae"),
			Birthday: birthday,
		},
	}

	for _, c := range cases {

		updatedUser := model.NewUser()
		updatedUser.Token = c.Token
		updatedUser.Email = c.Email
		updatedUser.UpdatedAt = time.Now()

		UserRepository := mock_repository.NewMockUserRepository(ctrl)
		UserRepository.EXPECT().FindByToken(c.Token).Return(updatedUser,nil)

		UserRepository.EXPECT().
			UpdateUser(updatedUser).Return(nil)

		UserRegisterUsecase := usecase.NewUserRegister(UserRepository)
		err := UserRegisterUsecase.RegisterUser(&c)
		if err != nil{
			t.Error(err)
		}
	}
}
