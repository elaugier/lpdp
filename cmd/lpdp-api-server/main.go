package main

import (
	"fmt"
	"os"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//DATABASEDRIVER ...
var DATABASEDRIVER = os.Getenv("DATABASEDRIVER")

//DATABASEHOSTNAME ...
var DATABASEHOSTNAME = os.Getenv("DATABASEHOSTNAME")

//DATABASEPORT ...
var DATABASEPORT = os.Getenv("DATABASEPORT")

//DATABASENAME ...
var DATABASENAME = os.Getenv("DATABASENAME")

//DATABASEUSERNAME ...
var DATABASEUSERNAME = os.Getenv("DATABASEUSERNAME")

//DATABASEPASSWORD ...
var DATABASEPASSWORD = os.Getenv("DATABASEPASSWORD")

func main() {
	conn, err := gorm.Open(DATABASEDRIVER, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		DATABASEHOSTNAME, DATABASEPORT, DATABASEUSERNAME, DATABASENAME, DATABASEPASSWORD))
	if err != nil {
		fmt.Printf("couldn't connect to database: %v\n", err)
		return
	}
	defer conn.Close()

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
