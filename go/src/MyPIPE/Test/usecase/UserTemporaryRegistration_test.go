package test

//import (
//	mock_repository "MyPIPE/Test/mock/repository"
//	"MyPIPE/domain/model"
//	"MyPIPE/usecase"
//	"github.com/golang/mock/gomock"
//	"testing"
//	"time"
//)
//
//func TestUserTemporaryRegistrationUsecase(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	cases := []model.User{
//		{
//			Email: model.UserEmail("taro@example.jp"),
//		},
//	}
//
//	for _, c := range cases {
//
//		testUserRepository := mock_repository.NewMockUserRepository(ctrl)
//
//		testUserRepository.EXPECT().FindByEmail(c.Email).Return(nil,nil)
//
//		testUserRepository.EXPECT().
//			SetUser(gomock.Any()).
//			DoAndReturn(func(arg *model.User)error{
//				if arg.Email == c.Email && arg.Birthday == time.Date(1000, 1, 1, 0, 0, 0, 0, time.Local){
//					return nil
//				}
//				t.Fail()
//				return nil
//		})
//
//		userTemporaryRegistration := usecase.NewUserTemporaryRegistration(testUserRepository)
//		err1 := userTemporaryRegistration.TemporaryRegister(&c)
//		if err1 != nil{
//			t.Error(err1)
//		}
//	}
//}
