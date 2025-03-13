package internal

import (
	"fmt"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func GetRepository(dir string) (*Repository, error) {
	currentRepository, err := git.PlainOpen(dir)
	if err != nil {
		return nil, err
	}

	return &Repository{currentRepository}, nil
}

func (repository *Repository) getBranches() ([]Branch, error) {
	branchIter, err := repository.Branches()
	if err != nil {
		return nil, err
	}

	var branches []Branch
	err = branchIter.ForEach(func(ref Branch) error {
		branches = append(branches, ref)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return branches, nil
}

func (repository *Repository) SelectBranch() (Branch, error) {
	branches, err := repository.getBranches()
	if err != nil {
		return nil, err
	}

	head, err := repository.Head()
	if err != nil {
		return nil, err
	}

	p := tea.NewProgram(initialBranchChoiceModel(branches, head))
	m, err := p.Run()
	if err != nil {
		return nil, err
	}

	branchSelected := m.(branchChoice).branchSelected
	return branchSelected, nil
}

func (repository *Repository) DoBranchAction(branchSelected Branch, action branchAction) error {
	switch action {
	case BranchActionExit:
		log.Info("Exiting..")
		return nil
	case BranchActionDelete:
		return repository.delete(branchSelected)
	case BranchActionMerge:
		return repository.merge(branchSelected)
	case BranchActionNewBranch:
		return repository.createNewBranch(branchSelected)
	case BranchActionCheckout:
		return repository.checkout(branchSelected)
	case BranchActionCopyName:
		return copyBranchName(branchSelected)
	}

	return nil
}

func (repository *Repository) delete(branch Branch) error {
	if repository.isHead(branch) {
		log.Warn("You can't delete the current HEAD, checkout before.")
		return nil
	}

	msgConfirmation := fmt.Sprintf("Are you sure you want to delete %s ?", renderElSelected(branch.Name().Short()))
	confirmBranchDeletion, err := getConfirmation(msgConfirmation)
	if err != nil {
		return err
	}

	if !confirmBranchDeletion {
		log.Info("Branch deletion cancelled.")
		return nil
	}

	refName := plumbing.ReferenceName(branch.Name())
	err = repository.Storer.RemoveReference(refName)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Branch %s successfully deleted locally.", renderElSelected(branch.Name().Short())))
	return nil
}

func (repository *Repository) merge(branch Branch) error {
	if repository.isHead(branch) {
		log.Warn("Cannot merge the current branch into itself. Please select a different target branch.")
		return nil
	}

	confirmMerge, err := getConfirmation(fmt.Sprintf("Are you sure you want to merge %s ?", renderElSelected(branch.Name().Short())))
	if err != nil {
		return err
	}

	if !confirmMerge {
		log.Info("Merge cancelled.")
		return nil
	}

	log.Warn("Not implemented yet..")
	return nil
}

func (repository *Repository) isHead(branch Branch) bool {
	head, err := repository.Head()
	if err != nil {
		return false
	}

	return head.Name() == branch.Name()
}

func (repository *Repository) createNewBranch(branch Branch) error {
	if !repository.isHead(branch) {
		log.Warn("You need to create a branch from the head, move on it first.")
		return nil
	}

	for {
		newBranchName, err := readInput("Enter the name of the new branch: ")
		if err != nil {
			return fmt.Errorf("failed to get input: %w", err)
		}

		if repository.isBranchNameAlreadyUsed(newBranchName) {
			warnMsg := fmt.Sprintf("%s is already used, please choose another name.", renderElSelected(newBranchName))
			log.Warn(warnMsg)
			continue
		}

		head, err := repository.Head()
		if err != nil {
			return err
		}

		newBranch := plumbing.NewHashReference(plumbing.ReferenceName("refs/heads/"+newBranchName), head.Hash())
		err = repository.Storer.SetReference(newBranch)
		if err != nil {
			return fmt.Errorf("failed to create new branch: %w", err)
		}

		msgSuccessfullyCreated := fmt.Sprintf("New branch %s based on %s created.", renderElSelected(newBranchName), renderElSelected(branch.Name().Short()))
		log.Info(msgSuccessfullyCreated)

		msgConfirmation := fmt.Sprintf("Do you want to checkout on the new branch %s created?", renderElSelected(newBranchName))
		checkoutOnBranchCreated, err := getConfirmation(msgConfirmation)
		if err != nil {
			return fmt.Errorf("failed to get confirmation: %w", err)
		}

		if checkoutOnBranchCreated {
			err := repository.checkout(newBranch)
			if err != nil {
				return fmt.Errorf("failed to checkout new branch: %w", err)
			}
		}

		return nil
	}
}

func (repository *Repository) isBranchNameAlreadyUsed(newBranchName string) bool {
	branches, err := repository.getBranches()
	if err != nil {
		return false
	}

	for _, branch := range branches {
		if branch.Name().Short() == newBranchName {
			return true
		}
	}

	return false
}

func (repository *Repository) checkout(branch Branch) error {
	if repository.isHead(branch) {
		warnMsg := fmt.Sprintf("You're alread on the branch %s", renderElSelected(branch.Name().Short()))
		log.Warn(warnMsg)
		return nil
	}

	worktree, err := repository.Worktree()
	if err != nil {
		return err
	}

	options := &git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/heads/" + branch.Name().Short()),
	}
	err = worktree.Checkout(options)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Successfully checked out branch %s.", renderElSelected(branch.Name().Short()))
	log.Info(msg)
	return nil
}

func copyBranchName(branch Branch) error {
	if clipboard.Unsupported {
		return fmt.Errorf("Clipboard not supported on this plateform.")
	}

	err := clipboard.WriteAll(branch.Name().Short())
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Branch name %s copied to clipboard.", renderElSelected(branch.Name().Short())))
	return nil
}
