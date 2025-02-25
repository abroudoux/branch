package actions

import "github.com/abroudoux/branch/internal/branches"

type BranchAction int

const (
	BranchActionExit BranchAction = iota
	BranchActionDelete
	BranchActionMerge
	BranchActionNewBranch
	BranchActionCheckout
	BranchActionCopyName
)

type branchActionChoice struct {
	actions        []BranchAction
	cursor         int
	actionSelected BranchAction
	branchSelected branches.BranchDetails
}

type mergeOptions = &git.MergeOptions
