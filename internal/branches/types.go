package branches

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Branch = *plumbing.Reference
type Repository = *git.Repository
type BranchWithSymbol struct {
	Name   string
	IsHead bool
	Branch Branch
}

type branchChoice struct {
	branches       []BranchWithSymbol
	cursor         int
	branchSelected BranchWithSymbol
}
