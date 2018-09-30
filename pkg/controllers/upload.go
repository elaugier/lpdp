package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Upload ...
func Upload(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	// c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
