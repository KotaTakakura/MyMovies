package infra

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// データベース
	Dialect = "mysql"

	// ユーザー名
	DBUser = "root"

	// パスワード
	DBPass = "root"

	// プロトコル
	DBProtocol = "tcp(MyPIPE-mysql:3306)"

	// DB名
	DBName = "mypipe"
)

func ConnectGorm() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s?parseTime=true&loc=Asia%%2FTokyo"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
