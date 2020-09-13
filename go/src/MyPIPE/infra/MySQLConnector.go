package infra

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func ConnectGorm() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s?parseTime=true&loc=Asia%%2FTokyo"
	connect := fmt.Sprintf(connectTemplate, os.Getenv("DBUser"), os.Getenv("DBPass"), os.Getenv("DBProtocol"), os.Getenv("DBName"))
	db, err := gorm.Open(os.Getenv("Dialect"), connect)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
