package actions

import (
	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
)

func DoBranchAction(repo git.Repository, branch branches.Branch, action BranchAction) error {
	switch action {
	case BranchActionExit:
		return nil
	case BranchActionDelete:
		return nil
	case BranchActionMerge:
		return nil
	case BranchActionNewBranch:
		return nil
	case BranchActionRename:
		return nil
	case BranchActionCheckout:
		return nil
	case BranchActionCopyName:
		return nil
	}

	return nil
}
