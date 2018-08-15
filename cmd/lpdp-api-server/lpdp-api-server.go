package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/elaugier/lpdp/pkg/config"
	"github.com/elaugier/lpdp/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	conf, err := config.Get()

	logFolder := os.ExpandEnv(conf.GetString("logFolder"))

	fmt.Println("logFolder : '" + logFolder + "'")

	timestamp := strconv.FormatInt(time.Time.UnixNano(time.Now()), 10)
	f, err := os.OpenFile(logFolder+"/"+timestamp+"_"+os.Args[0]+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)

	// set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	gin.DefaultWriter = multi

	databaseDriver := conf.GetString("database.driver")
	log.Printf("database driver selected: %s", databaseDriver)
	databaseHostname := conf.GetString("database.hostname")
	log.Printf("database hostname selected: %s", databaseHostname)
	databasePort := conf.GetString("database.port")
	log.Printf("database port selected: %s", databasePort)
	databaseDbName := conf.GetString("database.dbname")
	log.Printf("database dbname selected: %s", databaseDbName)
	databaseUsername := conf.GetString("database.user")
	log.Printf("database username selected: %s", databaseUsername)
	databasePassword := conf.GetString("database.password")
	log.Printf("database password found!")

	conn, err := gorm.Open(databaseDriver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		databaseHostname, databasePort, databaseUsername, databaseDbName, databasePassword))
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
