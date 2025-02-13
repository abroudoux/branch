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
		return delete(repo, branchSelected)
	case BranchActionMerge:
		return merge(repo, branchSelected)
	case BranchActionNewBranch:
		return createNewBranch(repo, branchSelected, allBranches, head)
	case BranchActionCheckout:
		return checkout(repo, branchSelected.Branch.Name().String(), head)
	case BranchActionCopyName:
		return copyBranchName(branchSelected)
	}

	return nil
}
