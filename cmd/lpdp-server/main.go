package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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

	cockroachPath := config.GetString("cockroachPath")
	cockroachArgs := config.GetString("cockroachArgs")

	cockroachProc := exec.Command(cockroachPath)

	cockroachProc.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CmdLine:       fmt.Sprintf("%s", cockroachArgs),
		CreationFlags: 0,
	}

	stdout, err := cockroachProc.StdoutPipe()
	if err != nil {
		log.Fatalf("Error on get stdOut of cockroach process")
	}
	stderr, err := cockroachProc.StderrPipe()
	if err != nil {
		log.Fatalf("Error on get stdErr of cockroach process")
	}

	g.Go(func() error {
		err = cockroachProc.Start()
		if err != nil {
			log.Fatalf("Error on starting cockroach : %s %s => %v", cockroachPath, cockroachArgs, err)
		}
		bufStdout := new(bytes.Buffer)
		bufStderr := new(bytes.Buffer)
		bufStdout.ReadFrom(stdout)
		bufStderr.ReadFrom(stderr)
		if err = cockroachProc.Wait(); err != nil {
			log.Fatalf("Error when execute %s %s => %v\r\nstdout = '%s'\r\nstderr = '%s'",
				cockroachPath, cockroachArgs, err, bufStdout.String(), bufStderr.String())
		}
		return err
	})

	backendHostname := config.GetString("backendHostname")
	backendPort := config.GetString("backendPort")
	frontendHostname := config.GetString("frontendHostname")
	frontendPort := config.GetString("frontendPort")

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

	// New functionality written in Go
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "New function")
	})

	urlBackend, _ := url.Parse(fmt.Sprintf("http://localhost:%s/", backendPort))
	http.Handle(fmt.Sprintf("%s/", backendHostname), httputil.NewSingleHostReverseProxy(urlBackend))

	urlFrontend, _ := url.Parse(fmt.Sprintf("http://localhost:%s/", frontendPort))
	http.Handle(fmt.Sprintf("%s/", frontendHostname), httputil.NewSingleHostReverseProxy(urlFrontend))

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
