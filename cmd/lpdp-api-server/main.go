package main

import (
	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {
	var p models.User
	p.ID = "test"
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
