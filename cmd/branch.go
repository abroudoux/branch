package main

import (
	"os"

	"github.com/abroudoux/branch/internal"
	"github.com/charmbracelet/log"
)

func main() {
	internal.HandleFlags()

	repository, err := internal.GetRepository(".")
	if err != nil {
		log.Warn("You're not in a git repository.")
		os.Exit(0)
	}

	branchSelected, err := repository.SelectBranch()
	if err != nil {
		internal.PrintErrorExitProgram(err)
	}

	if branchSelected == nil {
		log.Info("Program exited..")
		os.Exit(0)
	}

	action, err := internal.SelectAction(branchSelected)
	if err != nil {
		internal.PrintErrorExitProgram(err)
	}

	err = repository.DoBranchAction(branchSelected, action)
	if err != nil {
		internal.PrintErrorExitProgram(err)
	}
}
