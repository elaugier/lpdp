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
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/kardianos/osext"

	// jwt "github.com/appleboy/gin-jwt"
	"github.com/elaugier/lpdp/pkg/config"
	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g      errgroup.Group
	run    bool
	logger *logs.Instance
	//Db ...
	Db db.Instance
)

func main() {

	run = true
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		run = false
	}()

	/**
	 * Logger Initialization
	 */
	logger = logs.GetInstance()

	fullBinaryName, err := osext.Executable()
	if err != nil {
		logger.L.Fatal(err)
	}

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		logger.L.Fatal(err)
	}

	binaryName := strings.Replace(strings.Replace(fullBinaryName, folderPath, "", -1), string(os.PathSeparator), "", -1)

	logger.L.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	logger.L.SetPrefix(binaryName + " " + strconv.Itoa(os.Getpid()) + " ")

	configuration, err := config.Get()
	if err != nil {
		logger.L.Fatal(err)
	}

	timestampStart := strconv.FormatInt(time.Time.UnixNano(time.Now()), 10)
	logFile := os.ExpandEnv(configuration.GetString("logFolder")) + "/" + timestampStart + "_" + binaryName + ".log"
	logger.L.Println("log file location => '" + logFile + "'")
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.L.Fatal(err)
	}
	multi := io.MultiWriter(f, os.Stdout)
	logger.L.SetOutput(multi)

	// set log format
	logger.L.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	gin.DefaultWriter = multi
	gin.DefaultErrorWriter = multi
	gin.DisableConsoleColor()
	/**
	 * Cockroach Initialization
	 */
	cockroachPath := configuration.GetString("cockroachPath")
	cockroachArgs := configuration.GetString("cockroachArgs")
	cockroachProc := exec.Command(cockroachPath)
	cockroachProc.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CmdLine:       fmt.Sprintf("%s", cockroachArgs),
		CreationFlags: 0,
	}

	stdout, err := cockroachProc.StdoutPipe()
	if err != nil {
		logger.L.Fatalf("Error on get stdOut of cockroach process")
	}
	stderr, err := cockroachProc.StderrPipe()
	if err != nil {
		logger.L.Fatalf("Error on get stdErr of cockroach process")
	}

	go Scan(logger.L, stdout)
	go Scan(logger.L, stderr)

	g.Go(func() error {
		err = cockroachProc.Start()
		if err != nil {
			logger.L.Fatalf("Error on starting cockroach : %s %s => %v", cockroachPath, cockroachArgs, err)
		}
		bufStdout := new(bytes.Buffer)
		bufStderr := new(bytes.Buffer)
		bufStdout.ReadFrom(stdout)
		bufStderr.ReadFrom(stderr)
		if err = cockroachProc.Wait(); err != nil {
			logger.L.Fatalf("Error when execute %s %s => %v\r\nstdout = '%s'\r\nstderr = '%s'",
				cockroachPath, cockroachArgs, err, bufStdout.String(), bufStderr.String())
		}
		return err
	})

	backendHostname := configuration.GetString("backendHostname")
	backendPort := configuration.GetString("backendPort")
	frontendHostname := configuration.GetString("frontendHostname")
	frontendPort := configuration.GetString("frontendPort")

	database := db.GetInstance(true, multi)
	defer database.Close()
	//database.Connection.SetLogger(logger)
	database.DatabaseInitialization()

	backendServerAddr := fmt.Sprintf(":%s", backendPort)
	logger.L.Printf("backendServerAddr = '%s'", backendServerAddr)

	backendServer := &http.Server{
		Addr:           backendServerAddr,
		Handler:        server.BackendRouter(logger.L),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		ErrorLog:       logger.L,
		MaxHeaderBytes: 1 << 20,
	}

	frontendServerAddr := fmt.Sprintf(":%s", frontendPort)
	logger.L.Printf("frontendServerAddr = '%s'", frontendServerAddr)

	frontendServer := &http.Server{
		Addr:           frontendServerAddr,
		Handler:        server.FrontendRouter(logger.L),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		ErrorLog:       logger.L,
		MaxHeaderBytes: 1 << 20,
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
		return http.ListenAndServe(":80", logRequest(http.DefaultServeMux))
	})

	if err := g.Wait(); err != nil {
		logger.L.Fatal(err)
	}

}

//Scan ...
func Scan(logger *log.Logger, i io.ReadCloser) {
	// read command's stdout line by line
	in := bufio.NewScanner(i)

	for in.Scan() {
		logger.Printf(in.Text()) // write each line to your log, or anything you need
	}
	if err := in.Err(); err != nil {
		logger.Printf("error: %s", err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.L.Printf("%30s | %30s | %5s | %s\n", r.Host, r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
