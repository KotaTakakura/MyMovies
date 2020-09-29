package test

import (
	mock_repository "MyPIPE/Test/mock/repository"
	"MyPIPE/domain/model"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestFollowUser(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := usecase.NewFollowDTO(model.UserID(100),model.UserID(200))

	userRepoistory := mock_repository.NewMockUserRepository(ctrl)
	followUserRepository := mock_repository.NewMockFollowUserRepository(ctrl)

	findUserByIdFirst := userRepoistory.EXPECT().FindById(cases.UserID).Return(&model.User{
		ID:         cases.UserID,
	},nil)
	findUserByIdSecond := userRepoistory.EXPECT().FindById(cases.FollowID).Return(&model.User{
		ID:         cases.FollowID,
	},nil)

	findFollowUser := followUserRepository.EXPECT().FindByUserIdAndFollowId(cases.UserID,cases.FollowID).Return(nil)
	saveFollowUser := followUserRepository.
		EXPECT().
		Save(gomock.Any()).
		DoAndReturn(
			func(followUser *model.FollowUser)error{
				if followUser.UserID != cases.UserID{
					t.Error("Invalid UserID.")
				}
				if followUser.FollowID != cases.FollowID{
					t.Error("Invalid FollowID.")
				}
				return nil
			})

	gomock.InOrder(
		findFollowUser,
		findUserByIdFirst,
		findUserByIdSecond,
		saveFollowUser,
	)

	followUserUsecase := usecase.NewFollowUser(userRepoistory,followUserRepository)
	err := followUserUsecase.Follow(cases)
	if err != nil{
		t.Error("FollowUserUsecase Error.")
	}
}