package test

import (
	mock_domain_service "MyPIPE/Test/mock/domainService"
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestUserRegister(t *testing.T){


	birthday, _ := time.Parse("2006-01-02", "2000-01-01")

	cases := []model.User{
		{
			Name:	model.UserName("TestTaro"),
			Password: model.UserPassword("password"),
			Token: model.UserToken("Temporary-Register-Token"),
			Birthday: birthday,
		},
	}

	for _, c := range cases {
		ctrl := gomock.NewController(t)

		updatedUser := model.NewUser()
		updatedUser.Token = c.Token
		updatedUser.Email = c.Email
		updatedUser.UpdatedAt = time.Now()

		UserRepository := mock_repository.NewMockUserRepository(ctrl)
		UserRepository.EXPECT().FindByToken(c.Token).Return(updatedUser,nil)

		UserRepository.EXPECT().UpdateUser(updatedUser).DoAndReturn(
			func (user *model.User)error{
				if user.Token != ""{
					t.Error("Token Empty Error.")
				}
				if user.Password != c.Password{
					t.Error("Password Error.")
				}
				if user.Birthday != c.Birthday {
					t.Error("Birthday Error.")
				}
				return nil
			})

		userService := mock_domain_service.NewMockIUserService(ctrl)
		userService.EXPECT().CheckNameExists(c.Name).Return(false)

		UserRegisterUsecase := usecase.NewUserRegister(UserRepository,userService)
		err := UserRegisterUsecase.RegisterUser(&c)
		if err != nil{
			t.Error(err)
		}

		ctrl.Finish()
	}
}
