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
	if !branch.IsHead {
		logs.WarnMsg("You need to create a branch from the head, move on it first.")
		return nil
	}

	for {

		newBranchName, err := forms.AskInput("Enter the name of the new branch: ")
		if err != nil {
			return fmt.Errorf("failed to get input: %w", err)
		}

		if isBranchNameAlreadyUsed(newBranchName, branches) {
			warnMsg := fmt.Sprintf("%s is already used, please choose another name.", ui.RenderElementSelected(newBranchName))
			logs.WarnMsg(warnMsg)
			continue
		}

		newRef := plumbing.NewHashReference(plumbing.ReferenceName("refs/heads/"+newBranchName), head.Hash())
		err = repo.Storer.SetReference(newRef)
		if err != nil {
			return fmt.Errorf("failed to create new branch: %w", err)
		}

		msgSuccessfullyCreated := fmt.Sprintf("New branch %s based on %s created.", ui.RenderElementSelected(newBranchName), ui.RenderElementSelected(branch.Name))
		logs.Info(msgSuccessfullyCreated)

		msgConfirmation := fmt.Sprintf("Do you want to checkout on the new branch %s created?", ui.RenderElementSelected(newBranchName))
		checkoutOnBranchCreated, err := forms.AskConfirmation(msgConfirmation)
		if err != nil {
			return fmt.Errorf("failed to get confirmation: %w", err)
		}

		if checkoutOnBranchCreated {
			err := checkout(repo, "refs/heads/"+newBranchName)
			if err != nil {
				return fmt.Errorf("failed to checkout new branch: %w", err)
			}
			logs.Info(fmt.Sprintf("Switched to new branch %s", ui.RenderElementSelected(newBranchName)))
		}

		return nil
	}
}

func isBranchNameAlreadyUsed(newBranchName string, branches []branches.BranchWithSymbol) bool {
	for _, branch := range branches {
		if branch.Name == newBranchName {
			return true
		}
	}

	return false
}
