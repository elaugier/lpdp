package logs

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elaugier/lpdp/pkg/config"
	"github.com/gin-gonic/gin"

	"github.com/kardianos/osext"
)

//Instance ...
type Instance struct {
	//L ...
	L *log.Logger
}

//Inst ...
var Inst Instance

//NewInstance ...
func NewInstance() *Instance {
	l := log.New(os.Stdout, "", log.LstdFlags)
	fullBinaryName, err := osext.Executable()
	if err != nil {
		l.Fatal(err)
	}

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		l.Fatal(err)
	}

	binaryName := strings.Replace(strings.Replace(fullBinaryName, folderPath, "", -1), string(os.PathSeparator), "", -1)

	configuration, err := config.Get()
	if err != nil {
		l.Fatal(err)
	}

	timestampStart := strconv.FormatInt(time.Time.UnixNano(time.Now()), 10)
	logFile := os.ExpandEnv(configuration.GetString("logFolder")) + "/" + timestampStart + "_" + binaryName + ".log"
	l.Println("log file location => '" + logFile + "'")
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		l.Fatal(err)
	}

	l.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	l.SetPrefix(binaryName + " " + strconv.Itoa(os.Getpid()) + " ")
	multi := io.MultiWriter(f, os.Stdout)
	l.SetOutput(multi)

	// set log format
	l.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	gin.DefaultWriter = multi
	gin.DefaultErrorWriter = multi
	gin.DisableConsoleColor()
	Inst.L = l
	return &Instance{
		L: l,
	}
}

//GetInstance ...
func GetInstance() *log.Logger {
	if Inst == (Instance{}) {
		return NewInstance().L
	}
	return Inst.L
}
