package test

import (
	"MyPIPE/domain/model"
	"MyPIPE/infra"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
	"time"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestUserRepository(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	userRepository := infra.NewUserPersistence()
	userRepository.DatabaseAccessor = db

	user := model.User{}
	user.ID = 999
	user.Name = "John Doe"
	user.Password = "TestPassword"
	user.Email = "Test@example.com"
	user.Token = "TestToken"
	user.Birthday = time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)

	mock.ExpectBegin()
	// Mock設定
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `users` "+
			"(`id`,`name`,`password`,`email`,`birthday`,`token`,`created_at`,`updated_at`) "+
			"VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs(user.ID, user.Name, user.Password, user.Email, user.Birthday, user.Token, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userRepository.SetUser(&user)

	// 使用されたモックDBが期待通りの値を持っているかを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
