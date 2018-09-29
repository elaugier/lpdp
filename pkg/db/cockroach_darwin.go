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
	cockroachPath := configuration.GetString("cockroach.mac.cockroachPath")
	cockroachArgs := configuration.GetString("cockroach.mac.cockroachArgs")
	cockroachProc := exec.Command(cockroachPath, cockroachArgs)
	return cockroachProc
}
