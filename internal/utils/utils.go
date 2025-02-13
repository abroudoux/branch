package utils

import (
	"os"

	"github.com/abroudoux/branch/internal/logs"
)

func PrintErrorExitProgram(err error) {
	logs.Error("Error: ", err)
	os.Exit(1)
}
