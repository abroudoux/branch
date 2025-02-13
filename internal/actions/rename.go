package actions

import (
	"fmt"

	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/forms"
	"github.com/abroudoux/branch/internal/git"
	"github.com/abroudoux/branch/internal/logs"
	"github.com/abroudoux/branch/internal/ui"
	"github.com/go-git/go-git/v5/plumbing"
)

func rename(repo git.Repository, branch branches.BranchDetails, head branches.Branch) error {
	newName, err := forms.AskInput("Enter the new name of the branch: ")
	if err != nil {
		return err
	}

	if !branch.Head {
		err := checkout(repo, branch.Branch.Name().String(), head)
		if err != nil {
			return nil
		}
	}

	newRef := plumbing.NewHashReference(plumbing.ReferenceName("refs/heads/"+newName), head.Hash())
	err = repo.Storer.SetReference(newRef)
	if err != nil {
		return err
	}

	err = repo.Storer.RemoveReference(head.Name())
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Branch %s renamed on %s", ui.RenderElementSelected(head.Name().String()), ui.RenderElementSelected(newName))
	logs.Info(msg)
	return nil
}
