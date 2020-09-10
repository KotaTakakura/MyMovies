package test

import (
	"MyPIPE/domain/model"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUserPassword(t *testing.T){
	hashedPassword := model.NewUserPassword("testhashedpassword")
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte("testhashedpassword"))
	if err != nil {
		t.Fatal("ハッシュ化パスワードが一致しません。")
	}
}