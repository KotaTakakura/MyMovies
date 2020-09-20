package test

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"testing"
	"fmt"
)

func TestNewEvaluate(t *testing.T){
	user,_ := infra.NewUserPersistence().FindById(model.UserID(1012))
	//fmt.Println(user.GoodMovies)
	//fmt.Println(user.BadMovies)
	evaluate,err :=model.NewEvaluate("good")
	if err != nil{
		fmt.Println(err.Error())
		fmt.Println("~~~~~")
		return
	}
	//user.Evaluate(evaluate,user,model.MovieID(7))
	errr := user.Evaluate(evaluate,model.MovieID(6))
	if errr != nil{
		fmt.Println(errr.Error())
		fmt.Println("|||||")
		return
	}
	fmt.Println(user.GoodMovies)
	fmt.Println(user.BadMovies)
}

//func TestRelationBetweenUserAndGoodMovies(t *testing.T) {
//	q := &model.User{ID: model.UserID(1012)}
//	infra.ConnectGorm().Preload("GoodMovies").Find(q).QueryExpr()
//}
//
//func TestRelationBetweenUserAndPlayLists(t *testing.T) {
//	q := &model.User{ID: model.UserID(1012)}
//	infra.ConnectGorm().Preload("PlayLists").Find(q).QueryExpr()
//	//fmt.Println(q.PlayLists)
//}
//
//func TestGetById(t *testing.T){
//	userRepository := infra.NewUserPersistence()
//	users,_ :=userRepository.FindById(model.UserID(1012))
//	fmt.Println(users.Comments)
//}

//func TestSetGoodMovie(t *testing.T){
//	user,_ := infra.NewUserPersistence().FindById(model.UserID(1012))
//	_ = user.SetGoodMovie(model.MovieID(20))
	//fmt.Println("||||||||||||||")
	//fmt.Println(user.GoodMovies)
	//fmt.Println("||||||||||||||")
	//fmt.Println(user.BadMovies)
	//fmt.Println("||||||||||||||")
	//_ = infra.NewUserPersistence().UpdateUser(user)
//}
//
//func TestSetBadMovie(t *testing.T){
//	user,_ := infra.NewUserPersistence().FindById(model.UserID(1012))
//	_ = user.SetBadMovie(model.MovieID(20))
//	fmt.Println("||||||||||||||")
//	fmt.Println(user.GoodMovies)
//	fmt.Println("||||||||||||||")
//	fmt.Println(user.BadMovies)
//	fmt.Println("||||||||||||||")
//	//_ = infra.NewUserPersistence().UpdateUser(user)
//}

//func TestUnsetGoodMovie(t *testing.T){
//	user,_ := infra.NewUserPersistence().FindById(model.UserID(1012))
//	_ = user.UnsetGoodMovie(model.MovieID(20))
//	fmt.Println("||||||||||||||")
//	fmt.Println(user.GoodMovies)
//	fmt.Println("||||||||||||||")
//	fmt.Println(user.BadMovies)
//	fmt.Println("||||||||||||||")
	//_ = infra.NewUserPersistence().UpdateUser(user)
//}


//func TestUnsetBadMovie(t *testing.T){
//	user,_ := infra.NewUserPersistence().FindById(model.UserID(1012))
//	_ = user.UnsetBadMovie(model.MovieID(10))
//	fmt.Println("||||||||||||||")
//	fmt.Println(user.GoodMovies)
//	fmt.Println("||||||||||||||")
//	fmt.Println(user.BadMovies)
//	fmt.Println("||||||||||||||")
//	//_ = infra.NewUserPersistence().UpdateUser(user)
//}

