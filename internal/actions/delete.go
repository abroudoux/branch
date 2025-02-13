package actions

import (
	"fmt"

	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/forms"
	"github.com/abroudoux/branch/internal/git"
	"github.com/abroudoux/branch/internal/logs"
	"github.com/abroudoux/branch/internal/ui"
)

func delete(repo git.Repository, branch branches.BranchDetails) error {
	if branch.IsHead {
		return fmt.Errorf("You can't delete the current HEAD.")
	}

	confirmBranchDeletion, err := forms.AskConfirmation(fmt.Sprintf("Are you sure you want to delete '%s'?", ui.RenderElementSelected(branch.Name)))
	if err != nil {
		return err
	}

	if !confirmBranchDeletion {
		return fmt.Errorf("Branch deletion cancelled.")
	}

	err = repo.DeleteBranch(branch.Branch.Name().String())
	if err != nil {
		return err
	}

	logs.Info(fmt.Sprintf("Branch '%s' successfully deleted.", ui.RenderElementSelected(branch.Name)))
	return nil
}
