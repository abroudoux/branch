package main

import (
	_ "embed"
	"fmt"
	"os"

	git "github.com/abroudoux/branch/internal/git"
	menus "github.com/abroudoux/branch/internal/menus"
	repository "github.com/abroudoux/branch/internal/repository"
	utils "github.com/abroudoux/branch/internal/utils"
)

func main() {
	branches := git.GetBranchesWithDefaultIndication()

	if len(os.Args) > 1 {
		err := repository.FlagMode(branches)
		if err != nil {
			utils.PrintErrorAndExit(err)
		}
		os.Exit(0)
	}

	err := utils.IsInGitRepository()
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	branch, err := menus.ChooseBranch(branches)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	action, err := menus.ChooseAction(branch)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	err = doAction(branch, action)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}
}

func doAction(branch string, action string) error {
	switch action {
	case "Exit":
		fmt.Println("Exiting...")
		return nil
	case "Delete":
		return git.DeleteBranch(branch)
	case "Merge":
		return git.MergeBranch(branch)
	case "Branch":
		return git.CreateBranch(branch)
	case "Rename":
		return git.RenameBranch(branch)
	case "Checkout":
		return git.CheckoutBranch(branch)
	case "Name":
		return git.CopyName(branch)
	default:
		return fmt.Errorf("invalid action: %s", action)
	}
}