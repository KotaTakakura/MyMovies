package test

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"fmt"
	"testing"
)

func TestGetPlayList(t *testing.T) {
	playList := &model.PlayList{
		ID: 1,
	}
	infra.ConnectGorm().Find(&playList).QueryExpr()
	//fmt.Println(playList)
}

func TestGetPlayListItems(t *testing.T) {
	playList := &model.PlayList{
		ID: 1,
	}
	infra.ConnectGorm().Preload("PlayListItems").Find(&playList).QueryExpr()
	fmt.Println(playList)
}
