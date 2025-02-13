package main

import (
	"os"

	"github.com/abroudoux/branch/internal/actions"
	br "github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
	"github.com/abroudoux/branch/internal/logs"
)

func main() {
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

	head, err := br.GetHead(repo)
	if err != nil {
		logs.Error("Error: ", err)
		panic(err)
	}

	branchesWithDetails := br.CreateBranchesWithDetails(branches, head)
	branchSelected, err := br.SelectBranch(branchesWithDetails)
	if err != nil {
		logs.Error("Error: ", err)
		os.Exit(1)
	}

	if branchSelected.Name == "" {
		logs.Info("Program exited...")
		os.Exit(0)
	}

	actionSelected, err := actions.SelectAction(branchSelected)
	if err != nil {
		logs.Error("Error: ", err)
		panic(err)
	}

	err = actions.DoBranchAction(repo, branchSelected, branchesWithDetails, head, actionSelected)
	if err != nil {
		logs.Error("Error: ", err)
		panic(err)
	}
}
