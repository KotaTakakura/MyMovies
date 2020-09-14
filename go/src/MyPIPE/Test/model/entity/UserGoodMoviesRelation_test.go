package test

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"testing"
)

func TestRelationBetweenUserAndGoodMovies(t *testing.T) {
	q := &model.User{ID: model.UserID(1012)}
	infra.ConnectGorm().Preload("GoodMovies").Find(q).QueryExpr()
}

func TestRelationBetweenUserAndPlayLists(t *testing.T) {
	q := &model.User{ID: model.UserID(1012)}
	infra.ConnectGorm().Preload("PlayLists").Find(q).QueryExpr()
	//fmt.Println(q.PlayLists)
}

//func TestRelationFollowUsers(t *testing.T){
//	q := &model.User{ID: model.UserID(1012)}
//	infra.ConnectGorm().Preload("Follows").Find(q).QueryExpr()
//	fmt.Println(q.ID)
//	fmt.Println(q.Follows)
//}
