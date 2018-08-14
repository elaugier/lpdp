package main

import (
	"github.com/elaugier/lpdp/pkg/lpdp-models"
	"github.com/gin-gonic/gin"
)

func main() {
	var p models.Product
	p.Code = "test"
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
