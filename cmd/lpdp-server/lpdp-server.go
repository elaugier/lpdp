package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kardianos/osext"

	// jwt "github.com/appleboy/gin-jwt"
	"github.com/elaugier/lpdp/pkg/config"
	"github.com/elaugier/lpdp/pkg/server"
	//"github.com/elaugier/lpdp/pkg/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {

	fullBinaryName, err := osext.Executable()
	if err != nil {
		log.Fatal(err)
	}

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}

	binaryName := strings.Replace(strings.Replace(fullBinaryName, folderPath, "", -1), string(os.PathSeparator), "", -1)

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	log.SetPrefix(binaryName + " " + strconv.Itoa(os.Getpid()) + " ")

	config, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	timestampStart := strconv.FormatInt(time.Time.UnixNano(time.Now()), 10)
	logFile := os.ExpandEnv(config.GetString("logFolder")) + "/" + timestampStart + "_" + binaryName + ".log"
	log.Println("log file location => '" + logFile + "'")
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)

	// set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	gin.DefaultWriter = multi

	// db := db.NewInstance(
	// 	conf.GetString("database.driver"),
	// 	conf.GetString("database.hostname"),
	// 	conf.GetString("database.port"),
	// 	conf.GetString("database.dbname"),
	// 	conf.GetString("database.user"),
	// 	conf.GetString("database.password"),
	// )
	// defer db.Close()

	backendServer := &http.Server{
		Addr:         ":8080",
		Handler:      server.BackendRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		//ErrorLog:     *log,
	}

	frontendServer := &http.Server{
		Addr:         ":8081",
		Handler:      server.FrontendRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		//ErrorLog:     *log,
	}

	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "New function")
	})

	u1, _ := url.Parse("http://localhost:8080/")
	http.Handle("lpdpbackend.local/", httputil.NewSingleHostReverseProxy(u1))

	u2, _ := url.Parse("http://localhost:8081/")
	http.Handle("lpdpfrontend.local/", httputil.NewSingleHostReverseProxy(u2))

	g.Go(func() error {
		return backendServer.ListenAndServe()
	})

	g.Go(func() error {
		return frontendServer.ListenAndServe()
	})

	g.Go(func() error {
		return http.ListenAndServe(":80", nil)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	/* 	r := gin.Default()
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

	   	r.Run() // listen and serve on 0.0.0.0:8080 */
}
