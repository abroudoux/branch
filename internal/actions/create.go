package actions

import (
	"fmt"

	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/forms"
	"github.com/abroudoux/branch/internal/git"
	"github.com/abroudoux/branch/internal/logs"
	"github.com/abroudoux/branch/internal/ui"
	"github.com/go-git/go-git/v5/plumbing"
)

func createNewBranch(repo git.Repository, branch branches.BranchWithSymbol, branches []branches.BranchWithSymbol, head branches.Branch) error {
	newBranchName, err := forms.AskInput("Enter the name of the new branch: ")
	if err != nil {
		return err
	}

	if isBranchNameAlreadyUsed(newBranchName, branches) {
		warnMsg := fmt.Sprintf("%s is already used, please choose another name.", ui.RenderElementSelected(newBranchName))
		logs.WarnMsg(warnMsg)
		createNewBranch(repo, branch, branches, head)
	}

	if !branch.IsHead {
		err := checkout(repo, branch.Branch.Name().String())
		if err != nil {
			return nil
		}
	}

	newRef := plumbing.NewHashReference(plumbing.ReferenceName("refs/heads/"+newBranchName), head.Hash())
	err = repo.Storer.SetReference(newRef)
	if err != nil {
		return err
	}

	msgSuccessfullyCreated := fmt.Sprintf("New branch %s based on %s created.", ui.RenderElementSelected(newBranchName), ui.RenderElementSelected(branch.Name))
	logs.Info(msgSuccessfullyCreated)

	msgConfirmation := fmt.Sprintf("Do you want to checkout on the new branch %s created?", ui.RenderElementSelected(newBranchName))
	checkoutOnBranchCreated, err := forms.AskConfirmation(msgConfirmation)
	if err != nil {
		return err
	}

	if checkoutOnBranchCreated {
		err := checkout(repo, "refs/heads/"+newBranchName)
		if err != nil {
			return nil
		}
	}

	return nil
}

func isBranchNameAlreadyUsed(newBranchName string, branches []branches.BranchWithSymbol) bool {
	for _, branch := range branches {
		if branch.Name == newBranchName {
			return true
		}
	}

	return false
}
