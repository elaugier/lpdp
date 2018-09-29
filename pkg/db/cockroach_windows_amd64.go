package db

import (
	"fmt"
	"os/exec"
	"syscall"

	"github.com/spf13/viper"
)

//CockroachStarter ...
func CockroachStarter(configuration *viper.Viper) *exec.Cmd {
	/**
	 * Cockroach Initialization
	 */
	cockroachPath := configuration.GetString("cockroachPath")
	cockroachArgs := configuration.GetString("cockroachArgs")
	cockroachProc := exec.Command(cockroachPath, cockroachArgs)
	cockroachProc.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CmdLine:       fmt.Sprintf("%s", cockroachArgs),
		CreationFlags: 0,
	}
	return cockroachProc
}
