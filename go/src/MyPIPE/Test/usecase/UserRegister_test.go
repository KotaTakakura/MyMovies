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
			Name:	model.NewUserName("TestTaro"),
			Password: model.NewUserPassword("password"),
			Token: model.NewUserToken("ewpiurq-fejas-faes-fae"),
			Birthday: birthday,
		},
	}

	for _, c := range cases {

		updatedUser := model.NewUser()
		updatedUser.Token = c.Token
		updatedUser.Email = c.Email

		UserRepository := mock_repository.NewMockUserRepository(ctrl)
		UserRepository.EXPECT().FindByToken(c.Token).Return(updatedUser,nil)

		UserRepository.EXPECT().
			UpdateUser(updatedUser).
			DoAndReturn(func(updatedUser *model.User)error{
				if updatedUser.Token != "" {
					t.Error("ERROR1")
				}
				if updatedUser.Password != c.Password {
					t.Error("ERROR2")
				}
				if updatedUser.Name != c.Name {
					t.Error("ERROR3")
				}
				if updatedUser.Birthday != c.Birthday {
					t.Error("ERROR4")
				}
				return nil
		})

		UserRegisterUsecase := usecase.NewUserRegister(UserRepository)
		UserRegisterUsecase.RegisterUser(&c)
	}
}
