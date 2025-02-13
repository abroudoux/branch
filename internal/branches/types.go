package branches

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Branch = *plumbing.Reference
type Repository = *git.Repository
type BranchDetails struct {
	Name   string
	Head   bool
	Branch Branch
}

type branchChoice struct {
	branches       []BranchDetails
	cursor         int
	branchSelected BranchDetails
}
