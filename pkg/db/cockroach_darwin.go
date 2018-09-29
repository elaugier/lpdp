package db

import (
	"os/exec"

	"github.com/spf13/viper"
)

//CockroachStarter
func CockroachStarter(configuration *viper.Viper) *exec.Cmd {
	/**
	 * Cockroach Initialization
	 */
	cockroachPath := configuration.GetString("cockroachPath")
	cockroachArgs := configuration.GetString("cockroachArgs")
	cockroachProc := exec.Command(cockroachPath, cockroachArgs)
	return cockroachProc
}
