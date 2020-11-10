package demo

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"testing"
	"fmt"
)

func TestFind(t *testing.T){
	repo := infra.NewMovieStatusPersistence()

	result,_ := repo.Find(model.MovieID(50))

	fmt.Println(result)

	result.MovieStatus = model.MovieStatusValue(1)

	_ = repo.Save(result)

}
