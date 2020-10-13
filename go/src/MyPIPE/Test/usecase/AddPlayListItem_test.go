package test

import (
	"testing"
)

func TestAddPlayListItem(t *testing.T){
	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//
	//playListItemAddJson := &usecase.AddPlayListItemAddJson{
	//	PlayListID: model.PlayListID(100),
	//	UserID:     model.UserID(50),
	//	MovieID:    model.MovieID(30),
	//}
	//
	////PlayListRepository Mock
	//PlayListRepository := mock_repository.NewMockPlayListRepository(ctrl)
	//PlayListRepository.EXPECT().FindByID(playListItemAddJson.PlayListID).
	//	Return(&model.PlayList{
	//	ID:            model.PlayListID(100),
	//	UserID:        playListItemAddJson.UserID,
	//	Name:          "TestPlayList",
	//	PlayListItems: []model.MovieID{},
	//	CreatedAt:     time.Time{},
	//	UpdatedAt:     time.Time{},
	//},nil)
	//PlayListRepository.EXPECT().Save(gomock.Any()).DoAndReturn(
	//	func(playList *model.PlayList)error{
	//		if playList.UserID != playListItemAddJson.UserID{
	//			t.Error("UserID Invalid.")
	//		}
	//
	//		for _,v := range playList.PlayListItems{
	//			if v == playListItemAddJson.MovieID{
	//				return nil
	//			}
	//		}
	//
	//		t.Error("Item Doesn't Contained.")
	//		return nil
	//	})
	//
	////PlayListService Mock
	//PlayListService := mock_domain_service.NewMockIPlayListService(ctrl)
	//PlayListService.EXPECT().CanAddItem(playListItemAddJson.MovieID).Return(true)
	//
	//playListItemAddUsecase := usecase.NewAddPlayListItem(PlayListRepository,PlayListService)
	//err := playListItemAddUsecase.AddPlayListItem(*playListItemAddJson)
	//if err != nil{
	//	t.Error("AddPlayListItemUsecase Error.")
	//}

}