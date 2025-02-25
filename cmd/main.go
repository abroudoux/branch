package main

import (
	"os"

	"github.com/abroudoux/branch/internal/actions"
	br "github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
	"github.com/abroudoux/branch/internal/logs"
	"github.com/abroudoux/branch/internal/utils"
	"github.com/charmbracelet/log"
)

func main() {
	if len(os.Args) > 1 {
		option := os.Args[1]

		switch option {
		case "--help", "-h", "help":
			utils.HelpManual()
		case "--version", "-v", "version":
			utils.Version()
		default:
			logs.WarnMsg("Unknown option.")
			utils.HelpManual()
		}

		os.Exit(0)
	}

	repo, err := git.GetRepositoryCurrentDir()
	if err != nil {
		log.Warn("You are not in a git repository.")
		os.Exit(0)
	}

	branches, err := br.GetBranches(repo)
	if err != nil {
		utils.PrintErrorExitProgram(err)
	}

	head, err := br.GetHead(repo)
	if err != nil {
		utils.PrintErrorExitProgram(err)
	}

	branchesWithDetails := br.CreateBranchesWithDetails(branches, head)
	branchSelected, err := br.SelectBranch(branchesWithDetails)
	if err != nil {
		utils.PrintErrorExitProgram(err)
	}

	if branchSelected.Name == "" {
		logs.Info("Program exited...")
		os.Exit(0)
	}

	actionSelected, err := actions.SelectAction(branchSelected)
	if err != nil {
		utils.PrintErrorExitProgram(err)
	}

	err = actions.DoBranchAction(repo, branchSelected, branchesWithDetails, head, actionSelected)
	if err != nil {
		utils.PrintErrorExitProgram(err)
	}
}
