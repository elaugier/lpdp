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
)

func main() {

	conf, err := config.Get()
	if err != nil {
		log.Fatalf("Error on loading configuration : %v", err)
	}

	logFolder := os.ExpandEnv(conf.GetString("logFolder"))

	fmt.Println("logFolder : '" + logFolder + "'")

	timestamp := strconv.FormatInt(time.Time.UnixNano(time.Now()), 10)
	f, err := os.OpenFile(logFolder+"/"+timestamp+"_"+os.Args[0]+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error on open log file : %v", err)
	}
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)

	// set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	gin.DefaultWriter = multi

	db := db.NewInstance(
		conf.GetString("database.driver"),
		conf.GetString("database.hostname"),
		conf.GetString("database.port"),
		conf.GetString("database.dbname"),
		conf.GetString("database.user"),
		conf.GetString("database.password"),
	)
	defer db.Close()

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
