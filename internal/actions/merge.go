package actions

import (
	"fmt"

	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
	"github.com/abroudoux/branch/internal/logs"
	"github.com/abroudoux/branch/internal/ui"
)

func merge(repo git.Repository, branch branches.BranchDetails) error {
	branchRef, err := repo.Reference(branch.Branch.Name(), true)
	if err != nil {
		return err
	}

	options := git.MergeOptions{}
	err = repo.Merge(*branchRef, options)
	if err != nil {
		return err
	}

	msgMergeSuccess := fmt.Sprintf("Branch %s merged.", ui.RenderElementSelected(branch.Name))
	logs.Info(msgMergeSuccess)

	return nil
}
