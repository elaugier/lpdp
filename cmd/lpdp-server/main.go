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
	"os/signal"
	"runtime"
	"syscall"
	"time"

	// jwt "github.com/appleboy/gin-jwt"
	"github.com/elaugier/lpdp/pkg/config"
	"github.com/elaugier/lpdp/pkg/db"
	"github.com/elaugier/lpdp/pkg/logs"
	"github.com/elaugier/lpdp/pkg/server"
	"golang.org/x/sync/errgroup"
)

var (
	g      errgroup.Group
	run    bool
	logger *log.Logger
	//Db ...
	Db db.Instance
)

func main() {

	run = true
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	/**
	 * Logger Initialization
	 */
	logger = logs.GetInstance()

	configuration, err := config.Get()
	if err != nil {
		logger.Fatal(err)
	}

	//Log configuration
	logger.Printf("cockroach.win.cockroachPath : '%s'", configuration.GetString("cockroach.win.cockroachPath"))
	logger.Printf("cockroach.win.cockroachArgs : '%s'", configuration.GetString("cockroach.win.cockroachArgs"))
	logger.Printf("cockroach.lnx.cockroachPath : '%s'", configuration.GetString("cockroach.lnx.cockroachPath"))
	logger.Printf("cockroach.lnx.cockroachArgs : '%s'", configuration.GetString("cockroach.lnx.cockroachArgs"))
	logger.Printf("cockroach.mac.cockroachPath : '%s'", configuration.GetString("cockroach.mac.cockroachPath"))
	logger.Printf("cockroach.mac.cockroachArgs : '%s'", configuration.GetString("cockroach.mac.cockroachArgs"))

	cockroachProc := db.CockroachStarter(configuration)

	go func() {
		<-c
		logger.Println("intercept interruption : SIGTERM")
		if runtime.GOOS == "windows" {
			logger.Println("Windows doesn't support Interrupt : send Kill")
			if err := cockroachProc.Process.Signal(os.Kill); err != nil {
				log.Println("failed to kill process: ", err)
			}
		} else {
			if err := cockroachProc.Process.Signal(syscall.SIGTERM); err != nil {
				log.Println("failed to send SIGTERM: ", err)
			}
		}
		run = false
	}()

	stdout, err := cockroachProc.StdoutPipe()
	if err != nil {
		logger.Fatalf("Error on get stdOut of cockroach process")
	}
	stderr, err := cockroachProc.StderrPipe()
	if err != nil {
		logger.Fatalf("Error on get stdErr of cockroach process")
	}

	go Scan(logger, stdout)
	go Scan(logger, stderr)

	g.Go(func() error {
		err = cockroachProc.Start()
		if err != nil {
			logger.Fatalf("Error on starting cockroach : %v", err)
		}
		bufStdout := new(bytes.Buffer)
		bufStderr := new(bytes.Buffer)
		bufStdout.ReadFrom(stdout)
		bufStderr.ReadFrom(stderr)
		if err = cockroachProc.Wait(); err != nil {
			logger.Fatalf("Error when execute cockroach => %v\r\nstdout = '%s'\r\nstderr = '%s'",
				err, bufStdout.String(), bufStderr.String())
		}
		return err
	})

	backendHostname := configuration.GetString("backendHostname")
	backendPort := configuration.GetString("backendPort")
	frontendHostname := configuration.GetString("frontendHostname")
	frontendPort := configuration.GetString("frontendPort")

	database := db.GetInstance()
	defer database.Close()

	backendServerAddr := fmt.Sprintf(":%s", backendPort)
	logger.Printf("backendServerAddr = '%s'", backendServerAddr)

	backendServer := &http.Server{
		Addr:           backendServerAddr,
		Handler:        server.BackendRouter(logger),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		ErrorLog:       logger,
		MaxHeaderBytes: 1 << 20,
	}

	frontendServerAddr := fmt.Sprintf(":%s", frontendPort)
	logger.Printf("frontendServerAddr = '%s'", frontendServerAddr)

	frontendServer := &http.Server{
		Addr:           frontendServerAddr,
		Handler:        server.FrontendRouter(logger),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		ErrorLog:       logger,
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
		logger.Fatal(err)
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
		logger.Printf("%30s | %30s | %5s | %s\n", r.Host, r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
