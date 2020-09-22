package test

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"testing"
	"fmt"
)

func TestPlayListRelationWithItems(t *testing.T){
	playListrepo := infra.NewPlayListPersistence()
	//playList := model.NewPlayList()
	playListId,_ := model.NewPlayListID(2)
	playList,_ := playListrepo.FindByID(playListId)
	fmt.Println(playList.PlayListItems)
}
