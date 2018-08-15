package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

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

var (
	//Trace ...
	Trace *log.Logger
	//Info ...
	Info *log.Logger
	//Warning ...
	Warning *log.Logger
	//Error ...
	Error *log.Logger
)

//InitLog ...
func InitLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	InitLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

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
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
