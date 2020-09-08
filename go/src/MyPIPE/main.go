package main

import (
	"MyPIPE/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func init() {
	const location = "Asia/Tokyo"
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main(){
	router := gin.Default()

	video_file_controller := initVideoFileController()
	router.POST("/new",handler.TemporaryRegisterUser)

	router.POST("/", video_file_controller.Store)
	
	router.GET("/test", func(c *gin.Context){
		c.String(200,"fesfes.")
	})

	router.GET("/checking",video_file_controller.Index)

    router.Run()
}