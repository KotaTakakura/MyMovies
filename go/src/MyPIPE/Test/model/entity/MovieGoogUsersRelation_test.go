package test

import (
	"MyPIPE/domain/model"
	"fmt"
	"testing"
	"time"
)

//func TestRelationBetweenMovieAndGoodUsers(t *testing.T){
//	q := &model.Movie{ID:model.MovieID(1)}
//	infra.ConnectGorm().Preload("GoodUsers").Find(q).QueryExpr()
//	//fmt.Println(q.GoodUsers)
//}

func TestMoviePublic(t *testing.T){
	movie := &model.Movie{
		ID:          model.MovieID(10),
		StoreName:   model.MovieStoreName("test1"),
		DisplayName: model.MovieDisplayName(""),
		Description: model.MovieDescription("test3"),
		UserID:      model.UserID(100),
		Public:      model.MoviePublic(0),
		Status:      model.MovieStatus(0),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	err := movie.ChangePublic(model.MoviePublic(1))
	if err != nil{
		fmt.Println(err.Error())
	}
}