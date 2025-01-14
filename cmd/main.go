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

	branches := git.GetBranchesWithDefaultIndication()

	branch, err := menus.ChooseBranch(branches)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	action, err := menus.ChooseAction(branch)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	err = git.DoAction(branch, action)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}
}