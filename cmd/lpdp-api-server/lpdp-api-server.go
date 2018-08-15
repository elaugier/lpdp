package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/elaugier/lpdp/pkg/db"
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

	timestamp := strconv.FormatInt(time.Time.UnixNano(time.Now()), 10)
	f, err := os.OpenFile(os.Getenv("TEMP")+"/"+timestamp+"_"+os.Args[0]+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)

	// set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	gin.DefaultWriter = multi

	conn, err := gorm.Open(DATABASEDRIVER, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		DATABASEHOSTNAME, DATABASEPORT, DATABASEUSERNAME, DATABASENAME, DATABASEPASSWORD))
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
		return
	}
	log.Println("database connected")
	defer conn.Close()

	db.DatabaseInitialization(conn)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		auth := "false"
		claims := jwt.ExtractClaims(c)
		if claims != nil {
			auth = "true"
		}
		c.JSON(200, gin.H{
			"message":        "pong",
			"VerifyAudience": auth,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
