package app

import (
	"fmt"
	"gin/config"
	"gin/service"
	"github.com/gin-gonic/gin"
)

func UploadFileTest(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		panic(err)
	}

	fmt.Println(config.GetProjectTruePath() + "/upload/" + header.Filename)

	c.SaveUploadedFile(header, config.GetProjectTruePath()+"/upload/"+header.Filename)
	service.UploadFile(header)
}
