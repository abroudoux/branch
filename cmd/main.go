package main

import (
	_ "embed"
	"os"

	git "github.com/abroudoux/branch/internal/git"
	menus "github.com/abroudoux/branch/internal/menus"
	repository "github.com/abroudoux/branch/internal/repository"
	utils "github.com/abroudoux/branch/internal/utils"
)

func main() {
	if len(os.Args) > 1 {
		err := repository.FlagMode()
		if err != nil {
			utils.PrintErrorAndExit(err)
		}
		os.Exit(0)
	}

	err := utils.IsInGitRepository()
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	branches, err := git.GetBranchesWithDefaultIndication()
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	branchSelected, err := menus.ChooseBranch(branches)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	action, err := menus.ChooseAction(branchSelected)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	err = git.DoAction(branchSelected, action)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}
}