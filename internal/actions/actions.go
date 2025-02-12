package actions

import (
	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
)

func DoBranchAction(repo git.Repository, branch branches.BranchWithSymbol, action BranchAction) error {
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
		return nil
	case BranchActionRename:
		return nil
	case BranchActionCheckout:
		return nil
	case BranchActionCopyName:
		return copyBranchName(branch)
	}

	return nil
}
