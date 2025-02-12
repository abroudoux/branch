package actions

import (
	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
)

func DoBranchAction(repo git.Repository, branch branches.BranchWithSymbol, head branches.Branch, action BranchAction) error {
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
		return createNewBranch(repo, branch, head)
	case BranchActionRename:
		return rename(repo, branch, head)
	case BranchActionCheckout:
		return checkout(repo, branch.Branch.Name().String())
	case BranchActionCopyName:
		return copyBranchName(branch)
	}

	return nil
}
