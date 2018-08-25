package main

import (
	"bufio"
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
	g      errgroup.Group
	logger *log.Logger
)

func main() {

	/**
	 * Logger Initialization
	 */
	logger = log.New(os.Stdout, "", log.LstdFlags)

	fullBinaryName, err := osext.Executable()
	if err != nil {
		logger.Fatal(err)
	}

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		logger.Fatal(err)
	}

	binaryName := strings.Replace(strings.Replace(fullBinaryName, folderPath, "", -1), string(os.PathSeparator), "", -1)

	logger.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	logger.SetPrefix(binaryName + " " + strconv.Itoa(os.Getpid()) + " ")

	config, err := config.Get()
	if err != nil {
		logger.Fatal(err)
	}

	timestampStart := strconv.FormatInt(time.Time.UnixNano(time.Now()), 10)
	logFile := os.ExpandEnv(config.GetString("logFolder")) + "/" + timestampStart + "_" + binaryName + ".log"
	logger.Println("log file location => '" + logFile + "'")
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Fatal(err)
	}
	multi := io.MultiWriter(f, os.Stdout)
	logger.SetOutput(multi)

	// set log format
	logger.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	gin.DefaultWriter = multi
	gin.DefaultErrorWriter = multi
	/**
	 * Cockroach Initialization
	 */
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
		logger.Fatalf("Error on get stdOut of cockroach process")
	}
	stderr, err := cockroachProc.StderrPipe()
	if err != nil {
		logger.Fatalf("Error on get stdErr of cockroach process")
	}

	go Scan(stdout)
	go Scan(stderr)

	g.Go(func() error {
		err = cockroachProc.Start()
		if err != nil {
			logger.Fatalf("Error on starting cockroach : %s %s => %v", cockroachPath, cockroachArgs, err)
		}
		bufStdout := new(bytes.Buffer)
		bufStderr := new(bytes.Buffer)
		bufStdout.ReadFrom(stdout)
		bufStderr.ReadFrom(stderr)
		if err = cockroachProc.Wait(); err != nil {
			logger.Fatalf("Error when execute %s %s => %v\r\nstdout = '%s'\r\nstderr = '%s'",
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

	backendServerAddr := fmt.Sprintf(":%s", backendPort)
	logger.Printf("backendServerAddr = '%s'", backendServerAddr)

	backendServer := &http.Server{
		Addr:         backendServerAddr,
		Handler:      server.BackendRouter(logger),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     logger,
	}

	frontendServerAddr := fmt.Sprintf(":%s", frontendPort)
	logger.Printf("frontendServerAddr = '%s'", frontendServerAddr)

	frontendServer := &http.Server{
		Addr:         frontendServerAddr,
		Handler:      server.FrontendRouter(logger),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     logger,
	}

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
		logger.Fatal(err)
	}

}

//Scan ...
func Scan(i io.ReadCloser) {
	// read command's stdout line by line
	in := bufio.NewScanner(i)

	for in.Scan() {
		logger.Printf(in.Text()) // write each line to your log, or anything you need
	}
	if err := in.Err(); err != nil {
		logger.Printf("error: %s", err)
	}
}
