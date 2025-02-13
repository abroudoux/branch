package branches

import (
	"strings"
)

func CreateBranchesWithDetails(branches []Branch, head Branch) []BranchDetails {
	branchesWithSymbols := []BranchDetails{}
	for _, branch := range branches {
		branchName := string(branch.Name())
		branchNameTrimed := strings.TrimPrefix(string(branchName), "refs/heads/")

		if branch.Name() == head.Name() {
			branchesWithSymbols = append(branchesWithSymbols, BranchDetails{
				Name:   branchNameTrimed,
				IsHead: true,
				Branch: branch,
			})
		} else {
			branchesWithSymbols = append(branchesWithSymbols, BranchDetails{
				Name:   branchNameTrimed,
				IsHead: false,
				Branch: branch,
			})
		}
	}

	return branchesWithSymbols
}
