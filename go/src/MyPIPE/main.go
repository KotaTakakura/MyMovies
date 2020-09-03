package main

import(
	"MyPIPE/Controllers"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	video_file_controller := Controllers.VideoFileController{}

	router.POST("/", video_file_controller.Store)
	
	router.GET("/test", func(c *gin.Context){
		c.String(200,"OKdddfafaesfeas.")
	})

	router.GET("/checking",video_file_controller.Index)

    router.Run()
}