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

func delete(repo git.Repository, branch branches.BranchDetails) error {
	if branch.Head {
		logs.WarnMsg("You can't delete the current HEAD, checkout before.")
		return nil
	}

	confirmBranchDeletion, err := forms.AskConfirmation(fmt.Sprintf("Are you sure you want to delete %s?", ui.RenderElementSelected(branch.Name)))
	if err != nil {
		return err
	}

	if !confirmBranchDeletion {
		logs.Info("Branch deletion cancelled.")
		return nil
	}

	refName := plumbing.ReferenceName(branch.Branch.Name().String())
	err = repo.Storer.RemoveReference(refName)
	if err != nil {
		return fmt.Errorf("failed to delete branch: %w", err)
	}

	logs.Info(fmt.Sprintf("Branch %s successfully deleted.", ui.RenderElementSelected(branch.Name)))

	remoteBranchName := fmt.Sprintf("refs/remotes/origin/%s", branch.Name)
	_, err = repo.Reference(plumbing.ReferenceName(remoteBranchName), true)
	if err != nil {
		return nil
	}

	confirmRemoteDeletion, err := forms.AskConfirmation(fmt.Sprintf("Remote branch %s exists. Do you want to delete it?", ui.RenderElementSelected(branch.Name)))
	if err != nil {
		return err
	}

	if confirmRemoteDeletion {
		err = repo.DeleteBranch("testbranch")
		if err != nil {
			return fmt.Errorf("failed to delete remote branch: %w", err)
		}
		logs.Info(fmt.Sprintf("Remote branch %s successfully deleted.", ui.RenderElementSelected(branch.Name)))
		return nil
	}

	logs.Info("Remote branch deletion cancelled.")
	return nil
}
