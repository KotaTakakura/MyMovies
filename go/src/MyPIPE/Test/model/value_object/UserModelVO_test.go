package test

import (
	"MyPIPE/domain/model"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUserID(t *testing.T){
	trueCases := []struct{
		ID uint64
	}{
		{ID: 100},
		{ID: 1},
	}

	falseCases := []struct{
		ID uint64
	}{
		{ID: 0},
	}

	for _,trueCase := range trueCases{
		_,userErr := model.NewUserID(trueCase.ID)
		if userErr != nil{
			t.Error("UserID TrueCase Error.")
		}
	}

	for _,falseCase := range falseCases{
		_,userErr := model.NewUserID(falseCase.ID)
		if userErr == nil{
			t.Error("UserID FalseCase Error.")
		}
	}

}

func TestUserName(t *testing.T){
	trueCases := []struct{
		Name	string
	}{
		{Name:	"田中太郎"},
		{Name:	"John Doe"},
	}

	falseCases := []struct{
		Name	string
	}{
		{Name:	""}, //空
		//長すぎ
		{
			Name:
			"ああああああああああああああああああああああああああああああああああああああああああああああああああ" +
			"ああああああああああああああああああああああああああああああああああああああああああああああああああ" +
			"ああああああああああああああああああああああああああああああああああああああああああああああああああ" +
			"ああああああああああああああああああああああああああああああああああああああああああああああああああ" +
			"ああああああああああああああああああああああああああああああああああああああああああああああああああ" +
			"ああああああああああああああああああああああああああああああああああああああああああああああああああ" +
			"ああああああああああああああああああああああああああああああ",
		},
	}

	for _, c := range trueCases {
		_,err := model.NewUserName(c.Name)
		if err != nil{
			t.Fatal("ユーザー名生成のテスト失敗。(正常系)")
		}
	}

	for _, c := range falseCases {
		_,err := model.NewUserName(c.Name)
		if err == nil{
			t.Fatal("ユーザー名生成のテスト失敗。(異常系)")
		}
	}
}

func TestUserPassword(t *testing.T){

	trueCases := []struct{
		Password	string
	}{
		{Password:	"testhashedpassword"},
	}

	falseCases := []struct{
		Password	string
	}{
		//空
		{Password:	""},
		//短すぎ
		{Password:	"abc"},
		//長すぎ
		{Password:	"ああああああああああああ"},

	}

	for _, c := range trueCases {
		hashedPassword,_ := model.NewUserPassword(c.Password)
		hashErr := bcrypt.CompareHashAndPassword(
			[]byte(hashedPassword),
			[]byte(c.Password))
		if hashErr != nil {
			t.Fatal("ハッシュ化パスワードが一致しません。")
		}
	}

	for _, c := range falseCases {
		_,err := model.NewUserPassword(c.Password)
		if err == nil {
			t.Fatal("バリデーションに異常があります。")
		}

	}
}

func TestUserEmail(t *testing.T){

	trueCases := []struct{
		Email	string
	}{
		{Email:	"taka@example.com"},
	}

	falseCases := []struct{
		Email	string
	}{
		//空
		{Email:	""},
		//無効なメアド
		{Email:	"tatata"},

	}

	for _, c := range trueCases {
		_,err := model.NewUserEmail(c.Email)
		if err != nil {
			t.Fatal("メールアドレス生成に失敗。(正常系)")
		}
	}

	for _, c := range falseCases {
		_,err := model.NewUserEmail(c.Email)
		if err == nil {
			t.Fatal("メールアドレス生成に失敗。(異常系)。")
		}

	}
}

func TestUserToken(t *testing.T){

	trueCases := []struct{
		Token	string
	}{
		{Token:	"feifaop-fesfa-eiapsf-faesi"},
	}

	falseCases := []struct{
		Token	string
	}{
		//空
		{Token:	""},
	}

	for _, c := range trueCases {
		_,err := model.NewUserToken(c.Token)
		if err != nil {
			t.Fatal("トークン生成に失敗。(正常系)")
		}
	}

	for _, c := range falseCases {
		_,err := model.NewUserToken(c.Token)
		if err == nil {
			t.Fatal("トークン生成に失敗。(異常系)。")
		}

	}
}