package main

import (
	br "github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
	"github.com/abroudoux/branch/internal/logs"
)

func main() {
	// branches, err := git.GetBranchesWithDefaultIndication()
	// if err != nil {
	// 	utils.PrintErrorAndExit(err)
	// }

	// branchSelected, err := menus.ChooseBranch(branches)
	// if err != nil {
	// 	utils.PrintErrorAndExit(err)
	// }

	// action, err := menus.ChooseAction(branchSelected)
	// if err != nil {
	// 	utils.PrintErrorAndExit(err)
	// }

	// err = git.DoAction(branchSelected, action)
	// if err != nil {
	// 	utils.PrintErrorAndExit(err)
	// }
	//

	repo, err := git.GetRepositoryCurrentDir()
	if err != nil {
		logs.Error("Error: ", err)
		panic(err)
	}

	branches, err := br.GetBranches(repo)
	if err != nil {
		logs.Error("Error: ", err)
		panic(err)
	}

	for _, b := range branches {
		println(b.Name())
	}

	head, err := br.GetHead(repo)
	if err != nil {
		logs.Error("Error: ", err)
		panic(err)
	}

	println(head.Name())
}
