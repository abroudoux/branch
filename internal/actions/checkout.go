package actions

import (
	"fmt"

	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/git"
	"github.com/abroudoux/branch/internal/logs"
	"github.com/abroudoux/branch/internal/ui"
	"github.com/go-git/go-git/v5/plumbing"
)

func checkout(repo git.Repository, branchName string, head branches.Branch) error {
	if branchName == head.Name().String() {
		warnMsg := fmt.Sprintf("You're alread on the branch %s", ui.RenderElementSelected(branchName))
		logs.WarnMsg(warnMsg)
		return nil
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	options := &git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branchName),
	}
	err = worktree.Checkout(options)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Successfully checked out branch %s.", ui.RenderElementSelected(branchName))
	logs.Info(msg)
	return nil
}
