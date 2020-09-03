package Controllers

import (
	"fmt"

	"MyPIPE/Services"
	"github.com/gin-gonic/gin"
)

type VideoFileController struct {}

func (v *VideoFileController) Store(c *gin.Context){
	file, header, _ := c.Request.FormFile("file")
	fmt.Println(header)
	Services.UploadFileToS3(file)
}

func (v *VideoFileController) Index(c *gin.Context){
	c.String(200,"OKKK.")
}