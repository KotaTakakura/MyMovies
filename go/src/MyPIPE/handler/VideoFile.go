package handler

import (
	"fmt"

	Iservice "MyPIPE/Interfaces/Services"
	"github.com/gin-gonic/gin"
)

type VideoFileController struct {
	UploadToS3Service Iservice.IUploadToS3
}

func (v *VideoFileController) Store(c *gin.Context) {
	file, header, _ := c.Request.FormFile("file")
	fmt.Println(header)
	v.UploadToS3Service.Upload(file)
}

func (v *VideoFileController) Index(c *gin.Context) {
	c.String(200, "OKKK.")
}

func NewVideoFileController(i Iservice.IUploadToS3) *VideoFileController {
	return &VideoFileController{
		UploadToS3Service: i,
	}
}
