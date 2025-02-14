package actions

import (
	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
	"github.com/abroudoux/branch/internal/logs"
)

func merge(_ git.Repository, branch branches.BranchDetails) error {
	if branch.Head {
		logs.WarnMsg("Cannot merge the current branch into itself. Please select a different target branch.")
		return nil
	}

	logs.WarnMsg("Not implemented yet")
	return nil
}
