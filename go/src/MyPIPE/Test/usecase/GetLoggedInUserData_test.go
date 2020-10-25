package test

import (
	mock_queryService "MyPIPE/Test/mock/queryService"
	"MyPIPE/domain/model"
	"MyPIPE/domain/queryService"
	"MyPIPE/usecase"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestGetLoggedInUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	getLoggedInUserDataQueryService := mock_queryService.NewMockGetLoggedInUserDataQueryService(ctrl)
	getLoggedInUserData := usecase.NewGetLoggedInUserData(getLoggedInUserDataQueryService)

	trueCases := []usecase.GetLoggedInUserDataDTO{
		usecase.GetLoggedInUserDataDTO{
			UserID: model.UserID(10),
		},
	}

	for _, trueCase := range trueCases {

		usecaseReturn := &queryService.GetLoggedInUserDataDTO{
			ID:               uint64(trueCase.UserID),
			Name:             "TestName",
			Email:            "TestEmail@mail.com",
			Birthday:         time.Now(),
			ProfileImageName: "TestProfileImageName",
			CreatedAt:        time.Now(),
			AvatarName:       "TestAvatarName",
		}

		getLoggedInUserDataQueryService.EXPECT().FindByUserId(trueCase.UserID).Return(usecaseReturn)

		result := getLoggedInUserData.Find(&trueCase)

		if result != usecaseReturn {
			t.Fatal("Usecase Error.")
		}
	}
}
