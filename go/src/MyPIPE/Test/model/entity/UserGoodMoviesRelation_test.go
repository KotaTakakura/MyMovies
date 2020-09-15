package test

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"testing"
	"fmt"
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

func TestGetById(t *testing.T){
	userRepository := infra.NewUserPersistence()
	users,_ :=userRepository.FindById(model.UserID(1012))
	fmt.Println(users.Movies)
}