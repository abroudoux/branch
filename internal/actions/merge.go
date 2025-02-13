package actions

import (
	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
)

func merge(repo git.Repository, branch branches.BranchDetails) error {
	// if branch.Head {
	// 	// TODO
	// 	logs.WarnMsg("Cannot merge the current branch into itself. Please select a different target branch.")
	// 	return nil
	// }

	// branchRef, err := repo.Reference(branch.Branch.Name(), true)
	// if err != nil {
	// 	return err
	// }

	// options := git.MergeOptions{}
	// err = repo.Merge(*branchRef, options)
	// if err != nil {
	// 	return err
	// }

	// msgMergeSuccess := fmt.Sprintf("Branch %s merged.", ui.RenderElementSelected(branch.Name))
	// logs.Info(msgMergeSuccess)

	return nil
}
