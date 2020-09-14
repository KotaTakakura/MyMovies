package test

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"testing"
)

func TestRelationBetweenMovieAndGoodUsers(t *testing.T){
	q := &model.Movie{ID:model.MovieID(1)}
	infra.ConnectGorm().Preload("GoodUsers").Find(q).QueryExpr()
	//fmt.Println(q.GoodUsers)
}

