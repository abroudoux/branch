package actions

import (
	"fmt"

	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/forms"
	"github.com/abroudoux/branch/internal/logs"
	"github.com/abroudoux/branch/internal/ui"
	"github.com/charmbracelet/log"
	"github.com/go-git/go-git/v5"
)

func merge(repo git.Repository, branch branches.BranchDetails) error {
	if branch.Head {
		logs.WarnMsg("Cannot merge the current branch into itself. Please select a different target branch.")
		return nil
	}

	confirmMerge, err := forms.AskConfirmation(fmt.Sprintf("Are you sure you want to merge %s?", ui.RenderElementSelected(branch.Name)))
	if err != nil {
		return err
	}

	if !confirmMerge {
		logs.Info("Merge cancelled.")
		return nil
	}

	log.Warn("Not implemented yet..")
	return nil
}
