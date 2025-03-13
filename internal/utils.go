package internal

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/log"
)

func copyBranchName(branch Branch) error {
	if clipboard.Unsupported {
		return fmt.Errorf("Clipboard not supported on this plateform.")
	}

	err := clipboard.WriteAll(branch.Name().Short())
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Branch name %s copy to the clipboard.", renderElSelected(branch.Name().Short())))
	return nil
}

func cleanView() {
	fmt.Print("\033[H\033[2J")
}

func PrintErrorExitProgram(err error) {
	log.Error("Error: ", err)
	os.Exit(1)
}
