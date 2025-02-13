package actions

import (
	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
)

func DoBranchAction(repo git.Repository, branchSelected branches.BranchWithSymbol, allBranches []branches.BranchWithSymbol, head branches.Branch, action BranchAction) error {
	switch action {
	case BranchActionExit:
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
