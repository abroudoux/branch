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

func createNewBranch(repo git.Repository, branch branches.BranchWithSymbol, head branches.Branch) error {
	newBranchName, err := forms.AskInput("Enter the name of the new branch: ")
	if err != nil {
		return err
	}

	if !branch.IsHead {
		return fmt.Errorf("cannot create a new branch from a non-HEAD branch")
	}

	ref := plumbing.NewHashReference(plumbing.ReferenceName("refs/heads/"+newBranchName), head.Hash())
	err = repo.Storer.SetReference(ref)
	if err != nil {
		return err
	}

	msgSuccessfullyCreated := fmt.Sprintf("New branch %s based on %s created.", newBranchName, branch.Name)
	logs.Info(msgSuccessfullyCreated)

	msgConfirmation := fmt.Sprintf("Do you want to checkout on the new branch %s created?", ui.RenderElementSelected(newBranchName))
	checkoutOnBranchCreated, err := forms.AskConfirmation(msgConfirmation)
	if err != nil {
		return err
	}

	println(checkoutOnBranchCreated)

	if checkoutOnBranchCreated {
		// w, err := repo.Worktree()
		// if err != nil {
		// 	return err
		// }
		// err = w.Checkout(&git.CheckoutOptions{
		// 	Branch: plumbing.ReferenceName("refs/heads/" + newBranchName),
		// })
		// if err != nil {
		// 	return err
		// }
		// logs.Info(fmt.Sprintf("Checked out to new branch %s", newBranchName))
		return nil
	}

	return nil
}
