package actions

import (
	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
	"github.com/abroudoux/branch/internal/logs"
)

func DoBranchAction(repo git.Repository, branchSelected branches.BranchDetails, allBranches []branches.BranchDetails, head branches.Branch, action BranchAction) error {
	switch action {
	case BranchActionExit:
		logs.Info("Exiting..")
		return nil
	case BranchActionDelete:
		// TODO
		return nil
		// return delete(repo, branch)
	case BranchActionMerge:
		return nil
	case BranchActionNewBranch:
		return createNewBranch(repo, branchSelected, allBranches, head)
	case BranchActionRename:
		return rename(repo, branchSelected, head)
	case BranchActionCheckout:
		return checkout(repo, branchSelected.Branch.Name().String(), head)
	case BranchActionCopyName:
		return copyBranchName(branchSelected)
	}

	return nil
}
