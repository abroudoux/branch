package internal

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Repository struct {
	*git.Repository
}

type Branch = *plumbing.Reference

type branchChoice struct {
	head           Branch
	branches       []Branch
	cursor         int
	branchSelected Branch
}

type branchAction int

type branchActionChoice struct {
	actions        []branchAction
	cursor         int
	actionSelected branchAction
	branchSelected Branch
}
